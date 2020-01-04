package main

import (
  "log"
  "github.com/gin-gonic/gin"

  "./controllers"
)

var userControl = new(controllers.UserController)
var bountyControl = new(controllers.BountyController)

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
    Bounty.GET("/ListAll", bountyControl.ListAll)
    Bounty.POST("/Select", bountyControl.Select)
    Bounty.POST("/Register", bountyControl.Register)
    Bounty.POST("/Update", bountyControl.Update)
  }

  log.Fatal(router.Run())
}
