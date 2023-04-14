package service

import "github.com/joinusordie/app_todo/pkg/repository"

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
