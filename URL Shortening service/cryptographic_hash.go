package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

func main() {
	var input string
	fmt.Scanln(&input)

	mdaHash := md5.Sum([]byte(input))
	fmt.Printf("md5 gives %T result which is 128 (8*16) bits.\n", mdaHash)
	fmt.Printf("md5 hexadecimal result: %x\n", mdaHash)
	fmt.Printf("md5 binary result: %b\n", mdaHash)

	fmt.Println("-------------------------------------------------------------------------------------------------")

	base64Encoded := base64.StdEncoding.EncodeToString(mdaHash[:])
	base32Encoded := base32.StdEncoding.EncodeToString(mdaHash[:])
	fmt.Printf("base-64 encoding of mda5 with len %v (which is near 128 / 6): %v\n", len(base64Encoded), base64Encoded)
	fmt.Printf("base-32 encoding of mda5 with len %v (which is near 128 / 5): %v\n", len(base32Encoded), base32Encoded)

	fmt.Println("-------------------------------------------------------------------------------------------------")

	sha256Hash := sha256.Sum256([]byte(input))
	fmt.Printf("sha-256 gives %T result which is 256 (8*32) bits.\n", sha256Hash)
	fmt.Printf("sha-256 raw result: %v\n", sha256Hash)
	fmt.Printf("sha-256 hexadecimal result: %x\n", sha256Hash)
	fmt.Printf("sha-256 binary result: %b\n", sha256Hash)

	fmt.Println("-------------------------------------------------------------------------------------------------")

	base64Encoded = base64.StdEncoding.EncodeToString(sha256Hash[:])
	base32Encoded = base32.StdEncoding.EncodeToString(sha256Hash[:])
	fmt.Printf("base-64 encoding of sha-256 with len %v (which is near 256 / 6): %v\n", len(base64Encoded), base64Encoded)
	fmt.Printf("base-32 encoding of sha-256 with len %v (which is near 256 / 5): %v\n", len(base32Encoded), base32Encoded)
}
