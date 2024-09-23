package main

import (
	"fmt"
	util "pwWallet/util"
)

func main() {
	pHandler := wakeUp()
}

// Write to std out welcome message
// returns a ready password cipher
func wakeUp() *util.Handler {
	fmt.Println("Welcome to your password Wallet")
	return util.NewHandler()
}

func getExistingPassword(s string) string {
	return ""
}

func getNewPassword(s string) string {
	return ""
}
