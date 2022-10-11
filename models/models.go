package models

type Punishee struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Db interface {
	Init()
	List() ([]Punishee, error)
	AddPunishee(Punishee) error
	UpdatePunishee(Punishee) error
	RemovePunishee(string) error
}
