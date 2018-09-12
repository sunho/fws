package fws

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	rando "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/sunho/fws/server/runtime"
	"github.com/sunho/fws/server/store"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rando.Seed(time.Now().Unix())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rando.Intn(len(letterRunes))]
	}
	return string(b)
}

type fwsInterface struct {
	f *Fws
}

func (f *fwsInterface) GetStore() store.Store {
	return f.f.stor
}

func (f *fwsInterface) GetRunManager() *runtime.RunManager {
	return f.f.runManager
}

func (f *fwsInterface) GetBuildManager() *runtime.BuildManager {
	return f.f.buildManager
}

func (f *fwsInterface) CreateWebhookSecret() string {
	return randomString(20)
}

func (f *fwsInterface) CreateInviteKey(username string) string {
	return randomString(20)
}

func (f *fwsInterface) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (f *fwsInterface) HashPassword(password string) string {
	buf, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		glog.Fatalf("Something went wrong, err: %v", err)
	}
	return string(buf)
}

func (f *fwsInterface) CreateToken(id int, username string) string {
	plainText := []byte(strconv.Itoa(id) + ":" + username)

	block, err := aes.NewCipher([]byte(f.f.config.Secret))
	if err != nil {
		glog.Fatalf("Invalid secret, err: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		glog.Fatalf("Something went wrong, err: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func (f *fwsInterface) ParseToken(tok string) (int, string, bool) {
	cipherText, err := base64.StdEncoding.DecodeString(tok)
	if err != nil {
		return 0, "", false
	}

	block, err := aes.NewCipher([]byte(f.f.config.Secret))
	if err != nil {
		glog.Fatalf("Invalid secret, err: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return 0, "", false
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	str := string(cipherText)
	arr := strings.SplitN(str, ":", 2)
	id, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, "", false
	}
	username := arr[1]

	return id, username, true
}

func (f *fwsInterface) GetDistAddr() string {
	return f.f.config.Dist
}

func (f *fwsInterface) GetDistFolder() http.FileSystem {
	return f.f.dist
}

func (f *fwsInterface) GetIndex() []byte {
	return f.f.index
}
