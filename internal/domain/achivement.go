package domain

type Achivement struct {
	ID          string `json:"id" example:"1"`
	UserID      string `json:"userId" binding:"required" example:"1"`
	Name        string `json:"name" binding:"required" example:"Climbed a mountain"`
	Description string `json:"description" binding:"required" example:"An achivement"`
	Points      int    `json:"points" binding:"required" example:"200"`
}
