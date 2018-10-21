package handlers

import (
	"encoding/json"
	"github.com/donfrigo/quiz/structs"
	"github.com/labstack/echo"
	"io/ioutil"
	"math"
	"net/http"
)

func GetQuestions(questions structs.Questions) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, questions)
	}
}

func UpdateResults(results *structs.Results) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Read the Body content
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		}

		result := structs.Result{}
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			c.JSON(http.StatusBadRequest, "Malformed JSON")
		}

		// Checking if request has the required fields
		if result.User == "" {
			c.JSON(http.StatusBadRequest,"Invalid JSON Format")
		}

		if result.Result == "" {
			c.JSON(http.StatusBadRequest,"Invalid JSON Format")
		}

		*results = append(*results, result)

		percentage := calculatePercentage(*results)
		percentageObject := map[string]float64{"percentage": percentage}

		return c.JSON(http.StatusCreated, percentageObject)
	}
}

func calculatePercentage (results structs.Results) float64{

	length := len(results)
	correctAnswers := results[length-1].Result
	counter := 0.0

	if length == 1 {
		// first one to complete the quiz
		return 100.0

	} else {

		for j := 0; j <= length-2; j++ {
			result := results[j].Result
			if result < correctAnswers {
				counter++
			}
		}

		percentage := math.Round(counter / float64(length-1) * 100.0)

		return percentage
	}
}
