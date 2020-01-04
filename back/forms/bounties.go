package forms

type BountySelect struct {
  Id            int `json:"id" binding:"required"`
}

type BountyRegister struct {
  Title      string `json:"title" binding:"required"`
  Text       string `json:"text" binding:"required"`
  Bounty        int `json:"bounty" binding:"required"`
  Time_limit string `json:"time_limit" binding:"required"`
}

type BountyUpdate struct {
  Id            int `json:"id" binding:"required"`
  State         int `json:"state" binding:"required"`
  Title      string `json:"title" binding:"required"`
  Text       string `json:"text" binding:"required"`
  Bounty        int `json:"bounty" binding:"required"`
  Time_limit string `json:"time_limit" binding:"required"`
}
