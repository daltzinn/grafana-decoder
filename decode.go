package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	users := []string{
		"boris@data.vl|dc6becccbb57d34daf4a4e391d2015d3350c60df3608e9e99b5291e47f3e5cd39d156be220745be3cbe49353e35f53b51da8|LCBhdtJWjl",
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	for _, user := range users {
		userParts := strings.Split(user, "|")
		if len(userParts) != 3 {
			fmt.Println("Invalid user data format")
			continue
		}

		hexEncodedHash := userParts[1]
		saltString := userParts[2]

		hashBytes, err := hex.DecodeString(hexEncodedHash)
		if err != nil {
			fmt.Printf("Error decoding hex hash: %v\n", err)
			continue
		}

		base64EncodedHash := base64.StdEncoding.EncodeToString(hashBytes)
		base64EncodedSalt := base64.StdEncoding.EncodeToString([]byte(saltString))

		line := fmt.Sprintf("sha256:10000:%s:%s\n", base64EncodedSalt, base64EncodedHash)
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}
	}
}

