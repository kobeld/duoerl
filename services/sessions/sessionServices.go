package sessions

import (
	. "github.com/paulbellamy/mango"
)

const (
	USER_ID = "id"
)

func PutUserIdToSession(env Env, id string) {
	env.Session()[USER_ID] = id
}
