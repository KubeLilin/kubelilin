package dto

type DeliverablesTreeDTO struct {
	Title    string
	Key      string
	Children []DeliverablesTreeDTO
}
