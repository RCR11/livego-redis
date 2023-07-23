//Randomizes StreamKeys each server boot - https://github.com/Glamorgann/livego
package uid

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]ruhttps://github.com/Glamorgann/livegone, n)
	// avoid generating same code in order !
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
