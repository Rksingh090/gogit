package gitroutes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Rksingh090/gogit/git"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v55/github"
)

func ListInstallations(c *gin.Context) {

	t, err := git.GetToken()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}

	gc := github.NewClient(nil).WithAuthToken(t)

	inst, _, err := gc.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"error":        false,
		"message":      "working",
		"installation": inst,
	})
}
