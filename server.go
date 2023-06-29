package main

import (
	"fmt"
	"log"
	"mycrypto/aes"
	"mycrypto/util"
	"net/http"
)

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		text1 := r.FormValue("text1")
		text2 := r.FormValue("text2")

		plainText := text1
		key := text2
		plaintextInByte := []byte(plainText)
		util.PKCS7Padding(&plaintextInByte, 16)

		cipher := aes.Encrypto(plaintextInByte, []byte(key))
		fmt.Printf("\n[cipher]: %v\n", cipher)

		response := cipher
		fmt.Fprint(w, response)

		decryptoText := aes.Decrypto(cipher, []byte(key))
		util.PKCS7UnPadding(&decryptoText)
		fmt.Printf("\n[decryptoText]: %s\n", string(decryptoText))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/submit", handleSubmit)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
