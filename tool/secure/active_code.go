package secure

import (
	"math/rand"
	"sync"
	"time"
)

var (
	lockForCode sync.Mutex
	lettersNum  = []rune("0123456789")
)

// GenerateDigitActiveCode  随机生成数字验证码
func GenerateDigitActiveCode(lenth int) string {
	lockForCode.Lock()
	defer lockForCode.Unlock()
	token := make([]rune, lenth)

	mRandom := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range token {
		token[i] = lettersNum[mRandom.Intn(len(lettersNum))]
	}

	return string(token)
}
