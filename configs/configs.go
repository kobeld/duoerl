package configs

var (
	AppHost       = ""
	HttpPort      = ":8888"
	DBUrl         = "localhost"
	Database      = "duoerl_p"
	AssetsVersion = 0
)

const (
	EMAIL_REGEXP = `(^([^@\s]+)@((?:[-A-z0-9]+\.)+[A-z]{2,})$)|^$`
)
