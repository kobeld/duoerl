package utils

import (
	"encoding/hex"
	"fmt"
	"labix.org/v2/mgo/bson"
	"strings"
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

// ["1", "2", "3"] -> [ObjectId("1"), ObjectId("2"), ObjectId("3")]
func TurnPlainIdsToObjectIds(ids []string) (r []bson.ObjectId) {
	for _, id := range ids {
		if strings.Trim(id, " 　") == "" {
			continue
		}
		oId, err := ToObjectId(id)
		if err != nil {
			continue
		}
		r = append(r, oId)
	}
	return
}

// [ObjectId("1"), ObjectId("2"), ObjectId("3")]-> ["1", "2", "3"]
func TurnObjectIdToPlainIds(ids []bson.ObjectId) (r []string) {
	for _, id := range ids {
		r = append(r, id.Hex())
	}
	return
}

func IsInObjectIds(tragetId bson.ObjectId, ids []bson.ObjectId) bool {
	for _, id := range ids {
		if tragetId == id {
			return true
		}
	}
	return false
}
