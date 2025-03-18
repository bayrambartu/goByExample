package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	data2 := "roof123"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // convert to byte '[]byte'
	fmt.Println("Base64 Encoded:", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Base64 Decoded:", string(sDec))

	sEnc1 := b64.StdEncoding.EncodeToString([]byte(data2))
	fmt.Println("Base64 Encoded:", sEnc1)

	sEnc2, _ := b64.StdEncoding.DecodeString(sEnc1)
	fmt.Println("Base64 Decoded:", string(sEnc2))

	// Normal Base64 (StdEncoding) → uses + and / characters.
	//	- URL-Safe Base64 (URLEncoding) → uses - instead of + and _ instead of /.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("URL Safe Base64 Encoded:", uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("URL Safe Base64 Decoded:", string(uDec))

}
