package ports

import "testing/domains/dto"

type UserEntityPort interface {
	Get() ([]*dto.UserDto, error)
	Add(loginDto *dto.UserDto) (uint64, error)
	Delete(id uint64) (uint64, error)
}
