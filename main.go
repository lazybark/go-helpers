package main

import (
	"fmt"

	"github.com/lazybark/go-helpers/hasher"
)

func main() {
	// Get all types of hash
	SHA256, err := hasher.HashFilePath("example_file", hasher.SHA256, 8192)
	if err != nil {
		fmt.Println(err)
	}
	MD5, err := hasher.HashFilePath("example_file", hasher.MD5, 8192)
	if err != nil {
		fmt.Println(err)
	}
	SHA1, err := hasher.HashFilePath("example_file", hasher.SHA1, 8192)
	if err != nil {
		fmt.Println(err)
	}
	SHA512, err := hasher.HashFilePath("example_file", hasher.SHA512, 8192)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File hashes:")
	fmt.Println(SHA256)
	fmt.Println(MD5)
	fmt.Println(SHA1)
	fmt.Println(SHA512)

	fmt.Println("String hashes ('Some string for you'):")
	fmt.Println(hasher.HashString("Some string for you", hasher.SHA256))
	fmt.Println(hasher.HashString("Some string for you", hasher.MD5))
	fmt.Println(hasher.HashString("Some string for you", hasher.SHA1))
	fmt.Println(hasher.HashString("Some string for you", hasher.SHA512))

	fmt.Println("[]byte hashes:")
	b := []byte("Some string for you")
	fmt.Println(hasher.HashBytes(b, hasher.SHA256))
	fmt.Println(hasher.HashBytes(b, hasher.MD5))
	fmt.Println(hasher.HashBytes(b, hasher.SHA1))
	fmt.Println(hasher.HashBytes(b, hasher.SHA512))

	fmt.Println("Checking types (9, 1):")
	// Wring hash type
	wrong := hasher.HashType(9)
	fmt.Println(wrong)

	// Correct hash type
	right := hasher.HashType(1)
	fmt.Println(right)

}
