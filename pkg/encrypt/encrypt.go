package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"main/dao"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewSalt() string {
	x := rand.Intn(len(letters) - 10)
	s := make([]byte, 10)
	for i := 0; i < 10; i++ {
		s[i] = letters[x+i]
	}
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return string(s)
}

func PwdToHash(pwd string) dao.Password {
	s := NewSalt()
	h1 := md5.New()
	h2 := sha256.New()
	_, _ = h1.Write([]byte(pwd + s))
	_, _ = h2.Write([]byte(pwd + s))
	return dao.Password{
		MD5Hash:    string(h1.Sum(nil)),
		SHA256Hash: string(h2.Sum(nil)),
		Salt:       s,
	}
}
