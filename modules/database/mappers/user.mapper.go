package mapper

import (
	"testing/domains/dto"
	model "testing/modules/database/models"
)

type UserMapper struct {
	Dto   dto.UserDto
	Model model.User
}

func (um *UserMapper) DtoToModel() *model.User {
	user := um.Model
	user.ID = um.Dto.ID
	user.Email = um.Dto.Email
	user.Password = um.Dto.Password
	return &user
}

func (um *UserMapper) ModelToDto() *dto.UserDto {
	user := um.Dto
	user.ID = um.Model.ID
	user.Email = um.Model.Email
	user.Password = um.Model.Password
	return &user
}
