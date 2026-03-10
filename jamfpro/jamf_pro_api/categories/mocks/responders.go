package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CategoriesMock struct {
	*mocks.GenericMock
}

func NewCategoriesMock() *CategoriesMock {
	return &CategoriesMock{
		GenericMock: mocks.NewJSONMock("CategoriesMock"),
	}
}

func (m *CategoriesMock) RegisterMocks() {
	m.RegisterListCategoriesMock()
	m.RegisterGetCategoryMock()
	m.RegisterCreateCategoryMock()
	m.RegisterUpdateCategoryMock()
	m.RegisterDeleteCategoryMock()
	m.RegisterDeleteCategoriesBulkMock()
	m.RegisterGetCategoryHistoryMock()
	m.RegisterAddCategoryHistoryNotesMock()
}

func (m *CategoriesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *CategoriesMock) RegisterListCategoriesMock() {
	m.Register("GET", "/api/v1/categories", 200, "validate_list_categories.json")
}

func (m *CategoriesMock) RegisterListCategoriesRSQLMock() {
	m.Register("GET", "/api/v1/categories", 200, "validate_list_categories_rsql.json")
}

func (m *CategoriesMock) RegisterGetCategoryMock() {
	m.Register("GET", "/api/v1/categories/1", 200, "validate_get_category.json")
}

func (m *CategoriesMock) RegisterCreateCategoryMock() {
	m.Register("POST", "/api/v1/categories", 201, "validate_create_category.json")
}

func (m *CategoriesMock) RegisterUpdateCategoryMock() {
	m.Register("PUT", "/api/v1/categories/1", 200, "validate_update_category.json")
}

func (m *CategoriesMock) RegisterDeleteCategoryMock() {
	m.Register("DELETE", "/api/v1/categories/1", 204, "")
}

func (m *CategoriesMock) RegisterDeleteCategoriesBulkMock() {
	m.Register("POST", "/api/v1/categories/delete-multiple", 204, "")
}

func (m *CategoriesMock) RegisterGetCategoryHistoryMock() {
	m.Register("GET", "/api/v1/categories/1/history", 200, "validate_get_history.json")
}

func (m *CategoriesMock) RegisterAddCategoryHistoryNotesMock() {
	m.Register("POST", "/api/v1/categories/1/history", 201, "")
}

func (m *CategoriesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/categories/999", 404, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/categories", 409, "error_conflict.json", "")
}

func (m *CategoriesMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/categories", 500, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterGetCategoryHistoryErrorMock() {
	m.RegisterError("GET", "/api/v1/categories/1/history", 500, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterDeleteCategoryErrorMock() {
	m.RegisterError("DELETE", "/api/v1/categories/1", 500, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterDeleteCategoriesBulkErrorMock() {
	m.RegisterError("POST", "/api/v1/categories/delete-multiple", 500, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterUpdateCategoryErrorMock() {
	m.RegisterError("PUT", "/api/v1/categories/1", 404, "error_not_found.json", "")
}

func (m *CategoriesMock) RegisterAddCategoryHistoryNotesErrorMock() {
	m.RegisterError("POST", "/api/v1/categories/1/history", 500, "error_not_found.json", "")
}
