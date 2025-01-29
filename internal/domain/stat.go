package domain

type Stat struct {
	ID     string `json:"id" example:"1"`
	UserID string `json:"userId" binding:"required" example:"1"`
	Points int    `json:"points" binding:"required" example:"200"`
}
