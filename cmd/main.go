package main

import (
	"log"
	"os"
	"time"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
	"github.com/SassoStorTo/migra-studenti/pkg/router"
	dbutils "github.com/SassoStorTo/migra-studenti/pkg/services/databaseutils"
	"github.com/SassoStorTo/migra-studenti/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	utils.InitStoreSess()
	utils.InitSaveFolder()
	err := database.ConnectDB()
	if err != nil {
		log.Panic(err)
	}
	// dbutils.Reset()
	dbutils.SetupDb()

	engine := html.New("./views", ".html")

	engine.AddFunc("formatDate", func(t time.Time) string {
		return t.Format("2006-01-02") // Returns date in YYYY-MM-DD format
	})
	engine.AddFunc("isInList", utils.IsItemInList)
	// todo: add this fuction to utils
	engine.AddFunc("listToString", func(list []string) string {
		ret := ""
		for _, v := range list {

			ret += v + " ,"
		}
		return ret
	})

	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: true,
		AppName:       "italici",
		Views:         engine,
		// ViewsLayout:   "frontends/mypages/template",
	})

	// todo: Move this to configs
	// multiWriter := io.MultiWriter(os.Stdout)
	// if os.Getenv("LOG_FILE") == "TRUE" {
	// 	log.Println("Logging to file")
	// 	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer logFile.Close()
	// 	multiWriter = io.MultiWriter(multiWriter, logFile)
	// 	log.SetOutput(multiWriter)
	// }

	//Todo: add this to
	app.Use(logger.New(logger.Config{
		Next: nil,
		Done: nil,
		// Format:        "${date} ${time} | ${status} | ${latency} | ${ip} | ${method} | ${url} | ${error} | ${body} | ${reqHeaders} \n",
		// Format:        "${status} | ${latency} | ${ip} | ${method} | ${url} | ${body} | \n ${reqHeaders} \n",
		// Format:        "${date} ${time} | ${status} | ${latency} | ${ip} | ${method} | ${url} | ${error} | ${body} \n",
		Format:        "${date} ${time} | ${status} | ${latency} | ${ip} | ${method} | ${url} | ${error} \n",
		TimeFormat:    "02-01-2006 15:04:05",
		TimeZone:      "Local",
		TimeInterval:  time.Millisecond,
		Output:        os.Stdout, // todo: change this to multiWriter
		DisableColors: false,
	}))

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
