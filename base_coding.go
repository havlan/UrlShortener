package main

import (
	"fmt"
	"strings"
)

const (
	// https://stackoverflow.com/questions/1119722/base-62-conversion
	BASE     string = "23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	BASE_LEN int    = len(BASE)
)

type base_coder struct {
}

func baseUDecode(inn_str string) (num int) {
	for _, c := range inn_str {
		if strings.Index(BASE, string(c)) > -1 {
			num = num*BASE_LEN + strings.Index(BASE, string(c))
		} else {
			return -1
		}
	}
	return
}

func baseUEncode(inn int) (out_str string) {
	out_str = ""
	if inn == 0 {
		return string(BASE[0])
	}
	for inn != 0 {
		out_str = string(BASE[inn%BASE_LEN]) + out_str
		inn /= BASE_LEN
		fmt.Println(inn)
	}
	return
}
