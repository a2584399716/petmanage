package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
const secret = "authorChen"

func main() {
	fmt.Println(encryptPassword("admin"))
}

func encryptPassword(oldPassword string) string{
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}