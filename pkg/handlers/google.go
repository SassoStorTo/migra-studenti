package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/SassoStorTo/migra-studenti/pkg/models"
	"github.com/SassoStorTo/migra-studenti/pkg/services/auth"
	"github.com/SassoStorTo/migra-studenti/pkg/services/users"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	stateMutex sync.Mutex
	states     = make(map[string]bool)
)

func HandleLogin(c *fiber.Ctx) error {
	state, err := generateState()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	stateMutex.Lock()
	states[state] = true
	stateMutex.Unlock()

	url := googleOauthConfig.AuthCodeURL(state)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func HandleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if !validateState(state) {
		fmt.Println("Invalid oauth state")
		return c.Redirect("/", http.StatusTemporaryRedirect)
	}

	stateMutex.Lock()
	delete(states, state)
	stateMutex.Unlock()

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Code exchange failed: %s\n", err.Error())
		return c.Redirect("/", http.StatusTemporaryRedirect)
	}

	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Printf("Failed getting user info: %s\n", err.Error())
		return c.Redirect("/", http.StatusTemporaryRedirect)
	}

	defer response.Body.Close()
	userData, err := parseUserInfo(response.Body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	usr := models.NewUser(userData["email"].(string), userData["name"].(string),
		userData["hd"].(string), userData["picture"].(string),
		userData["verified_email"].(bool))

	usr.IsEditor = false
	usr.IsAdmin = false

	if len(*users.GetAll()) == 0 {
		usr.IsEditor = true
		usr.IsAdmin = true
	}

	err = usr.Save()

	if e := setRefreshCookie(usr, c); e != nil {
		return e
	}

	if err != nil {
		return c.Redirect("/wait-accept")
	}

	return c.Redirect("/")
}

func setRefreshCookie(user *models.User, c *fiber.Ctx) error {
	user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	refresh_token, exp, err := auth.GetRefreshToken(user)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  exp,
		Secure:   true,
		HTTPOnly: true, // accessible only by http (not js)
		Name:     "refresh_token",
		Value:    refresh_token,
	})

	return nil
}

// todo: save this to redis
func generateState() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}

func validateState(state string) bool {
	stateMutex.Lock()
	defer stateMutex.Unlock()
	return states[state]
}

func parseUserInfo(body io.Reader) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.NewDecoder(body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
