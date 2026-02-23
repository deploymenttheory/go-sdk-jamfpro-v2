# Function Metadata Parser & API Documentation Scraper

A Go tool that parses service CRUD files to extract API documentation URLs, then automatically scrapes and parses those documentation pages to gather comprehensive API metadata.

## Features

### Code Parsing
- ✅ Extracts function names and descriptions from Go source
- ✅ Identifies HTTP methods (GET, POST, PUT, PATCH, DELETE)
- ✅ Extracts API endpoints and paths
- ✅ Captures API documentation URLs from comments
- ✅ Identifies request and response types
- ✅ Detects RSQL query support
- ✅ Identifies paginated endpoints
- ✅ Extracts HTTP headers (Accept, Content-Type)

### Web Scraping
- ✅ Automatically fetches API documentation from Jamf Developer Portal
- ✅ Extracts endpoint details from documentation
- ✅ Parses request/response schemas
- ✅ Captures API parameters and descriptions
- ✅ Extracts HTTP response codes
- ✅ Retrieves detailed API descriptions

### Output
- ✅ Multiple output formats (text, JSON, CSV)
- ✅ Debug mode for detailed parsing and scraping insights
- ✅ Comprehensive reports combining code analysis and documentation

## Installation

```bash
cd tools/function_meta_data
go build -o function_meta_data
```

Or use the Makefile:

```bash
make build
```

## Usage

### Basic Usage

The tool parses the specified Go file to extract API documentation URLs, then automatically scrapes each URL to gather detailed API information:

```bash
./function_meta_data -file /path/to/crud.go
```

**Note:** The tool will make HTTP requests to fetch documentation pages. Ensure you have internet connectivity.

### With Debug Output

Debug mode shows detailed information about both the parsing and scraping process:

```bash
./function_meta_data -file /path/to/crud.go -debug
```

This will display:
- Go code parsing progress
- Each URL being scraped
- Extracted data from documentation pages
- Any errors encountered during scraping

### Different Output Formats

```bash
# Text format (default, human-readable)
./function_meta_data -file /path/to/crud.go -format text

# JSON format (machine-readable)
./function_meta_data -file /path/to/crud.go -format json

# CSV format (for spreadsheets)
./function_meta_data -file /path/to/crud.go -format csv
```

## Examples

### Analyze Computer Inventory Service

```bash
./function_meta_data \
  -file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
  -debug
```

### Analyze Classic API Policies Service

```bash
./function_meta_data \
  -file ../../jamfpro/services/classic_api/policies/crud.go \
  -format json > policies_metadata.json
```

### Generate CSV Report

```bash
./function_meta_data \
  -file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
  -format csv > computer_inventory_report.csv
```

## Output Format

### Text Format

```
================================================================================
FUNCTION METADATA EXTRACTION REPORT
================================================================================

Total Functions Analyzed: 15

[1] CreateV3
--------------------------------------------------------------------------------
Description:    Creates a new computer inventory record
HTTP Method:    POST
Endpoint:       /api/v3/computers-inventory
API Docs:       https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory
Request Type:   *ResourceComputerInventory
Response Type:  *ResourceComputerInventory
Accept:         mime.ApplicationJSON
Content-Type:   mime.ApplicationJSON
RSQL Support:   false
Paginated:      false
Parameters:
  - ctx: context.Context
  - request: *ResourceComputerInventory
```

### JSON Format

```json
[
  {
    "functionName": "CreateV3",
    "httpMethod": "POST",
    "endpoint": "/api/v3/computers-inventory",
    "apiDocsURL": "https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory",
    "description": "Creates a new computer inventory record",
    "requestType": "*ResourceComputerInventory",
    "responseType": "*ResourceComputerInventory",
    "supportsRSQL": false,
    "isPaginated": false
  }
]
```

### CSV Format

```csv
Function,HTTP Method,Endpoint,Request Type,Response Type,RSQL,Paginated,API Docs
CreateV3,POST,/api/v3/computers-inventory,*ResourceComputerInventory,*ResourceComputerInventory,false,false,https://developer.jamf.com/...
```

## Metadata Extracted

For each function, the tool extracts:

| Field | Description |
|-------|-------------|
| `functionName` | The name of the exported function |
| `httpMethod` | HTTP method (GET, POST, PUT, PATCH, DELETE) |
| `endpoint` | API endpoint path |
| `apiDocsURL` | Link to official Jamf Pro API documentation |
| `description` | Function description from comments |
| `requestType` | Type of the request parameter |
| `responseType` | Type of the response |
| `acceptHeader` | Value of Accept header |
| `contentType` | Value of Content-Type header |
| `supportsRSQL` | Whether function supports RSQL queries |
| `isPaginated` | Whether function uses pagination |
| `parameters` | List of all function parameters |

## Debug Mode

When `-debug` flag is enabled, the tool provides detailed insights into the parsing process:

```
🔍 DEBUG: Parsing file: /path/to/crud.go
🔍 DEBUG: Package name: computer_inventory
🔍 DEBUG: Total declarations: 25
🔍 DEBUG: Found 15 interface methods

🔍 DEBUG: Processing function: CreateV3
  📄 Description: Creates a new computer inventory record
  🌐 URL: POST /api/v3/computers-inventory
  📚 API Docs: https://developer.jamf.com/...
  🔧 HTTP Method: POST
  🎯 Endpoint: /api/v3/computers-inventory
  📥 Parameter: ctx context.Context
  📥 Parameter: request *ResourceComputerInventory
  📤 Return type 0: *ResourceComputerInventory
  📋 Accept: mime.ApplicationJSON
  📋 Content-Type: mime.ApplicationJSON
  🔍 RSQL Support: false
  📄 Paginated: false
  📨 Request Type: *ResourceComputerInventory
```

## Use Cases

1. **Documentation Generation**: Extract metadata to auto-generate API documentation
2. **API Coverage Analysis**: Identify which endpoints are implemented
3. **Code Review**: Understand API surface quickly
4. **Testing**: Generate test matrices based on extracted metadata
5. **Migration Planning**: Compare old vs new implementations
6. **API Client Generation**: Use metadata to generate client SDKs

## Requirements

- Go 1.19 or higher
- Access to Go source files to parse

## License

Part of the go-sdk-jamfpro-v2 project.
