package mocks

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"

type ActionLogsMock struct{ *mocks.GenericMock }

func NewActionLogsMock() *ActionLogsMock {
	return &ActionLogsMock{GenericMock: mocks.NewJSONMock("ActionLogsMock")}
}
