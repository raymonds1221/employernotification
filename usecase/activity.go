package usecase

import (
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ActivityInteractor implementation of activity usecase
type ActivityInteractor struct {
	repository repository.Activity
}

// NewActivityInteractor create new instance of activity interactor
func NewActivityInteractor(repository repository.Activity) *ActivityInteractor {
	return &ActivityInteractor{
		repository: repository,
	}
}

// GetActivities get all activities using the clientID
func (ai *ActivityInteractor) GetActivities(clientID string) ([]stream.Activity, error) {
	activities, err := ai.repository.GetActivities(clientID)

	if err != nil {
		return nil, err
	}

	return activities, nil
}
