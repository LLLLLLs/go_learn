// Time        : 2019/10/17
// Description :

package service

import (
	"errors"
	"fmt"
	"golearn/pkgtest/gomock"
)

type Service struct {
	repo gomock.Repository
}

func (s Service) CreateTest(key string, value interface{}) error {
	// do something
	err := s.repo.Create(key, value)
	if err != nil {
		return err
	}
	// do something
	return nil
}

var ErrKeyNotExist = errors.New("key not exist")

func (s Service) UpdateGetTest(key string, value interface{}) error {
	v, ok := s.repo.Get(key)
	if !ok {
		// ...
		return ErrKeyNotExist
	}
	fmt.Println("get key", key, "success,value =", v)
	// do something
	err := s.repo.Update(key, value)
	if err != nil {
		return err
	}
	// do something
	return nil
}

func (s Service) RemoveTest(key string) error {
	// do something
	err := s.repo.Remove(key)
	if err != nil {
		return err
	}
	// do something
	return nil
}
