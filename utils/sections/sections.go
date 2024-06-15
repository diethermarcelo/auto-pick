package utils

import (
	"fmt"

	data "auto-pick.com/structs/data"
	option "auto-pick.com/structs/option"
	section "auto-pick.com/structs/section"
	constants "auto-pick.com/utils/constants"
)

var IntroductionSection section.Section = section.Section{
	Title: "test",
	Options: []option.Option{
		option.Option{
			OptionKey:     'a',
			Description:   constants.CHECK_CAR_LIST,
			Type:          "change-section",
			TargetSection: "check-list",
		},
		option.Option{
			OptionKey:     'b',
			Description:   constants.CHECK_PERSONAL_ACCOUNT,
			Type:          "change-section",
			TargetSection: "personal-account",
		},
	},
	DisplaySectionMessage: func(data *data.Data) {
		fmt.Println("Welcome in Auto Pick App")
	},
}

var CarListSection section.Section = section.Section{
	Title: "car-list",
	Options: []option.Option{
		option.Option{
			OptionKey: 'a',
		},
	},
	DisplaySectionMessage: func(data *data.Data) {
		fmt.Println("Select Cars")
	},
}

var CarSection section.Section

var PersonalAccountSection section.Section = section.Section{
	Title: "personal-account",
	Options: []option.Option{
		option.Option{
			OptionKey:   'a',
			Description: "change first name",
			CallAction: func(data *data.Data) {
				fmt.Print("Please enter the first name: ")
				data.Scanner.Scan()
				changeName := data.Scanner.Text()
				data.Account.FirstName = changeName
				fmt.Println("Changed first name successfully")
			},
		},
		option.Option{
			OptionKey:   'b',
			Description: "change last name",
		},
		option.Option{
			OptionKey:   'c',
			Description: "change age",
		},
		option.Option{
			OptionKey:   'd',
			Description: "check bank account",
		},
		option.Option{
			OptionKey:     'e',
			Description:   constants.CHECK_CAR_LIST,
			Type:          "change-section",
			TargetSection: "check-list",
		},
	},
	DisplaySectionMessage: func(data *data.Data) {
		welcomeMessage := fmt.Sprintf("Hi %s!", data.Account.FirstName)

		fmt.Println(welcomeMessage)
	},
}
