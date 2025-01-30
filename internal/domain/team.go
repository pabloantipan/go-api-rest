package domain

type Team struct {
	ID     string `json:"id" example:"1"`
	UserID string `json:"userId" binding:"required" example:"1"`
	Name   string `json:"name" binding:"required" example:"DoeTeam"`
}
