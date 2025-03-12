package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/server"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(gin.Recovery())

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	server.RegisterRoutes(router)

	// // Get user value
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// // Authorized group (uses gin.BasicAuth() middleware)
	// // Same than:
	// // authorized := router.Group("/")
	// // authorized.Use(gin.BasicAuth(gin.Credentials{
	// //	  "foo":  "bar",
	// //	  "manu": "123",
	// //}))
	// authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// /* example curl for /admin with basicauth header
	//    Zm9vOmJhcg== is base64("foo:bar")

	// 	curl -X POST \
	//   	http://localhost:8080/admin \
	//   	-H 'authorization: Basic Zm9vOmJhcg==' \
	//   	-H 'content-type: application/json' \
	//   	-d '{"value":"bar"}'
	// */
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return router
}

func main() {
	config := cfg.SetServerConfig()
	connector.Connect()
	defer connector.Disconnect()

	router := setupRouter()
	router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
