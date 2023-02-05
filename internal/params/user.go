package params

import "template/internal/params/generics"

type UserCreateParam struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateParam struct {
	ID    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserListParams struct {
	generics.GenericFilter
}
