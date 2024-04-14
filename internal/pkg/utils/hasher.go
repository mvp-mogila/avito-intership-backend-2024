package utils

import (
	"crypto/md5"
	"strconv"
)

func MakeUniqueHash(featID, tagID int) string {
	hashVal := 31 * featID
	hashVal += 31 * tagID
	hasher := md5.New()
	hasher.Write([]byte(strconv.Itoa(hashVal)))
	return string(hasher.Sum(nil))
}
