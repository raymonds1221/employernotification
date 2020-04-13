package usecase

import (
	talentRequestModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/talentrequest"
	"github.com/Ubidy/Ubidy_EmployerNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// TalentRequestInteractor implementation of talentrequest usecase
type TalentRequestInteractor struct {
	repository repository.TalentRequest
}

// NewTalentRequestInteractor create new instance of talentrequest interactor
func NewTalentRequestInteractor(repository repository.TalentRequest) *TalentRequestInteractor {
	return &TalentRequestInteractor{
		repository: repository,
	}
}

// AddTalentRequestCancelActivity create an activity to agency when employer cancel a talentrequest
func (ti *TalentRequestInteractor) AddTalentRequestCancelActivity(talentrequest talentRequestModel.SuccessFeeTalentRequest) (stream.Activity, error) {
	activity, err := ti.repository.AddTalentRequestCancelActivity(talentrequest)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
