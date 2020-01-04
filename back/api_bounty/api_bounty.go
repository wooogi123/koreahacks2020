package api_bounty

import (
  "os"
  "log"
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Bounty struct {
  Id               int `json:"id"`
  State            int `json:"state"`
  Title         string `json:"title" binding:"required"`
  Text          string `json:"text"`
  Bounty           int `json:"bounty"`
  Time_limit    string `json:"time_limit"`
}

func (b Bounty) sel() (bt Bounty, err error) {
  row := db.QueryRow(`
    SELECT Bounty.id, Bounty.state, Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit
    FROM Bounty
    WHERE Bounty.id = ?
  `, b.Id)

  err = row.Scan(&bt.Id, &bt.State, &bt.Title, &bt.Text, &bt.Bounty, &bt.Time_limit)
  if err != nil {
    return
  }
  return
}

func (b Bounty) selAll() (bts[] Bounty, err error) {
  rows, err := db.Query(`
    SELECT Bounty.id, Bounty.state, Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit
    FROM Bounty
    WHERE Bounty.state = 1
  `)
  if err != nil {
    return
  }
  for rows.Next() {
    var bt Bounty
    rows.Scan(&bt.Id, &bt.State, &bt.Title, &bt.Text, &bt.Bounty, &bt.Time_limit)
    bts = append(bts, bt)
  }
  defer rows.Close()
  return
}

func (b Bounty) ins() (id int64, err error) {
  stmt, err := db.Prepare(`
    INSERT INTO Bounty (Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit)
    VALUES (?, ?, ?, ?)
  `)
  if err != nil {
    return
  }
  rs, err := stmt.Exec(b.Title, b.Text, b.Bounty, b.Time_limit)
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

func (b Bounty) upd() (bt Bounty, err error) {
  stmt, err := db.Prepare(`
    UPDATE Bounty
    SET Bounty.state = ?, Bounty.title = ?, Bounty.text = ?, Bounty.bounty = ?, Bounty.time_limit = ?, updated_at = CURRENT_TIMESTAMP
    WHERE Bounty.id = ?
  `)
  if err != nil {
    return
  }
  _, err = stmt.Exec(b.State, b.Title, b.Text, b.Bounty, b.Time_limit, b.Id)
  if err != nil {
    return
  }
  defer stmt.Close()
  bt, err = b.sel()
  if err != nil {
    return
  }
  return
}

func SelectBounty(c *gin.Context) {
  var err error
  db_source := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
  db, err = sql.Open("mysql", db_source)

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var b Bounty
  if err := c.ShouldBindJSON(&b); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }
  bt, err := b.sel()
  if err != nil {
    log.Println("Data not Select!")
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

func ListBounty(c *gin.Context) {
  var err error
  db_source := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
  db, err = sql.Open("mysql", db_source)

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var b Bounty
  bts, err := b.selAll()
  if err != nil {
    log.Println("Data not SelectAll!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": err,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "length": len(bts),
    "content": bts,
  })
  return
}

func RegisterBounty(c *gin.Context) {
  var err error
  db_source := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
  db, err = sql.Open("mysql", db_source)

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var b Bounty
  if err := c.ShouldBindJSON(&b); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }

  id, err := b.ins()
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
    "content": id,
  })
}

func UpdateBounty(c *gin.Context) {
  var err error
  db_source := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
  db, err = sql.Open("mysql", db_source)

  if err != nil {
    log.Fatal(err.Error())
  }
  defer db.Close()

  var b Bounty
  if err := c.ShouldBindJSON(&b); err != nil {
    log.Println("Data not Bind!")
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "content": "Data not Bind!",
    })
    return
  }

  bt, err := b.upd()
  if err != nil {
    log.Println("Data update Failed!")
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
}
