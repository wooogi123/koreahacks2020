package controllers

import (
  "github.com/gin-gonic/gin"
  "../models"
  "net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) SignIn(c *gin.Context) {
  if err := c.ShouldBindJSON(&userModel); err != nil {
    log.Println("Data not Bind!")
	c.JSON(http.StatusBadRequest, gin.H{
	  "status": "error",
	  "content": "Data not Bind!",
	})
	return
  }
  user, err := userModel.SignIn()
  if err != nil {
    log.Println("Sign In Error!")
	c.JSON(http.StatusUnauthorized, gin.H{
	  "status": "error",
	  "content": err,
	})
	return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
	"content": user,
  })
  return
}

func (u UserController) SignUp(c *gin.Context) {
  if err := c.ShouldBindJSON(&userModel); err != nil {
    log.Println("Data not Bind!")
	c.JSON(http.StatusBadRequest, gin.H{
	  "status": "error",
	  "content": "Data not Bind!",
	})
	return
  }
  user, err := userModel.SignUp()
  if err != nil {
    log.Println("Sign Up Error!")
	c.JSON(http.StatusBadRequest, gin.H{
	  "status": "error",
	  "content": err,
	})
	return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
	"content": user,
  })
  return
}

func (u UserController) Update(c *gin.Context) {
  if err := c.ShouldBindJSON(&userModel); err != nil {
    log.Println("Data not Bind!")
	c.JSON(http.StatusBadRequest, gin.H{
	  "status": "error",
	  "content": "Data not Bind!",
	})
	return
  }
  user, err := userModel.Update()
  if err != nil {
    log.Println("Update Error!")
	c.JSON(http.StatusBadRequest, gin.H{
	  "status": "error",
	  "content": err,
	})
	return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
	"content": user,
  })
  return
}
