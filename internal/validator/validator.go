package validator

import (
	"strings"
	"unicode"
)

type Result struct {
	Password string
	OK       bool
	Reasons  []string
}

// for checking passwords
type Checker interface {
	Check(password string) Result
}

// password validation
type PasswordChecker struct {
	minLength int
	blacklist map[string]bool
}

// makes new PasswordChecker with the given minimum length and banned passwords
func New(minLength int, banned []string) *PasswordChecker {
	bl := make(map[string]bool)
	for _, word := range banned {
		bl[word] = true
	}
	return &PasswordChecker{
		minLength: minLength,
		blacklist: bl,
	}
}

// validates a password against multiple criteria
func (pc *PasswordChecker) Check(password string) Result {
	var reasons []string

	if len(password) < pc.minLength {
		reasons = append(reasons, "too short")
	}

	if pc.blacklist[password] {
		reasons = append(reasons, "password is banned")
	}

	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSymbol = true
		}
	}

	if !hasUpper {
		reasons = append(reasons, "missing uppercase letter")
	}
	if !hasLower {
		reasons = append(reasons, "missing lowercase letter")
	}
	if !hasDigit {
		reasons = append(reasons, "missing digit")
	}
	if !hasSymbol {
		reasons = append(reasons, "missing symbol")
	}

	return Result{
		Password: password,
		OK:       len(reasons) == 0,
		Reasons:  reasons
	}
}
