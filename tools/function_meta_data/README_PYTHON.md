# Jamf Pro API Documentation Scraper (Python)

A Python tool that extracts API documentation URLs from Go source files and scrapes the Jamf Developer Portal to gather comprehensive API metadata using Playwright for JavaScript rendering.

## Features

### Web Scraping with Playwright
- ✅ Renders JavaScript-heavy documentation pages
- ✅ Extracts API endpoints, methods, and descriptions
- ✅ Captures parameters and response codes
- ✅ Retrieves request/response schemas
- ✅ Saves raw HTML for debugging

### Output Formats
- ✅ JSON - Machine-readable structured data
- ✅ CSV - Spreadsheet-friendly format
- ✅ Text - Human-readable reports

## Installation

### Requirements
- Python 3.8 or higher
- pip (Python package manager)

### Quick Setup

```bash
cd /Users/dafyddwatkins/GitHub/deploymenttheory/go-sdk-jamfpro-v2/tools/function_meta_data

# Run the setup script
chmod +x setup.sh
./setup.sh
```

This will:
1. Create a Python virtual environment
2. Install required dependencies (Playwright)
3. Download Chromium browser

### Manual Setup

```bash
# Create virtual environment
python3 -m venv venv

# Activate it
source venv/bin/activate

# Install dependencies
pip install -r requirements.txt

# Install Playwright browsers
playwright install chromium
```

## Usage

### Basic Usage

```bash
# Activate virtual environment (if not already active)
source venv/bin/activate

# Run the scraper
python scraper.py -file /path/to/crud.go -format json -debug
```

### Examples

```bash
# Scrape computer_inventory service with debug output
python scraper.py \
  -file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
  -format json \
  -debug \
  -save-html \
  > output.json 2> debug.log

# Generate CSV report
python scraper.py \
  -file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
  -format csv \
  > report.csv

# Human-readable text output
python scraper.py \
  -file ../../jamfpro/services/classic_api/policies/crud.go \
  -format text \
  -debug
```

### Using the Run Script

```bash
chmod +x run_python_scraper.sh
./run_python_scraper.sh
```

## Command-Line Options

```
-file, --file PATH          Path to the Go source file to parse (required)
-format, --format FORMAT    Output format: text, json, csv (default: json)
-debug, --debug             Enable debug output to stderr
-save-html, --save-html     Save raw HTML files to html_dumps/
-headless, --headless       Run browser in headless mode (default: True)
```

## Output Structure

### JSON Format

```json
[
  {
    "url": "https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory",
    "success": true,
    "title": "Return paginated Computer Inventory records",
    "method": "GET",
    "endpoint": "/api/v3/computers-inventory",
    "description": "Returns all computer inventory records with pagination support",
    "requestBody": "",
    "responseBody": "{ ... }",
    "parameters": [
      {
        "name": "page",
        "description": "Page number for pagination"
      }
    ],
    "responseCodes": [
      {
        "code": "200",
        "description": "Successful response"
      }
    ],
    "scrapedAt": "2026-02-23T08:00:00"
  }
]
```

## Troubleshooting

### Virtual Environment Not Found

```bash
# Make sure you're in the right directory
cd /Users/dafyddwatkins/GitHub/deploymenttheory/go-sdk-jamfpro-v2/tools/function_meta_data

# Run setup again
./setup.sh
```

### Playwright Browsers Not Installed

```bash
source venv/bin/activate
playwright install chromium
```

### Permission Denied

```bash
chmod +x setup.sh run_python_scraper.sh
```

## Advantages Over Go Version

- ✅ More mature Playwright implementation
- ✅ Better JavaScript rendering support
- ✅ Easier installation and setup
- ✅ Better error handling
- ✅ More reliable on different platforms

## Files Generated

- `output.json` - Scraped API documentation data
- `debug.log` - Debug output and scraping progress
- `html_dumps/` - Raw HTML files (with `-save-html` flag)
- `venv/` - Python virtual environment

## License

Part of the go-sdk-jamfpro-v2 project.
