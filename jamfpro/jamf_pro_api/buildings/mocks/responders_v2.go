package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// BuildingsMockV2 demonstrates the new generic mock pattern with comprehensive registration methods.
// This replaces the old 251-line responders.go with just ~60 lines.
type BuildingsMockV2 struct {
	*mocks.GenericMock
}

// NewBuildingsMockV2 creates a new mock for buildings testing.
func NewBuildingsMockV2() *BuildingsMockV2 {
	return &BuildingsMockV2{
		GenericMock: mocks.NewJSONMock("BuildingsMock"),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *BuildingsMockV2) RegisterMocks() {
	m.RegisterListBuildingsMock()
	m.RegisterGetBuildingMock()
	m.RegisterCreateBuildingMock()
	m.RegisterUpdateBuildingMock()
	m.RegisterDeleteBuildingMock()
	m.RegisterDeleteBuildingsByIDMock()
	m.RegisterGetBuildingHistoryMock()
	m.RegisterAddBuildingHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *BuildingsMockV2) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *BuildingsMockV2) RegisterListBuildingsMock() {
	m.Register("GET", "/api/v1/buildings", 200, "validate_list_buildings.json")
}

func (m *BuildingsMockV2) RegisterListBuildingsRSQLMock() {
	m.Register("GET", "/api/v1/buildings", 200, "validate_list_buildings_rsql.json")
}

func (m *BuildingsMockV2) RegisterGetBuildingMock() {
	m.Register("GET", "/api/v1/buildings/1", 200, "validate_get_building.json")
}

func (m *BuildingsMockV2) RegisterCreateBuildingMock() {
	m.Register("POST", "/api/v1/buildings", 201, "validate_create_building.json")
}

func (m *BuildingsMockV2) RegisterUpdateBuildingMock() {
	m.Register("PUT", "/api/v1/buildings/1", 200, "validate_update_building.json")
}

func (m *BuildingsMockV2) RegisterDeleteBuildingMock() {
	m.Register("DELETE", "/api/v1/buildings/1", 204, "")
}

func (m *BuildingsMockV2) RegisterDeleteBuildingsByIDMock() {
	m.Register("POST", "/api/v1/buildings/delete-multiple", 204, "")
}

func (m *BuildingsMockV2) RegisterGetBuildingHistoryMock() {
	m.Register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history.json")
}

func (m *BuildingsMockV2) RegisterGetBuildingHistoryNullDetailsMock() {
	m.Register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history_null_details.json")
}

func (m *BuildingsMockV2) RegisterAddBuildingHistoryNotesMock() {
	m.Register("POST", "/api/v1/buildings/1/history", 201, "validate_get_history.json")
}

func (m *BuildingsMockV2) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/buildings/999", 404, "error_not_found.json", "")
}

func (m *BuildingsMockV2) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/buildings", 409, "error_conflict.json", "")
}

func (m *BuildingsMockV2) RegisterExportBuildingsMock() {
	m.Register("POST", "/api/v1/buildings/export", 200, "")
}

func (m *BuildingsMockV2) RegisterExportBuildingHistoryMock() {
	m.Register("POST", "/api/v1/buildings/1/history/export", 200, "")
}
