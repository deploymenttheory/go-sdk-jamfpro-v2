package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDevicePrestagesMock struct {
	*mocks.GenericMock
}

func NewMobileDevicePrestagesMock() *MobileDevicePrestagesMock {
	return &MobileDevicePrestagesMock{
		GenericMock: mocks.NewJSONMock("MobileDevicePrestagesMock"),
	}
}

func (m *MobileDevicePrestagesMock) RegisterListMock() {
	m.Register("GET", "/api/v3/mobile-device-prestages", 200, "validate_list.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterCreateMock() {
	m.Register("POST", "/api/v3/mobile-device-prestages", 200, "validate_create.json")
}

func (m *MobileDevicePrestagesMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v3/mobile-device-prestages/"+id, 200, "")
}

func (m *MobileDevicePrestagesMock) RegisterGetScopeByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterReplaceScopeByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.Register("PUT", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterReplaceScopePutOnlyMock(id string) {
	m.Register("PUT", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterAddScopeByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.Register("POST", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterRemoveScopeByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.Register("POST", "/api/v2/mobile-device-prestages/"+id+"/scope/delete-multiple", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetSyncsByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/syncs", 200, "validate_syncs.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetLatestSyncByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-device-prestages/"+id+"/syncs/latest", 200, "validate_sync_latest.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetAttachmentsByIDMock(id string) {
	m.Register("GET", "/api/v3/mobile-device-prestages/"+id+"/attachments", 200, "validate_attachments.json")
}

func (m *MobileDevicePrestagesMock) RegisterUploadAttachmentMock(id string) {
	m.Register("POST", "/api/v3/mobile-device-prestages/"+id+"/attachments", 200, "validate_attachment_upload.json")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteAttachmentsByIDMock(id string) {
	m.Register("POST", "/api/v3/mobile-device-prestages/"+id+"/attachments/delete-multiple", 200, "")
}

func (m *MobileDevicePrestagesMock) RegisterGetHistoryByIDMock(id string) {
	m.Register("GET", "/api/v3/mobile-device-prestages/"+id+"/history", 200, "validate_history.json")
}

func (m *MobileDevicePrestagesMock) RegisterAddHistoryNoteByIDMock(id string) {
	m.Register("POST", "/api/v3/mobile-device-prestages/"+id+"/history", 200, "validate_add_history_note.json")
}

func (m *MobileDevicePrestagesMock) RegisterEmptyListMock() {
	m.Register("GET", "/api/v3/mobile-device-prestages", 200, "validate_empty_list.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetAllSyncsMock() {
	m.Register("GET", "/api/v2/mobile-device-prestages/syncs", 200, "validate_syncs.json")
}

func (m *MobileDevicePrestagesMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v3/mobile-device-prestages", 500, "error_internal.json", "no response")
}

func (m *MobileDevicePrestagesMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v3/mobile-device-prestages", 500, "error_internal.json", "no response")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteByIDNoResponseErrorMock(id string) {
	m.RegisterError("DELETE", "/api/v3/mobile-device-prestages/"+id, 500, "error_internal.json", "no response")
}
