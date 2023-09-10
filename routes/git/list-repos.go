package gitroutes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v55/github"
)

func ListRepos(c *gin.Context) {

	token := c.Param("id")
	client := github.NewClient(nil).WithAuthToken(token)

	repos, _, err := client.Apps.ListRepos(context.Background(), &github.ListOptions{})

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "working",
		"repos":   repos,
	})
}
