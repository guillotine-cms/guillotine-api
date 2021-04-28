package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.POST("/content-type", func(context *gin.Context) {

		// Create content type
	})

	r.GET("/content-type", func(context *gin.Context) {

        // Getting the gist of how things ist!
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
		}

		es, err := elasticsearch.NewClient(cfg)

		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}

		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		defer res.Body.Close()
		log.Println(res)
		// Create content type
	})

	r.Run()
}
