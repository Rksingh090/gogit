package gitroutes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Rksingh090/gogit/git"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v55/github"
)

func ExchangeCode(c *gin.Context) {
	inst_id := c.Query("installation_id")

	if inst_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": "installation_id is required",
		})
	}

	t, err := git.GetToken()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   true,
			"message": "error getting app jwt token",
		})
	}

	gc := github.NewClient(nil).WithAuthToken(t)

	installationId, err := strconv.ParseInt(inst_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": true,
			"data":  err.Error(),
		})
	}

	iToken, resp, err := gc.Apps.CreateInstallationToken(context.Background(), installationId, &github.InstallationTokenOptions{})
	defer resp.Body.Close()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  err.Error(),
		})
	}

	// Print the response
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  iToken,
	})
}
