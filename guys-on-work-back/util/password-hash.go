package util

// Imports
import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// Method takes a password as a string and returns the hashed password as a string along with a possible error
func HashPassword(password string) (string, error) {

    // Generate a hash from the password and the provided cost (number of iterations)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

    // Checking there was any error during the hash generation
    if err != nil {

        // If there is an error, return a formatted error message with additional information
        return "", fmt.Errorf("Failed to encrypt the password: %w", err)
    }

    // Return the hashed password as a string and null
    return string(hashedPassword), nil
}
