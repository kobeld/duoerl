package middlewares

import (
	. "github.com/paulbellamy/mango"
)

func RespondHtml() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		status, headers, body = app(env)
		if headers == nil {
			headers = Headers{}
		}
		if headers.Get("Content-Type") == "" {
			headers.Add("Content-Type", "text/html; charset=utf8")
		}

		if status == 0 {
			status = 200
		}
		return

	}
}

func RespondJson() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		status, headers, body = app(env)
		if headers == nil {
			headers = Headers{}
		}

		headers.Set("Content-Type", "application/json")

		if status == 0 {
			status = 200
		}

		return

	}

}
