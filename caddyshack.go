package main

import (
	"caddyshack/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg := readConfig("caddyshack.yml")

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	r.GET("/check", func(c *gin.Context) {
		domain := c.Query("domain")
		status := 401

		for _, s := range cfg.Domains {
			if strings.HasSuffix(domain, s) {
				status = 200
				break
			}
		}

		c.JSON(status, gin.H{
			"domain_checked": domain,
			"status":         status,
		})
	})

	_ = r.Run(cfg.Port)
}

func readConfig(filename string) models.Configuration {
	config := models.Configuration{}

	file, err := ioutil.ReadFile(filename)
	failOnError(err, "Failed to read in configuration file")

	err = yaml.Unmarshal(file, &config)
	failOnError(err, "Failed to parse configuration file")

	return config
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
