package main

import (
	"strings"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	symbols []rune = []rune{' ', '!', '\\', '/', '"', '#', '$', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', ']', '^', '_',
		'`', '{', '|', '}', '~'}
	lowerCaseAlphabet []rune = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'æ', 'ø', 'å'}
	allCases          []rune = func() []rune {
		var allCases []rune
		allCases = append(allCases, lowerCaseAlphabet...)
		for _, char := range lowerCaseAlphabet {
			allCases = append(allCases, []rune(strings.ToUpper(string(char)))...)
		}
		allCases = append(allCases, symbols...)
		return allCases
	}()
	encryptMode = binding.NewBool()
	key         int
)

func main() {
	encryptMode.Set(true)
	myApp := app.New()
	myWindow := myApp.NewWindow("Encrypt / Decrypt container")
	myWindow.Resize(fyne.Size{Width: 800, Height: 500})
	myWindow.SetFixedSize(true)
	myWindow.CenterOnScreen()

	encryptEntry := widget.NewMultiLineEntry()
	textArea := widget.NewMultiLineEntry()
	keyEntry := widget.NewEntry()
	label := widget.NewLabel(func() string {
		if ignoreError(encryptMode.Get()) {
			return "Encrypt"
		}
		return "Decrypt"
	}())
	var but *widget.Button
	updateBut := func() {
		if ignoreError(encryptMode.Get()) {
			but.SetText("I want to decrypt")
			label.SetText("Encrypt")
		} else {
			but.SetText("I want to encrypt")
			label.SetText("Decrypt")
		}
	}
	but = widget.NewButton("I want to decrypt", func() {
		encryptMode.Set(!ignoreError(encryptMode.Get()))
		if ignoreError(encryptMode.Get()) {
			updateBut()
		} else {
			updateBut()
		}
	})

	content := container.New(
		layout.NewVBoxLayout(),
		but,
		// if encrypt mode print encrypt else print decrypt
		label,
		&widget.Form{
			Items: []*widget.FormItem{
				{Text: "Entry", Widget: encryptEntry}, {Text: "Key", Widget: keyEntry}},
			SubmitText: "PRÆZEL",
			OnSubmit: func() {

				key = func() int {
					if keyEntry.Text == "" {
						return 0
					}
					// Convert keyentry.text to int
					keyb, err := strconv.Atoi(keyEntry.Text)
					if err != nil {
						return 0
					}
					return keyb
				}()
				textArea.SetText(rotateText(encryptEntry.Text))
			},
		},
		textArea,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func ignoreError(b bool, e error) bool {
	return b
}

// Takes a string and rotates each character by the provided amount.
func rotateText(inputText string) string {
	var result string
	findFunc := func(list []rune, target rune) (int, bool) {
		for index, char := range list {
			if char == target {
				return index, true
			}
		}
		return -1, false
	}

	for _, char := range inputText {
		rot := key
		if !ignoreError(encryptMode.Get()) {
			rot = -rot
		}
		if i, found := findFunc(allCases, char); found {
			result += string(allCases[modLikePython(i+rot, len(allCases))])
		} else {
			result += string(char)
		}
	}
	return result
}

func modLikePython(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
