package auth

import (
	"lambda/types"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword takes a RegisterUser struct and generates a hashed password for the user.
// It returns a pointer to a User struct with the hashed password or an error if the hashing fails.
func GeneratePassword(registerUser types.RegisterUser) (*types.User, error) {
	// Generate a bcrypt hash of the password.
	hash, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Return a new User struct with the username, email, and hashed password.
	return &types.User{
		Username:     registerUser.Username,
		Email:        registerUser.Email,
		PasswordHash: string(hash),
	}, nil
}

// ValidatePassword checks if the provided password matches the hashed password.
// It returns true if the password is correct, otherwise false.
func ValidatePassword(hashedPassword, password string) bool {
	// Compare the plaintext password with the hashed password.
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
