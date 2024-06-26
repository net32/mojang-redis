package server

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const VERSION = "1.0.4"

var name, author, url = "mojang-redis", "NeT32", "https://github.com/net32/mojang-redis"

func InitServer() {
	log.Printf("Starting %s developed by %s.", name, author)
	log.Printf("API Version: %s", VERSION)
	log.Printf("GitHub: %s", url)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    name,
			"author":  author,
			"url":     url,
			"version": VERSION,
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users/profiles/minecraft/:name", func(c *gin.Context) {
		userName := c.Params.ByName("name")
		writeJsonResponse(c, UsernameToUUID(userName))
	})
	r.POST("/profiles/minecraft", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
		}
		writeJsonResponse(c, UsernamesToUUIDs(b))
	})
	// This endpoint has been deprecated by Mojang and was removed on 13 September 2022 at 9:25 AM CET to "improve player safety and data privacy"
	r.GET("/user/profiles/:uuid/*action", func(c *gin.Context) {
		uuid := c.Params.ByName("uuid")
		writeJsonResponse(c, UuidToNameHistory(uuid))
	})
	r.GET("/session/minecraft/profile/:uuid", func(c *gin.Context) {
		uuid := c.Params.ByName("uuid")
		unsigned, exist := c.GetQuery("unsigned")
		if !exist {
			unsigned = "true"
		}
		writeJsonResponse(c, UuidToProfile(uuid, unsigned))
	})
	r.GET("/session/minecraft/hasJoined", func(c *gin.Context) {
		userName, _ := c.GetQuery("username")
		serverId, _ := c.GetQuery("serverId")
		writeJsonResponse(c, HasJoined(userName, serverId))
	})
	r.GET("/blockedservers", func(c *gin.Context) {
		response := BlockedServers()
		c.String(response.Code, response.Json)
	})
	r.GET("/auth/haspaid/:name", func(c *gin.Context) {
		userName := c.Params.ByName("name")
		response, _ := HasPaid(userName)
		c.String(200, response)
	})
	addr := ":" + GetEnv("PORT", "8080")
	log.Println("Listen: " + addr)
	r.Run(addr)
}

func writeJsonResponse(c *gin.Context, response MojangResponse) {
	c.Header("Content-Type", "application/json")
	c.String(response.Code, response.Json)
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
