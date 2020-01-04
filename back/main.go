package main

import (
  "log"
  "github.com/gin-gonic/gin"

  "./api_users"
  "./api_bounty"
)

func main() {
  router := gin.Default()

  Users := router.Group("/Users")
  {
    Users.POST("/login", api_users.LoginUser)
	Users.POST("/register", api_users.RegisterUser)
	Users.POST("/update", api_users.UpdateUser)
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
