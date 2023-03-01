/**
* @description :
* @author : Jarick
* @Date : 2022-01-25
* @Url : http://CloudWebOps
 */

package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// 字符串编码Hash算法：sha256,md5(crypto/md5: weak cryptographic primitive)
func EncrytedStringHandle(data []byte, encrytedType string) (EncrytedString string) {
	if encrytedType == "SHA256" {
		EncrytedString = func(data []byte) string {
			_sha256 := sha256.New()
			_sha256.Write(data)
			return hex.EncodeToString(_sha256.Sum([]byte("")))
		}(data)
	}
	// else if encrytedType == "MD5" {
	// 	EncrytedString = func(data []byte) string {
	// 		_md5 := md5.New()
	// 		_md5.Write(data)
	// 		return hex.EncodeToString(_md5.Sum([]byte("")))
	// 	}(data)
	// }
	return
}
