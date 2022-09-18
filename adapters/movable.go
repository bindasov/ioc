package adapters

import "github.com/bindasov/ioc/models"

type MovableAdapter interface {
	GetPosition() *models.Vector
	GetVelocity() *models.Vector
	SetPosition(*models.Vector)
}
