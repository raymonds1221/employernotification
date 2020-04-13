package repository

import "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/blueprint"

// Blueprint interface for blueprint repository
type Blueprint interface {
	GetBlueprint() (blueprint.Blueprint, error)
}
