package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"final-task/internal/validator"
)

func main() {
	// define flags
	passwordsFlag := flag.String("passwords", "", "Comma-separated list of passwords to check")
	minFlag := flag.Int("min", 8, "Minimum password length")
	bannedFlag := flag.String("banned", "", "Comma-separated list of banned passwords")

	flag.Parse()

	// validate passwords input
	if *passwordsFlag == "" {
		fmt.Println("Error: --passwords flag is required")
		os.Exit(1)
	}

	// parse inputs
	passwords := strings.Split(*passwordsFlag, ",")
	banned := strings.Split(*bannedFlag, ",")

	// create PasswordChecker
	checker := validator.New(*minFlag, banned)

	// check each password and print results
	for _, pw := range passwords {
		pw = strings.TrimSpace(pw)
		result := checker.Check(pw)

		if result.OK {
			color.Green("OK: %s", result.Password)
		} else {
			color.Red("FAIL: %s - %s", result.Password, strings.Join(result.Reasons, "; "))
		}
	}
}
