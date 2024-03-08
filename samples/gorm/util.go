package gorm

import (
	"math/rand"
	"strconv"
)

// 指定した長さのランダム文字列を生成
func RandomStr(length int) string {
	const CANDIDATES = "abcdefghijklmnopqrstuvwxyz"
	const LEN_CAND = len(CANDIDATES)
	bytes := make([]byte, 0, length) // 必要な capacity は明確
	for i := 0; i < length; i++ {
		bytes = append(bytes, CANDIDATES[rand.Intn(LEN_CAND)])
	}
	return string(bytes)
}

// 指定した長さのランダム数値を生成
func RandomInt(length int) int {
	const CANDIDATES = "1234567890"
	const LEN_CAND = len(CANDIDATES)
	bytes := make([]byte, 0, length) // 必要な capacity は明確
	for i := 0; i < length; i++ {
		bytes = append(bytes, CANDIDATES[rand.Intn(LEN_CAND)])
	}
	s := string(bytes)
	i, _ := strconv.Atoi(s)
	return i
}
