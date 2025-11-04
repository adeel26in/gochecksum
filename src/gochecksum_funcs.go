package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func md5_checksum() {

	var input string

	fmt.Print("What would you like to convert?: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input!")
		return
	}

	checksum := fmt.Sprintf("%x", md5.Sum([]byte(input)))

	fmt.Println("Checksum:", checksum)
}

func sha1_checksum() {
	var input string

	fmt.Print("What would you like to convert?: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input!")
		return
	}

	checksum := fmt.Sprintf("%x", sha1.Sum([]byte(input)))

	fmt.Println("Checksum:", checksum)
}

func sha256_checksum() {
	var input string

	fmt.Print("What would you like to convert?: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input!")
		return
	}

	checksum := fmt.Sprintf("%x", sha256.Sum256([]byte(input)))

	fmt.Println("Checksum:", checksum)
}

func sha512_checksum() {
	var input string

	fmt.Print("What would you like to convert?: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input!")
		return
	}

	checksum := fmt.Sprintf("%x", sha512.Sum512([]byte(input)))

	fmt.Println("Checksum:", checksum)
}
