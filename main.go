package main

import (
	"encoding/json"
	"github.com/donfrigo/quiz/handlers"
	"github.com/donfrigo/quiz/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main () {

	// get execution path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// read json files for initialization
	quiz_json, _ := ioutil.ReadFile(dir + "/quiz.json")
	results_json, _ := ioutil.ReadFile(dir + "/results.json")

	// parse json file quiz.json
	questions := structs.Questions{}
	if err := json.Unmarshal(quiz_json, &questions); err != nil {
		return
	}

	// parse json file results.json
	results := structs.Results{}
	if err := json.Unmarshal(results_json, &results); err != nil {
		return
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Define the HTTP routes
	e.File("/", dir+"/public/index.html")
	e.GET("/quiz", handlers.GetQuestions(questions))
	e.PUT("/results", handlers.UpdateResults(&results))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
