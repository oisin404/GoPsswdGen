package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/andlabs/ui"
	"github.com/atotto/clipboard"
)

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Password Generator", 400, 200, false)
		box := ui.NewVerticalBox()
		window.SetChild(box)

		lengthEntry := ui.NewEntry()
		lengthEntry.SetText("16")
		box.Append(ui.NewLabel("Length of the password:"), false)
		box.Append(lengthEntry, false)

		upperCheckbox := ui.NewCheckbox("Include uppercase letters")
		upperCheckbox.SetChecked(true)
		box.Append(upperCheckbox, false)

		numbersCheckbox := ui.NewCheckbox("Include numbers")
		numbersCheckbox.SetChecked(true)
		box.Append(numbersCheckbox, false)

		symbolsCheckbox := ui.NewCheckbox("Include symbols")
		symbolsCheckbox.SetChecked(true)
		box.Append(symbolsCheckbox, false)

		passwordEntry := ui.NewEntry()
		passwordEntry.SetReadOnly(true)
		box.Append(ui.NewLabel("Generated Password:"), false)
		box.Append(passwordEntry, false)

		generateButton := ui.NewButton("Generate Password")
		box.Append(generateButton, false)

		generateButton.OnClicked(func(*ui.Button) {
			length, err := strconv.Atoi(lengthEntry.Text())
			if err != nil {
				ui.MsgBoxError(window, "Error", "Invalid length")
				return
			}
			useUpper := upperCheckbox.Checked()
			useNumbers := numbersCheckbox.Checked()
			useSymbols := symbolsCheckbox.Checked()

			password, err := generatePassword(length, useUpper, useNumbers, useSymbols)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Error generating password")
				return
			}

			passwordEntry.SetText(password)
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		copyButton := ui.NewButton("Copy Password")
		box.Append(copyButton, false)

		copyButton.OnClicked(func(*ui.Button) {
			password := passwordEntry.Text()
			if password == "" {
				ui.MsgBoxError(window, "Error", "No password to copy")
				return
			}
			err := clipboard.WriteAll(password)
			if err != nil {
				ui.MsgBoxError(window, "Error", "Failed to copy password to clipboard")
				return
			}
			ui.MsgBox(window, "Success", "Password copied to clipboard")
		})

		window.Show()
	})

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
}

func generatePassword(length int, useUpper, useNumbers, useSymbols bool) (string, error) {
	const (
		lowerLetters = "abcdefghijklmnopqrstuvwxyz"
		upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers      = "0123456789"
		symbols      = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	)

	var charset string
	charset += lowerLetters
	if useUpper {
		charset += upperLetters
	}
	if useNumbers {
		charset += numbers
	}
	if useSymbols {
		charset += symbols
	}

	if len(charset) == 0 {
		return "", fmt.Errorf("no character sets selected")
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password), nil
}
