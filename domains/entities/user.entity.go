package entity

import (
	"testing/domains/dto"
	"testing/domains/ports"
)

type UserEntity struct {
	Provider ports.UserAccessPort
}

func (ue *UserEntity) Get() ([]*dto.UserDto, error) {
	orders, err := ue.Provider.Get()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (ue *UserEntity) Add(userDto *dto.UserDto) (uint64, error) {
	id, err := ue.Provider.Add(userDto)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ue *UserEntity) Delete(id uint64) (uint64, error) {
	id, err := ue.Provider.Delete(id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
