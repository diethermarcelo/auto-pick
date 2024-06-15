package structs

import (
	data "auto-pick.com/structs/data"
	section "auto-pick.com/structs/section"
	sections "auto-pick.com/utils/sections"
)

type Handler struct {
	CurrentSection section.Section
	CurrentOption  string // TODO: rune
	InitialData    *data.Data
}

func (h *Handler) SwitchSection(new_section section.Section) {
	new_section.Data = h.InitialData
	h.CurrentSection = new_section
}

func (h *Handler) ChangeSection(new_section string) {
	var NewSectionObject section.Section

	switch new_section {
	case "introduction":
		NewSectionObject = sections.IntroductionSection
	case "check-list":
		NewSectionObject = sections.CarListSection
	case "personal-account":
		NewSectionObject = sections.PersonalAccountSection
	default:

	}

	h.SwitchSection(NewSectionObject)

}
