package main

import (
  "log"
  "github.com/gin-gonic/gin"

  "./api_users"
)

func main() {
  router := gin.Default()

  Users := router.Group("/Users")
  {
    Users.POST("/login", api_users.LoginUser)
	Users.POST("/register", api_users.RegisterUser)
	Users.POST("/update", api_users.UpdateUser)
  }

  log.Fatal(router.Run())
}
