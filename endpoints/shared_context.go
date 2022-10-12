package endpoints

import "infraction-tracker/models"

type SharedContext struct {
	Db models.Db
}
