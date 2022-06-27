package prevmaine


import (
   "encoding/hex"
   "fmt"
   "github.com/capitalone/fpe/ff1"
   "strings"
   //"crypto/rand"
   "math/rand"
   "time"
)
func hypen_removal(pt string) ([]string,[]string ) { //takes input and removes special characters
    ct := make([]string,len(pt))
    original := make([]string,len(pt))

    j:=0
   for index,s := range(pt) {
      // fmt.Println(index)
      if s>=48 && s<=57 {
          original[j]=string(s)
          j+=1
      } else {
          ct[index]=string(s)
      }
   }
   return ct, original
}
func add_hypen(ct []string,original string) string { // adds the special characters and returns tokensied
    j:=0
   for i := range(ct) {
       if ct[i]=="" {
           ct[i]=string(original[j])
           j+=1
       }
   }
   return strings.Join(ct,"")
}

func randToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}


// panic(err) is just used for example purposes.
func Prevmain() {
	// Key and tweak should be byte arrays. Put your key and tweak here.
	// To make it easier for demo purposes, decode from a hex string here.
	key, err := hex.DecodeString("EF4359D8D580AA4F7F036D6F04FC6A94")
       if err != nil {
          panic(err)
       }

        rand.Seed(time.Now().UnixNano())
            min := 1
            max := 30

        size_of_tweak:=rand.Intn(max - min + 1) + min
        fmt.Println("Tweak length is :",size_of_tweak)
        token, _ := randToken(size_of_tweak)
        fmt.Println("Tweak is :",token)
        //fmt.Println(token)
       tweak, err := hex.DecodeString(token)
       if err != nil {
          panic(err)
       }




       // Create a new FF1 cipher "object"
       // 10 is the radix/base, and 8 is the tweak length.
       FF1, err := ff1.NewCipher(10,size_of_tweak, key, tweak)
       if err != nil {
          panic(err)
       }

       original := "123-456-789"
       fmt.Println("Original text is : ",original)
        ct,withouthypen:=hypen_removal(original)
        withouthypen2:=strings.Join(withouthypen ,"")
        fmt.Println("original text with removed hyphens is : ",withouthypen2)
        //withouthypen2=string.join(withouthypen)
       // Call the encryption function on an example SSN
       ciphertext, err := FF1.Encrypt(withouthypen2)
       if err != nil {
          panic(err)
       }
        res:=add_hypen(ct,ciphertext)
        fmt.Println("the result after encryption is : ",res)




       plaintext, err := FF1.Decrypt(ciphertext)
       if err != nil {
          panic(err)
       }
      //  res:=add_hypen(ct,plaintext)
    //     fmt.Println("Original:", original)
    //     fmt.Println("Ciphertext:", ciphertext)
       fmt.Println("Plaintext after decryption is : ", plaintext)
}
