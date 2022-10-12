package db_impl

import (
	"fmt"

	"infraction-tracker/models"
)

func punisheeIdx(ps []models.Punishee, name string) (int, error) {
	idx := -1
	for i := range ps {
		if ps[i].Name == name {
			idx = i
			break
		}
	}
	if idx == -1 {
		return idx, fmt.Errorf("not found")
	}
	return idx, nil
}

type MemoryDb struct {
	punishees []models.Punishee
}

func (db *MemoryDb) Init() {}
func (db *MemoryDb) List() ([]models.Punishee, error) {
	// make explicit copy to prevent modification outside database
	cp := make([]models.Punishee, len(db.punishees))
	copy(cp, db.punishees)
	return cp, nil
}
func (db *MemoryDb) AddPunishee(p models.Punishee) error {
	_, err := punisheeIdx(db.punishees, p.Name)
	if err == nil {
		return fmt.Errorf("no duplicate names allowed")
	}
	db.punishees = append(db.punishees, p)
	return nil
}
func (db *MemoryDb) UpdatePunishee(p models.Punishee) error {
	idx, err := punisheeIdx(db.punishees, p.Name)
	if err != nil {
		return err
	}

	db.punishees[idx] = p
	return nil
}
func (db *MemoryDb) RemovePunishee(name string) error {
	idx, err := punisheeIdx(db.punishees, name)
	if err != nil {
		return err
	}

	length := len(db.punishees) - 1
	db.punishees[idx] = db.punishees[length]
	db.punishees = db.punishees[:length]

	return nil
}
