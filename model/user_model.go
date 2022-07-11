package model

type UserRegisterPayload struct {
	FullName string `json:"fullName" validate:"required,gte=4,lte=40" label:"fullName"`
	Gender   string `json:"gender" validate:"required,oneof=male female" label:"gender"`
	Email    string `json:"email" validate:"required,gte=5,uniquedb=users email" label:"email"`
	Password string `json:"password" validate:"required,gte=5,lte=18" label:"password"`
}
