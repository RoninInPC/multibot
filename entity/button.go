package entity

type Button struct {
	Handler  string
	Function func(update Update)
}

type UsesButton struct {
	Handler string
	Buttons Rows
}

type Row []Button

type Rows []Row
