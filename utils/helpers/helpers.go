package utils

import (
	"fmt"

	option "auto-pick.com/structs/option"
	section "auto-pick.com/structs/section"
)

func SwitchSection(current_section section.Section, new_section section.Section) {
	current_section = new_section
}

func PrintOptions(options []option.Option) {
	for _, option := range options {
		formatted := fmt.Sprintf("%c.) %s \n", option.OptionKey, option.Description)
		fmt.Print(formatted)
	}
}
