package main

import (
	"fmt"
)

func main() {

	var checksumtype string

	fmt.Println("Welcome to gochecksum!")
	fmt.Println("WARNING! CONVERTS EVERYTHING TO STRING!")
	fmt.Println("Only supports: MD5, SHA1, SHA256 and SHA512")

	fmt.Print("What checksum would you like to convert to?: ")

	fmt.Scan(&checksumtype)

	switch checksumtype {

	case "MD5", "md5":
		md5_checksum()
		fmt.Println("Don't use MD5! It is recommended to use SHA256 or SHA512.")
	case "SHA1", "sha1":
		sha1_checksum()
		fmt.Println("Don't use SHA1! It is recommended to use SHA256 or SHA512.")
	case "SHA256", "sha256":
		sha256_checksum()
	case "SHA512", "sha512":
		sha512_checksum()
	default:
		fmt.Println("Error reading input")
	}

}
