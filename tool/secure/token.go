package secure

import (
	"math/rand"
	"sync"
	"time"
)

var (
	lock    sync.Mutex
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// GenerateToken  create Token
func GenerateToken(lenth int) string {
	lock.Lock()
	defer lock.Unlock()
	token := make([]rune, lenth)

	mRandom := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range token {
		token[i] = letters[mRandom.Intn(len(letters))]
	}

	return string(token)
}
