package models

import (
  "../db"
  "../forms"
)

type Bounty struct {
  Id            int `json:"id"`
  State         int `json:"state"`
  Title      string `json:"title"`
  Text       string `json:"text"`
  Bounty        int `json:"bounty"`
  Time_limit string `json:"time_limit"`
}

func (b Bounty) ListAll() (bts []Bounty, err error) {
  db := db.GetDB()
  bt := Bounty{}

  rows, err := db.Query(`
    SELECT Bounty.id, Bounty.state, Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit
    FROM Bounty
  `)
  if err != nil {
    return
  }
  defer rows.Close()

  for rows.Next() {
    err = rows.Scan(&bt.Id, &bt.State, &bt.Title, &bt.Text, &bt.Bounty, &bt.Time_limit)
    if err != nil {
      return
    }
    bts = append(bts, bt)
  }
  return
}

func (b Bounty) Select(btPayload forms.BountySelect) (bt Bounty, err error) {
  db := db.GetDB()
  bt = Bounty{
    Id: btPayload.Id,
  }

  row := db.QueryRow(`
    SELECT Bounty.state, Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit
    FROM Bounty
    WHERE Bounty.id = ?
  `, bt.Id)
  err = row.Scan(&bt.State, &bt.Title, &bt.Text, &bt.Bounty, &bt.Time_limit)
  if err != nil {
    return
  }
  return
}

func (b Bounty) Register(btPayload forms.BountyRegister) (bt Bounty, err error) {
  db := db.GetDB()
  bt = Bounty{
    Title: btPayload.Title,
    Text: btPayload.Text,
    Bounty: btPayload.Bounty,
    Time_limit: btPayload.Time_limit,
  }

  stmt, err := db.Prepare(`
    INSERT INTO Bounty(Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit)
    VALUES(?, ?, ?, ?)
  `)
  if err != nil {
    return
  }
  rs, err := stmt.Exec(bt.Title, bt.Text, bt.Bounty, bt.Time_limit)
  if err != nil {
    return
  }
  id, err := rs.LastInsertId()
  if err != nil {
    return
  }
  defer stmt.Close()

  row := db.QueryRow(`
    SELECT Bounty.id, Bounty.state, Bounty.title, Bounty.text, Bounty.bounty, Bounty.time_limit
    FROM Bounty
    WHERE Bounty.id = ?
  `, id)
  err = row.Scan(&bt.Id, &bt.State, &bt.Title, &bt.Text, &bt.Bounty, &bt.Time_limit)
  if err != nil {
    return
  }
  return
}

func (b Bounty) Update(btPayload forms.BountyUpdate) (bt Bounty, err error) {
  db := db.GetDB()
  bt = Bounty{
    Id: btPayload.Id,
    State: btPayload.State,
    Title: btPayload.Title,
    Text: btPayload.Text,
    Bounty: btPayload.Bounty,
    Time_limit: btPayload.Time_limit,
  }

  stmt, err := db.Prepare(`
    UPDATE Bounty
    SET Bounty.state = ?, Bounty.title = ?, Bounty.text = ?, Bounty.bounty = ?, Bounty.time_limit = ?, Bounty.updated_at = CURRENT_TIMESTAMP
    WHERE Bounty.id = ?
  `)
  if err != nil {
    return
  }
  _, err = stmt.Exec(bt.State, bt.Title, bt.Text, bt.Bounty, bt.Time_limit, bt.Id)
  if err != nil {
    return
  }
  defer stmt.Close()
  return
}
