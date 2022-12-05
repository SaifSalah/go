package main

import "fmt"

func main() { 

	input := []byte("super secret message :|") 
	key := []byte("keybecede")
	encoded := Encoder(input, key)
	fmt.Println("Encoded:", encoded)

	decoded := Decoder(encoded, key) 
	fmt.Println("Decoded:", string(decoded))

}

func Encoder(input []byte, key []byte) []byte { 
	result := make([]byte, len(input)) 
	for i := 0; i < len(input); i++ {
		 result[i] = input[i] ^ key[i%len(key)]
		  } 
	return result
 }

 func Decoder(input []byte, key []byte) []byte {
	 result := make([]byte, len(input)) 
	 for i := 0; i < len(input); i++ {
		 result[i] = input[i] ^ key[i%len(key)]
		}
		   return result
}