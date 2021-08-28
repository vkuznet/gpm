package main

import (
	"math/rand"
	"time"
)

// code is based on https://github.com/AlanBar13/pass-generator

const voc string = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers string = "0123456789"
const symbols string = "!@#$%&*+_-="

func init() {
	rand.Seed(time.Now().UnixNano())
}

// helper function to copy password to clipboard
func copyPassword() {
}

func createPassword(length int, hasNumbers bool, hasSymbols bool) string {
	chars := voc
	if hasNumbers {
		chars = chars + numbers
	}
	if hasSymbols {
		chars = chars + symbols
	}
	return generatePassword(length, chars)
}

func generatePassword(length int, chars string) string {
	password := ""
	for i := 0; i < length; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}
	return password
}
