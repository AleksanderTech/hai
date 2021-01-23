package common

import "math/rand"

func RandomCode(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, n)

	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
