package services

import (
	"github.com/kobeld/duoerl/models/accounts"
	"github.com/kobeld/duoerlapi"
)

func toApiAccount(account *accounts.Account) *duoerlapi.Account {
	apiAccount := new(duoerlapi.Account)

	if account != nil {
		apiAccount = &duoerlapi.Account{
			Id:    account.Id.Hex(),
			Name:  account.Name,
			Email: account.Email,
		}
	}

	return apiAccount
}
