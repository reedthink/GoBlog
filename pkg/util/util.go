package util

import "encoding/hex"
import "crypto/md5"

//加密图片名称
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

/**
关于MD5：
虽然MD5已经被破解，无法作为校验了，但是拿来掩盖原照片信息还是可用的。
*/
