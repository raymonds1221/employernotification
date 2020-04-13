package usecase

import (
	talentRequestModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/talentrequest"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// TalentRequestInteractor interface for talentrequest usecase
type TalentRequestInteractor interface {
	AddTalentRequestCancelActivity(talentrequest talentRequestModel.SuccessFeeTalentRequest) (stream.Activity, error)
}
