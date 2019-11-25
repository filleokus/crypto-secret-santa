package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/* Creates a MD5 hash
 * Takes a passphrase or any string
  * Returns a hash of hexadecimal value */

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/* First we create a block cipher based on a hashed passphase
 * Then, we need to wrap it in GCM with a nonce
 * Nonce needs to be used when decrypting, so we store it
 * alongside the encrypted data */
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

/* First we create a new block cipher using the HASHED passphrase
 * We wrap it around GCM and get the nonce size
 * Then we need to separate the nonce and the encrypted data */
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

/* Shuffling an slice of strings*/
func Shuffle(array []string, source rand.Source) {
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func printSlice(slice []string) {
	for _, value := range slice {
		fmt.Printf("%d\n", value)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())

	var rsgMembers = []string{"Dave", "Mike", "Kevin", "Branne", "Jens", "Per", "Max", "Jacob"}
	var pool = []string{"Dave", "Mike", "Kevin", "Branne", "Jens", "Per", "Max", "Jacob"}

	var buyer = ""
	var reciever = ""
	var turn = 1

	fmt.Println("--- Shuffling both the lists ...---")
	Shuffle(rsgMembers, source)
	Shuffle(pool, source)
	// DEBUG
	// fmt.Println("--- Printing shuffled members---")
	// printSlice(rsgMembers)
	// fmt.Println("--- Printing shuffled pool ---")
	// printSlice(pool)

	for i := 1; i <= 8; i++ {
		buyer = rsgMembers[0]
		reciever = pool[0]

		//buyer != reciever
		//TODO works sometime lol
		// hacking a do->while
		for ok := true; ok; ok = buyer == reciever {
			// fmt.Println("ERROR!")
			// fmt.Println("Buyer and Reciever are the same!")
			// fmt.Println(buyer + " " + reciever)
			Shuffle(pool, source)
			reciever = pool[0]
			// fmt.Println(buyer + " " + reciever)
		}

		fmt.Println("--- TURN " + strconv.Itoa(turn) + " ---")
		fmt.Println(buyer + " buys for: ")

		/* Encrypting and decrypting TEST*/
		ciphertext := encrypt([]byte(reciever), "rsgchristmas2019")
		fmt.Printf("Encrypted: %x\n", ciphertext)
		plaintext := decrypt(ciphertext, "rsgchristmas2019")
		fmt.Printf("Decrypted: %s\n", plaintext)

		/* Removing both the buyer and reciver from the pools */

		// fmt.Println("Removing " + buyer + " from the list")
		rsgMembers = Remove(rsgMembers, 0)
		// printSlice(rsgMembers)

		// fmt.Println("Removing " + reciever + " from the list")
		pool = Remove(pool, 0)
		// printSlice(pool)
		turn++
	}

}
