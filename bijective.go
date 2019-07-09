package main

import (
	"strings"
)

var dictionary = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var arr []string
var base int

func init() {
	arr = strings.Split(dictionary, "")
	base = len(arr)
}

func Encode(i int) string {
	if i == 0 {
		return arr[0]
	}
	var result []string
	for i > 0 {
		result = append(result, arr[i%base])
		i = int(i / base)
	}
	for from, to := 0, len(result)-1; from < to; from, to = from+1, to-1 {
		result[from], result[to] = result[to], result[from] // 交换
	}
	return strings.Join(result, "")
}

func Decode(code string) int {
	i := 0
	input := strings.Split(code, "")
	for _, char := range input {
		for index, value := range arr {
			if value == char {
				i = i * base + index
				break
			}
		}
	}
	return i
}
