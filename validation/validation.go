// Package validation provides methods to validate the various models for this
// project.
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
	ErrNotFound = errors.New("resource not found")
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

// IsValidForum verifies the given forum is valid.
func IsValidForum(forum *pb.Forum) error {
	if IsEmptyString(forum.Title) {
		return ErrEmptyTitle
	}
	if IsEmptyString(forum.Description) {
		return ErrEmptyDescription
	}
	if forum.UserId == 0 {
		// it shouldn't get this far without the UserId
		return fmt.Errorf("user id can't be null")
	}

	return nil
}

// IsValidThread verifies the given thread is valid.
func IsValidThread(thread *pb.Thread) error {
	if IsEmptyString(thread.Title) {
		return ErrEmptyTitle
	}
	if IsEmptyString(thread.Msg) {
		return ErrEmptyMsg
	}
	if thread.UserId == 0 {
		// it shouldn't get this far without the UserId
		return fmt.Errorf("user id can't be null")
	}
	if thread.ForumId == 0 {
		// it shouldn't get this far without the ForumId
		return fmt.Errorf("forum id can't be null")
	}

	return nil
}

// IsValidPost verifies the given post is valid.
func IsValidPost(post *pb.Post) error {
	if IsEmptyString(post.Msg) {
		return ErrEmptyMsg
	}
	if post.UserId == 0 {
		// it shouldn't get this far without the UserId
		return fmt.Errorf("user id can't be null")
	}
	if post.ThreadId == 0 {
		// it shouldn't get this far without the ThreadId
		return fmt.Errorf("thread id can't be null")
	}

	return nil
}

// IsValidPlace verifies the given place is valid.
func IsValidPlace(place *pb.Place) error {
	// Name not empty
	if IsEmptyString(place.Name) {
		return ErrEmptyName
	}
	// Description not empty
	if IsEmptyString(place.Description) {
		return ErrEmptyDescription
	}
	if place.UserId == 0 {
		// it shouldn't get this far without the UserId
		return fmt.Errorf("user id can't be null")
	}
	// Lat & Long shouldn't be zero
	if place.Latitude == 0 || place.Longitude == 0 {
		return fmt.Errorf("coordinates cannot be zero")
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
	special := false
	number := false
	upper := false
	lower := false

	// Check length
	if len(s) < minPwLen || len(s) > maxPwLen {
		return false
	}

	// Check other requirements
	for _, c := range s {
		if special && number && upper && lower {
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

	for _, v := range []bool{special, number, upper, lower} {
		if !v {
			return false
		}
	}

	// No errors
	return true
}

// NormalizeEmail normalizes email string
func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
