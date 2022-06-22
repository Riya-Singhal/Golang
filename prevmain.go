package prevmain

import (
	"encoding/hex"
	"fmt"
	"github.com/capitalone/fpe/ff1"
)

// panic(err) is just used for example purposes.
func Prevmn() {
	// Key and tweak should be byte arrays. Put your key and tweak here.
	// To make it easier for demo purposes, decode from a hex string here.
	key, err1 := hex.DecodeString("EF4359D8D580AA4F7F036D6F04FC6A94")
	if err1 != nil {
		panic(err1)
	}
	tweak, err1 := hex.DecodeString("D8E7920AFA330A73")
	if err1 != nil {
		panic(err1)
	}

// 	// Create a new FF1 cipher "object"
// 	// 10 is the radix/base, and 8 is the tweak length.
// 	FF1, err1 := ff1.NewCipher(10, 8, key, tweak)
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	original := "123456789"

// 	// Call the encryption function on an example SSN
// 	ciphertext, err1 := FF1.Encrypt(original)
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	plaintext, err1 := FF1.Decrypt(ciphertext)
// 	if err1 != nil {
// 		panic(err1)
// 	}

// 	fmt.Println("Original:", original)
// 	fmt.Println("Ciphertext:", ciphertext)
// 	fmt.Println("Plaintext:", plaintext)
}
