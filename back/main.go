package main

import (
  "log"
  "github.com/gin-gonic/gin"

  "./controllers"
  "./api_bounty"
)

var userControl = new(controllers.UserController)

func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
    c.Header("Access-Control-Allow-Credentials", "true")
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST")
    c.Next()
  }
}

func main() {
  router := gin.Default()
  router.Use(CORSMiddleware())

  Users := router.Group("/Users")
  {
    Users.POST("/SignIn", userControl.SignIn)
    Users.POST("/SignUp", userControl.SignUp)
    Users.POST("/Update", userControl.Update)
  }

  Bounty := router.Group("/Bounty")
  {
    Bounty.GET("/list", api_bounty.ListBounty)
    Bounty.POST("/select", api_bounty.SelectBounty)
    Bounty.POST("/register", api_bounty.RegisterBounty)
    Bounty.POST("/update", api_bounty.UpdateBounty)
  }

  log.Fatal(router.Run())
}
