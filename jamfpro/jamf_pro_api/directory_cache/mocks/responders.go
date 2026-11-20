package mocks

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"

type DirectoryCacheMock struct{ *mocks.GenericMock }

func NewDirectoryCacheMock() *DirectoryCacheMock {
	return &DirectoryCacheMock{GenericMock: mocks.NewJSONMock("DirectoryCacheMock")}
}
