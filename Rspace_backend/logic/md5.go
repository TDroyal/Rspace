package logic

import (
	"crypto/md5"
	"fmt"
)

// 对str进行md5加密  返回32位的加密字符串
func GetMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
