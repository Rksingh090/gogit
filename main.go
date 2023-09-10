package main

import (
	"fmt"
	"net/http"

	gitroutes "github.com/Rksingh090/gogit/routes/git"
	webhookRoute "github.com/Rksingh090/gogit/routes/webhook"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		fmt.Print("JGJJ")
		c.JSON(http.StatusOK, gin.H{
			"message": "API is running",
		})
	})

	gitRoute := r.Group("git")
	{
		gitRoute.GET("/installations", gitroutes.ListInstallations)
		gitRoute.GET("/repos/:id", gitroutes.ListRepos)
		gitRoute.GET("/exchange", gitroutes.ExchangeCode)
		gitRoute.POST("/event", webhookRoute.HandleWebHook)
	}

	r.Run(":9000")
}
