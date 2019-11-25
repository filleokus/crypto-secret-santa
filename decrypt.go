package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"reflect"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {
	/* The input needed for ciphertext is byte[]*/
	fmt.Println("Enter your ciphertext")
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		log.Printf("Failed to read: %v", scanner.Err())
		return
	}
	ciphertext := scanner.Text()
	//TODO fix the types
	fmt.Println(ciphertext)
	fmt.Println(reflect.TypeOf(ciphertext))
	plaintext := decrypt(ciphertext, "rsgchristmas2019")
	fmt.Printf("Decrypted: %s\n", plaintext)

}
