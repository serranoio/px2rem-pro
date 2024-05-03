package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	ConversionFactor string `json:"conversionFactor"`
	Input            string `json:"input"`
}

func createServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/convert", func(c *gin.Context) {

		// Parse the JSON body into your struct
		var message Message
		if err := json.NewDecoder(c.Request.Body).Decode(&message); err != nil {
			c.JSON(http.StatusBadRequest, "bad")
		}

		// Do something with the parsed data
		fmt.Printf("Received data: %+v\n", message)
		cF, err := strconv.ParseFloat(message.ConversionFactor, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "bad")
		}

		config := config{
			conversionFactor: cF,
			doNotInclude:     "",
		}

		contents := parseContents(config, message.Input)

		contentsStruct := struct {
			Contents string `json:"contents"`
		}{
			Contents: contents,
		}

		c.JSON(http.StatusOK, contentsStruct)
	})

	r.Run(":6883")
}
