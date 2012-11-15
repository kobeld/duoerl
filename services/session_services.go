package services

import (
	"github.com/kobeld/duoerl/models/accounts"
	. "github.com/paulbellamy/mango"
)

const (
	ACCOUNT_ID    = "duoerl_id"
	LOGIN_ACCOUNT = "login_account"
)

func IsCurrentAccountWithId(env Env, id string) bool {
	if id == FetchAccountIdFromSession(env) {
		return true
	}
	return false
}

func PutAccountIdToSession(env Env, id string) {
	env.Session()[ACCOUNT_ID] = id
}

func FetchAccountIdFromSession(env Env) (id string) {
	value := env.Session()[ACCOUNT_ID]
	if value != nil {
		id = value.(string)
	}
	return
}

func PutAccountToEnv(env Env, account *accounts.Account) {
	env[LOGIN_ACCOUNT] = account
}

func FetchAccountFromEnv(env Env) (account *accounts.Account) {
	value := env[LOGIN_ACCOUNT]
	if value != nil {
		account = value.(*accounts.Account)
	}
	return
}

func DeleteAccountInSession(env Env) {
	s := env.Session()
	if s[ACCOUNT_ID] != nil {
		delete(s, ACCOUNT_ID)
	}

	if env[LOGIN_ACCOUNT] != nil {
		delete(env, LOGIN_ACCOUNT)
	}
}
