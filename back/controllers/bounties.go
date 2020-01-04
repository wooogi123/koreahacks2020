package controllers

import (
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "../models"
  "../forms"
)

type BountyController struct{}

var btModel = new (models.Bounty)

func (b BountyController) ListAll(c *gin.Context) {
  bts, err := btModel.ListAll()
  if err != nil {
    log.Println("ListAll Error")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "ListAll Error",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "content": bts,
  })
  return
}

func (b BountyController) Select(c *gin.Context) {
   model := new(forms.BountySelect)
   if err := c.ShouldBindJSON(&model); err != nil {
     log.Println("Data not Bind!")
     c.JSON(http.StatusBadRequest, gin.H{
       "status": "error",
       "content": "Data not Bind!",
     })
     return
   }
   bt, err := btModel.Select(*model)
   if err != nil {
     log.Println("Select Error!")
     c.JSON(http.StatusBadRequest, gin.H{
       "status": "error",
       "content": err,
     })
     return
   }

   c.JSON(http.StatusOK, gin.H{
     "status": "success",
     "content": bt,
   })
   return
}

func (b BountyController) Register(c *gin.Context) {
  model := new(forms.BountyRegister)
  if err := c.ShouldBindJSON(&model); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  bt, err := btModel.Register(*model)
  if err != nil {
    log.Println("Register Error!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "content": bt,
  })
  return
}

func (b BountyController) Update(c *gin.Context) {
  model := new(form.BountyUpdate)
  if err := c.ShouldBindJSON(&model); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  bt, err := btModel.Update(*model)
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
    "content": bt,
  })
  return
}
