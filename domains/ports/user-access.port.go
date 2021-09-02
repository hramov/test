package ports

import "testing/domains/dto"

type UserAccessPort interface {
	Get() ([]*dto.UserDto, error)
	Add(*dto.UserDto) (uint64, error)
	Delete(uint64) (uint64, error)
}
