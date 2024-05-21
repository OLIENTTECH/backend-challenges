package password

import (
	"crypto/rand"

	"github.com/OLIENTTECH/backend-challenges/pkg/log"
)

func MakeRandomStr() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		log.Error("unexpected error...", log.Ferror(err))
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result
}
