package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Will create hash password
// It should never panic if plainText is given properly
func CreateHash(plainText string) (hashText string) {
	preparedPlainText := preparePasswordInput(plainText)
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPlainText), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashText = string(passwordHashInBytes)
	return
}

// Compare hash to plain text, if same it will no return error
// If not same, it will return error
func CompareHash(plainText string, hashText string) (err error) {
	preparedPlainText := preparePasswordInput(plainText)
	plainTextInBytes := []byte(preparedPlainText)
	hashTextInBytes := []byte(hashText)
	err = bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	return
}

// Bcrypt truncates strings which are longer than 72 characters.
// This prepares the plainText input, so that more than that can be used.
func preparePasswordInput(plainText string) (preparedPasswordInput string) {
	// Creates a SHA512 hash, trimmed to 64 characters, so that it fits in bcrypt
	hashedInput := sha512.Sum512_256([]byte(plainText))
	// Bcrypt terminates at NULL bytes, so we need to trim these away
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPasswordInput = string(trimmedHash)
	return
}

// This is example for make a hash string and comparing whether plain string is same as hashed string
// This useful when you are working in password.
// For example, you must never try to save password using plain text from user directly in db.
// bcrypt is todays solution to save password in db rather than md5,
// because it will generates wide range combination of hash result
func main() {
	// specify the password
	password := "mypassword"

	for i := 0; i < 10; i++ {
		hashedPassword := CreateHash(password) // should return different random string
		err := CompareHash(password, hashedPassword)
		if err == nil {
			// You'll see that there is no same hash in each iteration
			fmt.Printf("%s is an hash version of %s \n", hashedPassword, password)
		} else {
			fmt.Println(err.Error())
		}
	}

	// Real case
	// For example, get user password from input
	userPassword := "this-is-user-password"
	userPasswordWrong := "wrong-password-input-by-user"
	userPasswordHashed := CreateHash(userPassword)

	// Should return no error, since user input == real password that has been hashed befoe
	if err := CompareHash(userPassword, userPasswordHashed); err == nil {
		fmt.Printf("You have inputed the right password!\n")
	}

	// Will return wrong password, because err is not nil
	// And hashed string (userPasswordHashed) is not one of any version of plain string text (userPasswordWrong)
	if err := CompareHash(userPasswordWrong, userPasswordHashed); err != nil {
		fmt.Printf("You have inputed the wrong password!\n")
	}
}
