package validation

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode"

	"github.com/rs401/letsgorip/pb"
)

var (
	// ErrEmptyName error for empty name
	ErrEmptyName = errors.New("name cannot be empty")
	// ErrEmptyEmail error for empty email
	ErrEmptyEmail = errors.New("email cannot be empty")
	// ErrEmptyPassword error for empty password
	ErrEmptyPassword = errors.New("password cannot be empty")
	// ErrEmptyTitle error for empty Title
	ErrEmptyTitle = errors.New("title cannot be empty")
	// ErrEmptyDescription error for empty Description
	ErrEmptyDescription = errors.New("description cannot be empty")
	// ErrEmptyMsg error for empty Msg
	ErrEmptyMsg = errors.New("message cannot be empty")
	// ErrInvalidEmail error for invalid email
	ErrInvalidEmail = errors.New("email not valid")
	// ErrEmailExists error for email already exists
	ErrEmailExists = errors.New("email already exists")
	// ErrNameExists error for name already exists
	ErrNameExists = errors.New("name already exists")
	// ErrNotFound error for not found
	ErrNotFound = errors.New("user not found")
	// ErrInvalidPassword error for invalid password
	ErrInvalidPassword = errors.New("invalid password, 8-50 characters, one upper, lower, number and special character")

	maxPwLen int = 50
	minPwLen int = 8
)

// IsValidSignUp takes a *SignUpRequest and verifies if the request is valid
func IsValidSignUp(user *pb.User) error {
	if IsEmptyString(user.Name) {
		return ErrEmptyName
	}
	if IsEmptyString(user.Email) {
		return ErrEmptyEmail
	}
	if IsEmptyString(user.Password) {
		return ErrEmptyPassword
	}
	if !IsValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if !IsValidPassword(user.Password) {
		return ErrInvalidPassword
	}

	return nil
}

func IsValidForum(forum *pb.Forum) error {
	if IsEmptyString(forum.Title) {
		return ErrEmptyTitle
	}
	if IsEmptyString(forum.Description) {
		return ErrEmptyDescription
	}

	return nil
}

// IsEmptyString verifies if a string is empty
func IsEmptyString(in string) bool {
	return strings.TrimSpace(in) == ""
}

// IsValidEmail verifies if an email is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidPassword verifies if a password is valid
func IsValidPassword(s string) bool {
	var (
		isMin   bool
		special bool
		number  bool
		upper   bool
		lower   bool
		errStr  string
	)

	// Check length
	if len(s) < minPwLen || len(s) > maxPwLen {
		isMin = false
		errStr += fmt.Sprintf("password length must be between %d and %d, ", minPwLen, maxPwLen)
	}

	// Check other requirements
	for _, c := range s {
		if special && number && upper && lower && isMin {
			break
		}

		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	// Append error messages
	if !special {
		errStr += "should contain at least a single special character, "
	}
	if !number {
		errStr += "should contain at least a single digit, "
	}
	if !lower {
		errStr += "should contain at least a single lowercase letter, "
	}
	if !upper {
		errStr += "should contain at least single uppercase letter, "
	}

	// If there are any errors
	if len(errStr) > 0 {
		ErrInvalidPassword = errors.New(errStr)
		return false
	}

	// No errors
	return true
}

// NormalizeEmail normalizes email string
func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
