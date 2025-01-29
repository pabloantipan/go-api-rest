package domain

type Client struct {
	ID      string `json:"id" example:"1"`
	Name    string `json:"name" binding:"required" example:"DoeCompany"`
	Rol     string `json:"rol" binding:"required" example:"12345678-9"`
	Email   string `json:"email" binding:"required" example:"hello@doe.cl"`
	Phone   string `json:"phone" binding:"required" example:"+56912345678"`
	Address string `json:"address" binding:"required" example:"Av. Doe 1234"`
}
