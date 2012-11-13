package sessions

import (
	"github.com/kobeld/duoerl/models/accounts"
	. "github.com/paulbellamy/mango"
)

const (
	ACCOUNT_ID    = "duoerl_id"
	LOGIN_ACCOUNT = "login_account"
)

func PutAccountToSession(env Env, account *accounts.Account) {
	env.Session()[LOGIN_ACCOUNT] = account
}

func PutAccountIdToSession(env Env, id string) {
	env.Session()[ACCOUNT_ID] = id
}

func FetchAccountIdFromSession(env Env) (string, bool) {
	value, exist := env.Session()[ACCOUNT_ID]
	return value.(string), exist
}
