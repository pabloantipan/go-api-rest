package domain

type Player struct {
	ID       string  `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Position string  `json:"position" binding:"required"`
	Rating   float64 `json:"rating" binding:"required"`
}
