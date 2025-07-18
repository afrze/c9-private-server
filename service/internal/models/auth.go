package models

type RegisterDTO struct {
	PAccId    string `json:"pAccId" binding:"required,min=4,max=20,alphanum"`
	PPassword string `json:"pPassword" binding:"required,min=6,max=20"`
}
