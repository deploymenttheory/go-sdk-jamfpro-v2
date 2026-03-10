package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type BookmarksMock struct {
	*mocks.GenericMock
}

func NewBookmarksMock() *BookmarksMock {
	return &BookmarksMock{
		GenericMock: mocks.NewJSONMock("BookmarksMock"),
	}
}

func (m *BookmarksMock) RegisterMocks() {
	m.Register("GET", "/api/v1/bookmarks", 200, "validate_list.json")
	m.Register("GET", "/api/v1/bookmarks/1", 200, "validate_get.json")
	m.Register("POST", "/api/v1/bookmarks", 201, "validate_create.json")
	m.Register("PUT", "/api/v1/bookmarks/1", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/bookmarks/1", 204, "")
}

func (m *BookmarksMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v1/bookmarks", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/bookmarks/999", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/bookmarks", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/bookmarks/999", 500, "error_internal.json", "")
	m.RegisterError("DELETE", "/api/v1/bookmarks/999", 500, "error_internal.json", "")
}
