package forms

type UserSignIn struct {
  Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type UserSignUp struct {
  Nickname string `json:"nickname" binding:"required"`
  Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
  Group    string `json:"group" binding:"required"`
}
