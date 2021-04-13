package user

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/bububa/kwai-openapi/core"
	"github.com/bububa/kwai-openapi/model/user"
)

func GetUserPhone(clt *core.Client, accessToken string) (string, error) {
	req := &user.GetUserPhoneRequest{
		AppID:       clt.AppID,
		AccessToken: accessToken,
	}
	var ret user.GetUserPhoneResponse
	err := clt.Get(req, &ret)
	if err != nil {
		return "", err
	}
	return ret.EncryptedPhone, nil
}

func DecryptUserPhone(clt *core.Client, raw string) (string, error) {
	arr := strings.Split(raw, ":")
	if len(arr) != 2 {
		return "", errors.New("invalid raw data")
	}
	iv, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		return "", err
	}
	secret, err := base64.StdEncoding.DecodeString(clt.Secret)
	if err != nil {
		return "", err
	}
	src, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return "", err
	}
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err = aes.NewCipher(secret)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(src))
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return string(decrypted), nil
}
