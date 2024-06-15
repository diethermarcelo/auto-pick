package structs

import (
	data "auto-pick.com/structs/data"
	option "auto-pick.com/structs/option"
)

type Section struct {
	Title                 string
	Options               []option.Option
	DisplaySectionMessage func(*data.Data)
	Data                  *data.Data
}

func (s Section) GetOption(ChoseOptionKey string) option.Option {
	var SelectedOption option.Option
	for _, option := range s.Options {
		if string(option.OptionKey) == ChoseOptionKey {
			SelectedOption = option
		}
	}

	return SelectedOption
}

func (s *Section) RegisterData(data *data.Data) {
	s.Data = data
}

func (s *Section) DisplayMessage() {
	s.DisplaySectionMessage(s.Data)
}
