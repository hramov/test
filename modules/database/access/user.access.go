package access

import (
	"testing/domains/dto"
	mapper "testing/modules/database/mappers"
	model "testing/modules/database/models"

	"gorm.io/gorm"
)

type UserProvider struct {
	User  *model.User
	Users []*model.User
	DB    *gorm.DB
}

func (up *UserProvider) Get() ([]*dto.UserDto, error) {
	var users []*dto.UserDto
	up.DB.Find(&up.Users)
	for i := 0; i < len(up.Users); i++ {
		um := mapper.UserMapper{Model: *up.Users[i]}
		users = append(users, um.ModelToDto())
	}
	return users, nil
}

func (up *UserProvider) Add(userDto *dto.UserDto) (uint64, error) {
	um := mapper.UserMapper{Dto: *userDto}
	userModel := um.DtoToModel()
	up.DB.Create(userModel)
	return userDto.ID, nil
}

func (up *UserProvider) Delete(id uint64) (uint64, error) {
	up.DB.Delete("id=?", id)
	return id, nil
}
