package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

func BuildSignString(method, path string, body *bytes.Buffer, queryParams url.Values) (sign string) {
	sign = method + "\n"
	sign += Md5(body.Bytes()) + "\n"

	url := path

	keys := make([]string, 0, len(queryParams))
	for k := range queryParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	querys := make([]string, 0)
	for _, k := range keys {
		if k == "sign" {
			continue
		}

		val := k

		if len(queryParams[k]) != 0 {
			val += "=" + queryParams[k][0]
		}

		querys = append(querys, val)
	}

	if len(querys) > 0 {
		url += "?" + strings.Join(querys, "&")
	}

	sign += url

	return
}

func CreateSign(signString, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(signString))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
func VerifySign(sign, secret string) bool {

	return false
}

func Md5(msg []byte) string {
	h := md5.New()
	h.Write(msg)
	return hex.EncodeToString(h.Sum(nil))
}
