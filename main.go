package main

import (
	"errors"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0"

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func checkURL(URL string) error {
	match, err := regexp.MatchString(`((http|https):\/\/|)(www\.|)youtube\.com\/(channel\/|user\/)[a-zA-Z0-9\-]{1,}`, URL)
	if err != nil {
		panic(err)
	}
	if match != true {
		return errors.New("The submitted URL isn't a channel URL")
	}
	return nil
}

func extractHandle(c *gin.Context) {
	URL := c.Request.URL.Path[9:]

	err := checkURL(URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}

	out, err := exec.Command("./extract_videos.sh", URL).Output()
	if err != nil {
		panic(err)
	}

	ids := deleteEmpty(strings.Split(string(out), "\n"))

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "seeds": ids})
}

func rootHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "version": version})
}

func main() {
	router := gin.Default()
	router.UseRawPath = true
	router.RedirectFixedPath = false
	v1 := router.Group("/")
	{
		v1.GET("/", rootHandle)
		v1.GET("/extract/*url", extractHandle)
	}
	router.Run(":4243")
}
