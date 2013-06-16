package utils

import (
	"encoding/hex"
	"fmt"
	"labix.org/v2/mgo/bson"
)

func ToObjectId(id string) (bid bson.ObjectId, err error) {
	var d []byte
	d, err = hex.DecodeString(id)
	if err != nil || len(d) != 12 {
		err = fmt.Errorf("Invalid input to ObjectIdHex: %q", id)
		return
	}
	bid = bson.ObjectId(d)
	return
}
