package bot

import "multibot/entity"

type Button struct {
}

type Bot interface {
	NewButtonFunctional(button entity.Button)
	NewButtonSendButtons()
}
