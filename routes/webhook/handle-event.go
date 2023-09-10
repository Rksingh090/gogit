package webhookRoute

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleWebHook(c *gin.Context) {

	fmt.Println("Webhook Called")

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	eventType := c.GetHeader("X-GitHub-Event")
	if eventType == "push" {
		// This is a push event
		fmt.Println("Received push event")
	}

	// Create a JSON.RawMessage to store the raw JSON data
	rawData := json.RawMessage(body)

	// Or unmarshal it into a map or struct if needed
	var data map[string]interface{}
	if err := json.Unmarshal(rawData, &data); err != nil {
		fmt.Println("Error unmarshaling raw JSON:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "OK",
		"data":    data,
	})
	return
}
