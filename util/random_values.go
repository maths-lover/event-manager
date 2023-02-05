package util

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums     = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPhoneNum() string {
	var sb strings.Builder
	k := len(nums)

	for i := 0; i < 10; i++ {
		c := nums[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomDecNum() float64 {
	var sb strings.Builder
	k := len(nums)

	// build integer part
	for i := 0; i < 3; i++ {
		c := nums[rand.Intn(k)]
		sb.WriteByte(c)
	}

	sb.WriteByte('.')

	// build fractional part
	for i := 0; i < 2; i++ {
		c := nums[rand.Intn(k)]
		sb.WriteByte(c)
	}

	num, err := strconv.ParseFloat(sb.String(), 64)
	if err != nil {
		log.Fatal("Couldn't generate a floating number:", err)
	}

	return num
}
