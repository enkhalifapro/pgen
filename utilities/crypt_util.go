package utilities

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

type CryptUtil struct {
}

func (c *CryptUtil) RandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	symbols := big.NewInt(int64(len(alphanum)))
	states := big.NewInt(0)
	states.Exp(symbols, big.NewInt(int64(n)), nil)
	r, err := rand.Int(rand.Reader, states)
	if err != nil {
		panic(err)
	}
	var bytes = make([]byte, n)
	r2 := big.NewInt(0)
	symbol := big.NewInt(0)
	for i := range bytes {
		r2.DivMod(r, symbols, symbol)
		r, r2 = r2, r
		bytes[i] = alphanum[symbol.Int64()]
	}
	return string(bytes)
}

func (c *CryptUtil) Encrypt(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	passwordSha256Hash := hex.EncodeToString(h.Sum(nil))
	return passwordSha256Hash
}

func (c *CryptUtil) Bcrypt(str string) string {
	byteStr := []byte(str)
	hashedPassword, err := bcrypt.GenerateFromPassword(byteStr, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (c *CryptUtil) CompareHashAndPassword(hashedPassword string, password string) bool {
	hashedPasswordByte := []byte(hashedPassword)
	passwordByte := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, passwordByte)
	if err == nil {
		return true
	}
	fmt.Println(err)
	return false
}

func (c *CryptUtil) NewEncryptedToken() string {
	randomStr := c.RandomString(100)
	token := c.Encrypt(randomStr)
	return token
}

func (c *CryptUtil) GenerateShortId() string {
	shortIDAlpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	rand.Read(b)
	s := []byte{}
	for _, v := range b {
		s = append(s, shortIDAlpha[v%62])
	}
	return string(s)
}
