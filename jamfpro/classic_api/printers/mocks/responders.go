package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PrintersMock struct {
	*mocks.GenericMock
}

func NewPrintersMock() *PrintersMock {
	return &PrintersMock{
		GenericMock: mocks.NewXMLMock("PrintersMock"),
	}
}

func (m *PrintersMock) RegisterMocks() {
	m.RegisterListPrintersMock()
	m.RegisterGetPrinterByIDMock()
	m.RegisterGetPrinterByNameMock()
	m.RegisterCreatePrinterMock()
	m.RegisterUpdatePrinterByIDMock()
	m.RegisterUpdatePrinterByNameMock()
	m.RegisterDeletePrinterByIDMock()
	m.RegisterDeletePrinterByNameMock()
}

func (m *PrintersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PrintersMock) RegisterListPrintersMock() {
	m.Register("GET", "/JSSResource/printers", 200, "validate_list_printers.xml")
}

func (m *PrintersMock) RegisterGetPrinterByIDMock() {
	m.Register("GET", "/JSSResource/printers/id/1", 200, "validate_get_printer.xml")
}

func (m *PrintersMock) RegisterGetPrinterByNameMock() {
	m.Register("GET", "/JSSResource/printers/name/Office Printer", 200, "validate_get_printer.xml")
}

func (m *PrintersMock) RegisterCreatePrinterMock() {
	m.Register("POST", "/JSSResource/printers/id/0", 201, "validate_create_printer.xml")
}

func (m *PrintersMock) RegisterUpdatePrinterByIDMock() {
	m.Register("PUT", "/JSSResource/printers/id/1", 200, "validate_update_printer.xml")
}

func (m *PrintersMock) RegisterUpdatePrinterByNameMock() {
	m.Register("PUT", "/JSSResource/printers/name/Office Printer", 200, "validate_update_printer.xml")
}

func (m *PrintersMock) RegisterDeletePrinterByIDMock() {
	m.Register("DELETE", "/JSSResource/printers/id/1", 200, "")
}

func (m *PrintersMock) RegisterDeletePrinterByNameMock() {
	m.Register("DELETE", "/JSSResource/printers/name/Office Printer", 200, "")
}

func (m *PrintersMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/printers/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *PrintersMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/printers/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A printer with that name already exists")
}
