package main

import (
	"bufio"
	"fmt"
	"os"

	account "auto-pick.com/structs/account"
	data "auto-pick.com/structs/data"
	handler "auto-pick.com/structs/handler"
	constants "auto-pick.com/utils/constants"
	helpers "auto-pick.com/utils/helpers"
	sections "auto-pick.com/utils/sections"
)

func main() {
	// car := structs.Car{
	// 	Name:     "Test",
	// 	Category: "Sedan",
	// }

	scanner := bufio.NewScanner(os.Stdin)

	InitialData := data.Data{
		Account: account.Account{
			FirstName: "Diether",
			LastName:  "Marcelo",
		},
		Scanner: scanner,
	}

	handler := handler.Handler{
		CurrentSection: sections.IntroductionSection,
		InitialData:    &InitialData,
	}

	for handler.CurrentOption != constants.QuitOption {
		handler.CurrentSection.DisplayMessage()
		helpers.PrintOptions(handler.CurrentSection.Options)
		scanner.Scan()
		handler.CurrentOption = scanner.Text()

		CurrentOption := handler.CurrentSection.GetOption(handler.CurrentOption)

		if CurrentOption.Type == "change-section" {
			handler.ChangeSection(CurrentOption.TargetSection)
		} else {
			CurrentOption.CallAction(handler.InitialData)
		}

		// handler.SwitchSection(sections.CarListSection)
		// TODO: quit Option
		fmt.Println("Current option after", handler.CurrentOption)
	}

	// fmt.Print(car)
}
