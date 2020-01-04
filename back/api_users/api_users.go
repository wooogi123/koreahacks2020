package api_users

import (
  "os"
  "log"
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB


type User struct {
  Nickname  string `json:"nickname"`
  Email     string `json:"email" binding:"required"`
  Password  string `json:"password" binding:"required"`
  Group     string `json:"group"`
  Star         int `json:"star"`
  Finish_count int `json:"finish_count"`
  Point        int `json:"point"`
}

func (u User) sel() (user User, err error) {
  row := db.QueryRow(`
    SELECT Users.nickname, Users.email, Users.password, Users.group, Users.star, Users.finish_count, Users.point
    FROM Users
    WHERE Users.email = ? AND Users.password = ?
  `, u.Email, u.Password)

  err = row.Scan(
    &user.Nickname,
    &user.Email,
    &user.Password,
    &user.Group,
    &user.Star,
    &user.Finish_count,
    &user.Point)
  if err != nil {
    return
  }
  return
}

func (u User) ins() (id int64, err error) {
  stmt, err := db.Prepare(`
    INSERT INTO Users (
    Users.nickname, Users.email, Users.password, Users.group, Users.star, Users.finish_count, Users.point
    ) VALUES (?, ?, ?, ?, ?, ?, ?)
  `)
  if err != nil {
    return
  }
  rs, err := stmt.Exec(u.Nickname, u.Email, u.Password, u.Group, u.Star, u.Finish_count, u.Point)
  if err != nil {
    return
  }
  id, err = rs.LastInsertId()
  if err != nil {
    return
  }
  defer stmt.Close()
  return
}

func (u User) upd() (user User, err error) {
  stmt, err := db.Prepare(`
    UPDATE Users
    SET Users.nickname = ?, Users.password = ?, Users.group = ?, Users.star = ?, Users.finish_count = ?, Users.point = ?, updated_at = CURRENT_TIMESTAMP
    WHERE Users.email = ?
  `)
  if err != nil {
    return
  }
  _, err = stmt.Exec(u.Nickname, u.Password, u.Group, u.Star, u.Finish_count, u.Point, u.Email)
  if err != nil {
    return
  }
  defer stmt.Close()

  user, err = u.sel()
  if err != nil {
    return
  }
  return
}

func LoginUser(c *gin.Context) {
  var err error
  db, err = sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME"))

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var u User
  if err := c.ShouldBindJSON(&u); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  user, err := u.sel()
  if err != nil {
    log.Println("Authorized Failed!")
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

func RegisterUser(c *gin.Context) {
  var err error
  db, err = sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME"))

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var u User
  if err := c.ShouldBindJSON(&u); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    log.Println("Data not Bind!")
    return
  }

  id, err := u.ins()
  if err != nil {
    log.Println("Data insert Failed!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "id": id,
    "content": u,
  })
  return
}

func UpdateUser(c *gin.Context) {
  var err error
  db, err = sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME"))

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var u User
  if err := c.ShouldBindJSON(&u); err != nil {
    log.Println("Data bind Failed!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data bind Failed!",
    })
    return
  }

  user, err := u.upd()
  if err != nil {
    log.Println("Date update Failed!")
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
