package usecase

import "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/blueprint"

// BlueprintInteractor interface for Blueprint usecase
type BlueprintInteractor interface {
	GetBlueprint() (blueprint.Blueprint, error)
}
