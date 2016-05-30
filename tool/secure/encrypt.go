package secure

import (
	"encoding/hex"

	"crypto/md5"
)

// doMd5 return md5-ed data
func DoMd5(data string) (encrypto string, err error) {
	md5Hash := md5.New()
	_, err = md5Hash.Write([]byte(data))
	if err != nil {
		return "", err
	}
	hashedBytes := md5Hash.Sum(nil)
	hashedData := hex.EncodeToString(hashedBytes)
	return hashedData, nil
}
