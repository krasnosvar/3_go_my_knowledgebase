package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	enc := base64.StdEncoding.EncodeToString([]byte("secret"))
	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(enc, string(dec))
}
