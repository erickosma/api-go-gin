package domain

type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email"  binding:"required,email"`
	WhatsApp string `form:"whatsapp" json:"whatsapp" binding:"required"`
}
