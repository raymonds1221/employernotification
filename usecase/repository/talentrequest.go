package repository

import (
	talentRequestModel "github.com/Ubidy/Ubidy_EmployerNotificationAPI/domain/talentrequest"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// TalentRequest interface for talentrequest repository
type TalentRequest interface {
	AddTalentRequestCancelActivity(talentrequest talentRequestModel.SuccessFeeTalentRequest) (stream.Activity, error)
}
