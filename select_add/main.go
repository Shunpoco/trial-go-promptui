package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	items := []string{"Vim", "Emacs", "Subline", "VSCode", "Atom"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "What's your text editor",
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()
		if index == -1 {
			// items = append(items, result)
			break
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %s\n", result)
}
