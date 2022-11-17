package commons

import (
	"crypto/md5"
	"fmt"
)

func GetStringMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
