package structs

// import (
// 	action "auto-pick.com/interfaces/action"
// )

import (
	data "auto-pick.com/structs/data"
)

type CallActionHandler func(data *data.Data)

type Option struct {
	OptionKey     rune
	Description   string
	Type          string
	TargetSection string
	CallAction    CallActionHandler
}

func (o *Option) Action() {}

// func New(actionHandler ActionHandlerInterface) *Option {
// 	return &Option{
// 		Action: actionHandler,
// 	}
// }
