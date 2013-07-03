package services

import (
	"github.com/kobeld/duoerl/models/users"
	. "github.com/paulbellamy/mango"
)

const (
	USER_ID    = "duoerl_id"
	LOGIN_USER = "login_user"
)

func IsCurrentUserWithId(env Env, id string) bool {
	if id == FetchUserIdFromSession(env) {
		return true
	}
	return false
}

func PutUserIdToSession(env Env, id string) {
	env.Session()[USER_ID] = id
}

func FetchUserIdFromSession(env Env) (id string) {
	value := env.Session()[USER_ID]
	if value != nil {
		id = value.(string)
	}
	return
}

func PutUserToEnv(env Env, user *users.User) {
	env[LOGIN_USER] = user
}

func FetchUserFromEnv(env Env) (user *users.User) {
	value := env[LOGIN_USER]
	if value != nil {
		user = value.(*users.User)
	}
	return
}

func DeleteUserInSession(env Env) {
	s := env.Session()
	if s[USER_ID] != nil {
		delete(s, USER_ID)
	}

	if env[LOGIN_USER] != nil {
		delete(env, LOGIN_USER)
	}
}
