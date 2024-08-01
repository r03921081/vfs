package service

import "r03921081/vfs/repository"

func init() {
	UserService = NewUserService()
}

var (
	Register = repository.UserRepository.Register
)
