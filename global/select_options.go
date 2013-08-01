package global

const (
	FROM_STROE  string = "from_01"
	FROM_ONLINE        = "from_02"
)

var GotFromOptions = map[string]string{
	FROM_STROE:  "Store",
	FROM_ONLINE: "Online",
}
