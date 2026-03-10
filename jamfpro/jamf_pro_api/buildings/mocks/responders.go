package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type BuildingsMock struct {
	*mocks.GenericMock
}

func NewBuildingsMock() *BuildingsMock {
	return &BuildingsMock{
		GenericMock: mocks.NewJSONMock("BuildingsMock"),
	}
}

func (m *BuildingsMock) RegisterMocks() {
	m.RegisterListBuildingsMock()
	m.RegisterGetBuildingMock()
	m.RegisterCreateBuildingMock()
	m.RegisterUpdateBuildingMock()
	m.RegisterDeleteBuildingMock()
	m.RegisterDeleteBuildingsByIDMock()
	m.RegisterGetBuildingHistoryMock()
	m.RegisterAddBuildingHistoryNotesMock()
}

func (m *BuildingsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
	m.RegisterError("GET", "/api/v1/buildings", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/buildings/delete-multiple", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/buildings/export", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/buildings/1/history/export", 500, "error_internal.json", "")
}

func (m *BuildingsMock) RegisterListBuildingsMock() {
	m.Register("GET", "/api/v1/buildings", 200, "validate_list_buildings.json")
}

func (m *BuildingsMock) RegisterListBuildingsRSQLMock() {
	m.Register("GET", "/api/v1/buildings", 200, "validate_list_buildings_rsql.json")
}

func (m *BuildingsMock) RegisterGetBuildingMock() {
	m.Register("GET", "/api/v1/buildings/1", 200, "validate_get_building.json")
}

func (m *BuildingsMock) RegisterCreateBuildingMock() {
	m.Register("POST", "/api/v1/buildings", 201, "validate_create_building.json")
}

func (m *BuildingsMock) RegisterUpdateBuildingMock() {
	m.Register("PUT", "/api/v1/buildings/1", 200, "validate_update_building.json")
}

func (m *BuildingsMock) RegisterDeleteBuildingMock() {
	m.Register("DELETE", "/api/v1/buildings/1", 204, "")
}

func (m *BuildingsMock) RegisterDeleteBuildingsByIDMock() {
	m.Register("POST", "/api/v1/buildings/delete-multiple", 204, "")
}

func (m *BuildingsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/buildings/999", 404, "error_not_found.json", "")
}

func (m *BuildingsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/buildings", 409, "error_conflict.json", "")
}

func (m *BuildingsMock) RegisterGetBuildingHistoryMock() {
	m.Register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history.json")
}

func (m *BuildingsMock) RegisterGetBuildingHistoryNullDetailsMock() {
	m.Register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history_null_details.json")
}

func (m *BuildingsMock) RegisterAddBuildingHistoryNotesMock() {
	m.Register("POST", "/api/v1/buildings/1/history", 201, "")
}

func (m *BuildingsMock) RegisterExportBuildingsMock() {
	m.Register("POST", "/api/v1/buildings/export", 200, "validate_list_buildings.json")
}

func (m *BuildingsMock) RegisterExportBuildingHistoryMock() {
	m.Register("POST", "/api/v1/buildings/1/history/export", 200, "validate_get_history.json")
}
