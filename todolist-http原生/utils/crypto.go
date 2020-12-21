package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(text string) string {
	//hash := md5.New()
	//fmt.Fprintf(hash, text)
	//return fmt.Sprintf("%x", hash.Sum(nil))
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}
