package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     "906162141711-i44qvcal8epjbh38t5kc9mpbk0gvehla.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/auth/callback", // todo: move this to env
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

	user := models.NewUser(userData["email"].(string), userData["name"].(string),
		userData["hd"].(string), userData["picture"].(string),
		userData["verified_email"].(bool))

	err = user.Save()

	if e := setRefreshCookie(user, c); e != nil {
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

	time := time.Now().Add(time.Hour * 24 * 30 * 6) // 6 months
	// refresh_token, err := NewToken(user.Id, user.IsAdmin, user.IsEditor, true, time)
	refresh_token, err := NewToken(user.Id, true, true, true, time)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Expires:  time,
		Secure:   true,
		HTTPOnly: true, // accessible only by http (not js)
		Name:     "refresh_token",
		Value:    refresh_token,
	})

	return nil
}

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
