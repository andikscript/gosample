package util

import (
	"math/rand"
	"time"
)

var number = "0123456789"
var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var word = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomTrxId() string {
	sl := make([]byte, 19)

	for i := 0; i < 19; i++ {
		sl[i] = number[rand.Intn(len(number))]
	}

	return string(sl)
}

func RandomString(length int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	sl := make([]byte, length)

	for i := 0; i < length; i++ {
		sl[i] = letters[random.Intn(len(number))]
	}

	return string(sl)
}

func RandomNumber(length int) string {
	sl := make([]byte, length)

	for i := 0; i < length; i++ {
		sl[i] = number[rand.Intn(len(number))]
	}

	return string(sl)
}

func RandomNumberString(length int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	sl := make([]byte, length)

	for i := 0; i < length; i++ {
		sl[i] = word[random.Intn(len(word))]
	}

	return string(sl)
}
