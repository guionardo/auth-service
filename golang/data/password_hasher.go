package data

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"github.com/guionardo/auth-service/golang/setup"
)

type (
	PasswordHasher struct {
		method       string
		hashFunction func(string) string
	}
)

var hasher *PasswordHasher

func init() {
	method := setup.GetConfiguration().PASSWORD_HASH_METHOD

	functions := map[string]*PasswordHasher{
		"md5":    getHasher("md5", md5.New().Sum),
		"sha1":   getHasher("sha1", sha1.New().Sum),
		"sha256": getHasher("sha256", sha256.New().Sum),
	}

	ok := false
	if hasher, ok = functions[method]; !ok {
		keys := make([]string, 0, len(functions))
		for k := range functions {
			keys = append(keys, k)
		}
		log.Fatalf("INVALID PASSWORD_HASH_METOD '%s' EXPECTED '%v'", method, keys)
	}
	log.Printf("PasswordHasher method %s", method)

}

func getHasher(name string, hashFunc func([]byte) []byte) *PasswordHasher {

	return &PasswordHasher{
		method: name,
		hashFunction: func(text string) string {
			data := []byte(text)
			return fmt.Sprintf("%x", hashFunc(data))
		},
	}
}

func GetHasher() *PasswordHasher {
	return hasher
}

func (hasher *PasswordHasher) Hash(text string) string {
	return strings.ToLower(hasher.Hash(text))
}
