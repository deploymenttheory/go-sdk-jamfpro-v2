package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// FunctionMetadata holds extracted metadata for a single function
type FunctionMetadata struct {
	FunctionName     string
	HTTPMethod       string
	Endpoint         string
	URL              string
	APIDocsURL       string
	Description      string
	RequestType      string
	ResponseType     string
	ErrorReturn      bool
	InterfaceComment string
	Parameters       []Parameter
	AcceptHeader     string
	ContentType      string
	SupportsRSQL     bool
	IsPaginated      bool
	ScrapedDocs      *ScrapedDocumentation // Added for web scraping
}

// ScrapedDocumentation holds information scraped from API documentation
type ScrapedDocumentation struct {
	URL              string
	Title            string
	Method           string
	Endpoint         string
	Description      string
	RequestBody      string
	ResponseBody     string
	Parameters       []APIParameter
	ResponseCodes    []ResponseCode
	RawHTML          string // Store raw HTML for debugging
	ScrapedAt        time.Time
	Success          bool
	Error            string
}

// APIParameter represents a parameter from the documentation
type APIParameter struct {
	Name        string
	Type        string
	Required    bool
	Description string
	Location    string // query, path, body, header
}

// ResponseCode represents an HTTP response code
type ResponseCode struct {
	Code        string
	Description string
}

// Parameter represents a function parameter
type Parameter struct {
	Name string
	Type string
}

var (
	debug    = false
	saveHTML = false
)

func main() {
	// Command-line flags
	filePath := flag.String("file", "", "Path to the Go file to parse")
	debugFlag := flag.Bool("debug", false, "Enable debug output")
	outputFormat := flag.String("format", "text", "Output format: text, json, csv")
	saveHTMLFlag := flag.Bool("save-html", false, "Save raw HTML files for inspection")

	flag.Parse()

	if *filePath == "" {
		fmt.Fprintln(os.Stderr, "Usage: function_meta_data -file <path-to-go-file> [-debug] [-format text|json|csv]")
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintln(os.Stderr, "  function_meta_data -file /path/to/crud.go -debug")
		os.Exit(1)
	}

	debug = *debugFlag
	saveHTML = *saveHTMLFlag

	if debug {
		fmt.Fprintf(os.Stderr, "🔍 DEBUG: Extracting API documentation URLs from: %s\n", *filePath)
	}
	if saveHTML {
		fmt.Fprintf(os.Stderr, "💾 Will save raw HTML files to ./html_dumps/\n")
		os.MkdirAll("html_dumps", 0755)
	}

	// Extract URLs from the file
	urls, err := extractDocURLs(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error extracting URLs: %v\n", err)
		os.Exit(1)
	}

	if debug {
		fmt.Fprintf(os.Stderr, "🔍 DEBUG: Found %d API documentation URLs\n", len(urls))
	}

	// Scrape each URL
	var scrapedDocs []*ScrapedDocumentation
	for _, url := range urls {
		doc := scrapeDocumentation(url)
		scrapedDocs = append(scrapedDocs, doc)
	}

	// Output results
	switch *outputFormat {
	case "json":
		printScrapedJSON(scrapedDocs)
	case "csv":
		printScrapedCSV(scrapedDocs)
	default:
		printScrapedText(scrapedDocs)
	}
}

// extractDocURLs extracts all Jamf Pro API documentation URLs from a Go source file
func extractDocURLs(filePath string) ([]string, error) {
	fset := token.NewFileSet()

	// Parse the file
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	var urls []string
	seen := make(map[string]bool)

	// Extract URLs from all comments
	for _, commentGroup := range file.Comments {
		text := commentGroup.Text()
		url := extractAPIDocsURL(text)
		if url != "" && !seen[url] {
			urls = append(urls, url)
			seen[url] = true
			if debug {
				fmt.Fprintf(os.Stderr, "  📎 Found URL: %s\n", url)
			}
		}
	}

	return urls, nil
}

// parseFile parses a Go source file and extracts function metadata
func parseFile(filePath string) ([]FunctionMetadata, error) {
	fset := token.NewFileSet()

	if debug {
		fmt.Printf("🔍 DEBUG: Reading file: %s\n", filePath)
	}

	// Parse the file
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	if debug {
		fmt.Printf("🔍 DEBUG: Package name: %s\n", file.Name.Name)
		fmt.Printf("🔍 DEBUG: Total declarations: %d\n", len(file.Decls))
	}

	var metadata []FunctionMetadata

	// First pass: Extract interface comments for later matching
	interfaceComments := extractInterfaceComments(file)

	if debug {
		fmt.Printf("🔍 DEBUG: Found %d interface methods\n", len(interfaceComments))
	}

	// Second pass: Extract function implementations
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Skip non-exported functions
		if !funcDecl.Name.IsExported() {
			continue
		}

		// Skip constructor functions
		if strings.HasPrefix(funcDecl.Name.Name, "New") {
			continue
		}

		if debug {
			fmt.Printf("\n🔍 DEBUG: Processing function: %s\n", funcDecl.Name.Name)
		}

		meta := extractFunctionMetadata(funcDecl, fset, interfaceComments)
		if meta.FunctionName != "" {
			// Scrape documentation if URL is available
			if meta.APIDocsURL != "" {
				meta.ScrapedDocs = scrapeDocumentation(meta.APIDocsURL)
			}
			metadata = append(metadata, meta)
		}
	}

	return metadata, nil
}

// extractInterfaceComments extracts comments from interface method declarations
func extractInterfaceComments(file *ast.File) map[string]string {
	comments := make(map[string]string)

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			if debug {
				fmt.Printf("🔍 DEBUG: Found interface: %s\n", typeSpec.Name.Name)
			}

			// Extract method comments from interface
			for _, method := range interfaceType.Methods.List {
				if len(method.Names) == 0 {
					continue
				}

				methodName := method.Names[0].Name
				if method.Doc != nil {
					comment := method.Doc.Text()
					comments[methodName] = strings.TrimSpace(comment)

					if debug {
						fmt.Printf("  📝 Method: %s\n", methodName)
						fmt.Printf("     Comment: %s\n", truncate(comment, 60))
					}
				}
			}
		}
	}

	return comments
}

// extractFunctionMetadata extracts metadata from a function declaration
func extractFunctionMetadata(funcDecl *ast.FuncDecl, fset *token.FileSet, interfaceComments map[string]string) FunctionMetadata {
	meta := FunctionMetadata{
		FunctionName: funcDecl.Name.Name,
		Parameters:   make([]Parameter, 0),
	}

	// Extract function comment
	if funcDecl.Doc != nil {
		comment := funcDecl.Doc.Text()
		meta.Description = extractDescription(comment)
		meta.URL = extractURL(comment)
		meta.APIDocsURL = extractAPIDocsURL(comment)

		if debug {
			fmt.Printf("  📄 Description: %s\n", truncate(meta.Description, 60))
			fmt.Printf("  🌐 URL: %s\n", meta.URL)
			fmt.Printf("  📚 API Docs: %s\n", meta.APIDocsURL)
		}
	}

	// Get interface comment if available
	if interfaceComment, ok := interfaceComments[meta.FunctionName]; ok {
		meta.InterfaceComment = interfaceComment

		// Extract additional info from interface comment if not in function comment
		if meta.APIDocsURL == "" {
			meta.APIDocsURL = extractAPIDocsURL(interfaceComment)
		}
		if meta.Description == "" {
			meta.Description = extractDescription(interfaceComment)
		}
	}

	// Extract HTTP method from URL comment
	if meta.URL != "" {
		meta.HTTPMethod = extractHTTPMethod(meta.URL)
		meta.Endpoint = extractEndpoint(meta.URL)

		if debug {
			fmt.Printf("  🔧 HTTP Method: %s\n", meta.HTTPMethod)
			fmt.Printf("  🎯 Endpoint: %s\n", meta.Endpoint)
		}
	}

	// Extract parameters
	if funcDecl.Type.Params != nil {
		for _, param := range funcDecl.Type.Params.List {
			paramType := exprToString(param.Type)
			for _, name := range param.Names {
				meta.Parameters = append(meta.Parameters, Parameter{
					Name: name.Name,
					Type: paramType,
				})

				if debug {
					fmt.Printf("  📥 Parameter: %s %s\n", name.Name, paramType)
				}
			}
		}
	}

	// Extract return types
	if funcDecl.Type.Results != nil {
		results := funcDecl.Type.Results.List
		if len(results) > 0 {
			// First non-error return type
			for i, result := range results {
				typeStr := exprToString(result.Type)

				if typeStr == "error" {
					meta.ErrorReturn = true
				} else if typeStr == "*interfaces.Response" {
					// Skip response type
					continue
				} else if i == 0 {
					meta.ResponseType = typeStr
				}

				if debug {
					fmt.Printf("  📤 Return type %d: %s\n", i, typeStr)
				}
			}
		}
	}

	// Analyze function body for additional metadata
	if funcDecl.Body != nil {
		meta.AcceptHeader = extractHeaderValue(funcDecl.Body, "Accept")
		meta.ContentType = extractHeaderValue(funcDecl.Body, "Content-Type")
		meta.SupportsRSQL = checkRSQLSupport(funcDecl.Body)
		meta.IsPaginated = checkPagination(funcDecl.Body)

		// Extract request type from function body
		if meta.Parameters != nil {
			for _, param := range meta.Parameters {
				if strings.HasPrefix(param.Type, "*") && !strings.Contains(param.Type, "context") {
					meta.RequestType = param.Type
					break
				}
			}
		}

		if debug {
			fmt.Printf("  📋 Accept: %s\n", meta.AcceptHeader)
			fmt.Printf("  📋 Content-Type: %s\n", meta.ContentType)
			fmt.Printf("  🔍 RSQL Support: %v\n", meta.SupportsRSQL)
			fmt.Printf("  📄 Paginated: %v\n", meta.IsPaginated)
			fmt.Printf("  📨 Request Type: %s\n", meta.RequestType)
		}
	}

	return meta
}

// extractDescription extracts the description from a comment
func extractDescription(comment string) string {
	lines := strings.Split(comment, "\n")
	if len(lines) == 0 {
		return ""
	}

	// First non-empty line is usually the description
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "URL:") && !strings.HasPrefix(line, "Jamf Pro API") {
			return line
		}
	}
	return ""
}

// extractURL extracts the URL line from a comment (e.g., "URL: GET /api/v3/...")
func extractURL(comment string) string {
	re := regexp.MustCompile(`URL:\s*(.+)`)
	matches := re.FindStringSubmatch(comment)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

// extractAPIDocsURL extracts the Jamf Pro API docs URL from a comment
func extractAPIDocsURL(comment string) string {
	re := regexp.MustCompile(`https://developer\.jamf\.com/[^\s)]+`)
	matches := re.FindStringSubmatch(comment)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// extractHTTPMethod extracts HTTP method from URL string (e.g., "GET /api/v3/..." -> "GET")
func extractHTTPMethod(url string) string {
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	urlUpper := strings.ToUpper(url)
	for _, method := range methods {
		if strings.HasPrefix(urlUpper, method) {
			return method
		}
	}
	return ""
}

// extractEndpoint extracts the endpoint path from URL string (e.g., "GET /api/v3/..." -> "/api/v3/...")
func extractEndpoint(url string) string {
	parts := strings.Fields(url)
	if len(parts) >= 2 {
		return parts[1]
	}
	return ""
}

// extractHeaderValue extracts a header value from function body
func extractHeaderValue(body *ast.BlockStmt, headerName string) string {
	var headerValue string

	ast.Inspect(body, func(n ast.Node) bool {
		if compLit, ok := n.(*ast.CompositeLit); ok {
			for _, elt := range compLit.Elts {
				if kvExpr, ok := elt.(*ast.KeyValueExpr); ok {
					keyStr := exprToString(kvExpr.Key)
					valStr := exprToString(kvExpr.Value)

					if keyStr == fmt.Sprintf(`"%s"`, headerName) {
						if strings.Contains(valStr, "mime.") {
							headerValue = valStr
							return false
						}
					}
				}
			}
		}
		return true
	})

	return headerValue
}

// checkRSQLSupport checks if the function supports RSQL queries
func checkRSQLSupport(body *ast.BlockStmt) bool {
	hasRSQL := false

	ast.Inspect(body, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			if ident.Name == "rsqlQuery" || strings.Contains(ident.Name, "RSQL") {
				hasRSQL = true
				return false
			}
		}
		return true
	})

	return hasRSQL
}

// checkPagination checks if the function uses pagination
func checkPagination(body *ast.BlockStmt) bool {
	hasPagination := false

	ast.Inspect(body, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			if ident.Name == "GetPaginated" || ident.Name == "mergePage" {
				hasPagination = true
				return false
			}
		}
		return true
	})

	return hasPagination
}

// exprToString converts an AST expression to a string representation
func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return "*" + exprToString(e.X)
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(e.Elt)
	case *ast.MapType:
		return "map[" + exprToString(e.Key) + "]" + exprToString(e.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return fmt.Sprintf("%T", expr)
	}
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// scrapeDocumentation fetches and parses API documentation from a URL
func scrapeDocumentation(url string) *ScrapedDocumentation {
	doc := &ScrapedDocumentation{
		URL:       url,
		ScrapedAt: time.Now(),
		Success:   false,
	}

	if debug {
		fmt.Fprintf(os.Stderr, "\n🌐 Scraping documentation from: %s\n", url)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Fetch the documentation page
	resp, err := client.Get(url)
	if err != nil {
		doc.Error = fmt.Sprintf("Failed to fetch URL: %v", err)
		if debug {
			fmt.Fprintf(os.Stderr, "  ❌ Error: %s\n", doc.Error)
		}
		return doc
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		doc.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status)
		if debug {
			fmt.Fprintf(os.Stderr, "  ❌ Error: %s\n", doc.Error)
		}
		return doc
	}

	// Read the HTML content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		doc.Error = fmt.Sprintf("Failed to read response: %v", err)
		if debug {
			fmt.Fprintf(os.Stderr, "  ❌ Error: %s\n", doc.Error)
		}
		return doc
	}

	doc.RawHTML = string(body)
	doc.Success = true

	if debug {
		fmt.Fprintf(os.Stderr, "  📄 HTML Length: %d bytes\n", len(body))
		fmt.Fprintf(os.Stderr, "  📄 HTML Preview (first 500 chars):\n%s\n", truncate(string(body), 500))
	}

	// Save HTML to file if requested
	if saveHTML {
		// Create filename from URL
		filename := regexp.MustCompile(`[^a-zA-Z0-9-]`).ReplaceAllString(url, "_")
		filepath := fmt.Sprintf("html_dumps/%s.html", filename)
		if err := os.WriteFile(filepath, body, 0644); err != nil {
			if debug {
				fmt.Fprintf(os.Stderr, "  ⚠️  Failed to save HTML: %v\n", err)
			}
		} else if debug {
			fmt.Fprintf(os.Stderr, "  💾 Saved HTML to: %s\n", filepath)
		}
	}

	// Parse the HTML to extract API information
	parseJamfDocumentation(doc, string(body))

	if debug {
		fmt.Fprintf(os.Stderr, "  ✅ Successfully scraped documentation\n")
		if doc.Title != "" {
			fmt.Fprintf(os.Stderr, "  📄 Title: %s\n", doc.Title)
		}
		if doc.Method != "" {
			fmt.Fprintf(os.Stderr, "  🔧 Method: %s\n", doc.Method)
		}
		if doc.Endpoint != "" {
			fmt.Fprintf(os.Stderr, "  🎯 Endpoint: %s\n", doc.Endpoint)
		}
		if doc.Description != "" {
			fmt.Fprintf(os.Stderr, "  📝 Description: %s\n", truncate(doc.Description, 100))
		}
		fmt.Fprintf(os.Stderr, "  📊 Parameters: %d\n", len(doc.Parameters))
		fmt.Fprintf(os.Stderr, "  📨 Response codes: %d\n", len(doc.ResponseCodes))
	}

	return doc
}

// parseJamfDocumentation parses the HTML content to extract API documentation
func parseJamfDocumentation(doc *ScrapedDocumentation, html string) {
	if debug {
		fmt.Fprintf(os.Stderr, "\n  🔍 PARSING HTML CONTENT:\n")
	}

	// Extract title - look for common patterns in Jamf documentation
	titleRe := regexp.MustCompile(`<title>([^<]+)</title>`)
	if matches := titleRe.FindStringSubmatch(html); len(matches) > 1 {
		doc.Title = strings.TrimSpace(matches[1])
		if debug {
			fmt.Fprintf(os.Stderr, "  ✓ Title found: %s\n", doc.Title)
		}
	} else if debug {
		fmt.Fprintf(os.Stderr, "  ✗ Title not found\n")
	}

	// Extract HTTP method and endpoint
	// Pattern: "POST /api/v1/..." or similar
	endpointRe := regexp.MustCompile(`(?i)(GET|POST|PUT|PATCH|DELETE)\s+(/api/[^\s<>"]+)`)
	if matches := endpointRe.FindStringSubmatch(html); len(matches) > 2 {
		doc.Method = strings.ToUpper(matches[1])
		doc.Endpoint = matches[2]
		if debug {
			fmt.Fprintf(os.Stderr, "  ✓ Endpoint found: %s %s\n", doc.Method, doc.Endpoint)
		}
	} else if debug {
		fmt.Fprintf(os.Stderr, "  ✗ Endpoint not found with pattern: (GET|POST|PUT|PATCH|DELETE)\\s+(/api/...)\n")
		// Try to find any /api/ mentions for debugging
		apiRe := regexp.MustCompile(`/api/[^\s<>"]+`)
		if apiMatches := apiRe.FindAllString(html, 5); len(apiMatches) > 0 {
			fmt.Fprintf(os.Stderr, "    Found API paths in HTML: %v\n", apiMatches)
		}
	}

	// Extract description - look for meta description or first paragraph
	descRe := regexp.MustCompile(`<meta\s+name="description"\s+content="([^"]+)"`)
	if matches := descRe.FindStringSubmatch(html); len(matches) > 1 {
		doc.Description = strings.TrimSpace(matches[1])
		if debug {
			fmt.Fprintf(os.Stderr, "  ✓ Meta description found: %s\n", truncate(doc.Description, 80))
		}
	} else if debug {
		fmt.Fprintf(os.Stderr, "  ✗ Meta description not found\n")
	}

	// If no meta description, try to find description in common div classes
	if doc.Description == "" {
		descDivRe := regexp.MustCompile(`<div[^>]*class="[^"]*description[^"]*"[^>]*>([^<]+)</div>`)
		if matches := descDivRe.FindStringSubmatch(html); len(matches) > 1 {
			doc.Description = strings.TrimSpace(stripHTML(matches[1]))
			if debug {
				fmt.Fprintf(os.Stderr, "  ✓ Div description found: %s\n", truncate(doc.Description, 80))
			}
		}
	}

	// Extract request body schema
	requestRe := regexp.MustCompile(`(?s)"Request Body"[^{]*(\{[^}]+\}|\[.*?\])`)
	if matches := requestRe.FindStringSubmatch(html); len(matches) > 1 {
		doc.RequestBody = strings.TrimSpace(matches[1])
		if debug {
			fmt.Fprintf(os.Stderr, "  ✓ Request body found: %s\n", truncate(doc.RequestBody, 80))
		}
	} else if debug {
		fmt.Fprintf(os.Stderr, "  ✗ Request body not found\n")
		// Look for alternatives
		if strings.Contains(html, "Request Body") {
			fmt.Fprintf(os.Stderr, "    (but 'Request Body' text exists in HTML)\n")
		}
		if strings.Contains(html, "request") || strings.Contains(html, "Request") {
			fmt.Fprintf(os.Stderr, "    Found 'request' mentions in HTML\n")
		}
	}

	// Extract response body schema
	responseRe := regexp.MustCompile(`(?s)"Response"[^{]*(\{[^}]+\}|\[.*?\])`)
	if matches := responseRe.FindStringSubmatch(html); len(matches) > 1 {
		doc.ResponseBody = strings.TrimSpace(matches[1])
		if debug {
			fmt.Fprintf(os.Stderr, "  ✓ Response body found: %s\n", truncate(doc.ResponseBody, 80))
		}
	} else if debug {
		fmt.Fprintf(os.Stderr, "  ✗ Response body not found\n")
		// Look for JSON-like structures
		jsonRe := regexp.MustCompile(`\{[^}]{20,}\}`)
		if jsonMatches := jsonRe.FindAllString(html, 3); len(jsonMatches) > 0 {
			fmt.Fprintf(os.Stderr, "    Found JSON-like structures: %d instances\n", len(jsonMatches))
			for i, match := range jsonMatches {
				if i < 2 {
					fmt.Fprintf(os.Stderr, "      Example %d: %s\n", i+1, truncate(match, 100))
				}
			}
		}
	}

	// Extract parameters - this is complex and depends on Jamf's HTML structure
	// Look for table rows or list items containing parameter information
	if debug {
		fmt.Fprintf(os.Stderr, "  🔍 Searching for parameters...\n")
	}
	extractParameters(doc, html)

	// Extract response codes
	if debug {
		fmt.Fprintf(os.Stderr, "  🔍 Searching for response codes...\n")
	}
	extractResponseCodes(doc, html)
}

// extractParameters extracts parameter information from HTML
func extractParameters(doc *ScrapedDocumentation, html string) {
	// Pattern to find parameter tables or lists
	// This is a simplified version - may need adjustment based on actual HTML structure
	paramRe := regexp.MustCompile(`(?s)<tr[^>]*>.*?<td[^>]*>([^<]+)</td>.*?<td[^>]*>([^<]+)</td>.*?</tr>`)
	matches := paramRe.FindAllStringSubmatch(html, -1)

	if debug {
		fmt.Fprintf(os.Stderr, "    Found %d table rows with 2+ cells\n", len(matches))
	}

	for i, match := range matches {
		if len(match) > 2 {
			name := strings.TrimSpace(stripHTML(match[1]))
			description := strings.TrimSpace(stripHTML(match[2]))

			// Skip if it looks like a header row
			if strings.ToLower(name) == "name" || strings.ToLower(name) == "parameter" {
				if debug && i < 5 {
					fmt.Fprintf(os.Stderr, "    Skipping header row: %s\n", name)
				}
				continue
			}

			param := APIParameter{
				Name:        name,
				Description: description,
			}
			doc.Parameters = append(doc.Parameters, param)

			if debug && i < 5 {
				fmt.Fprintf(os.Stderr, "    ✓ Parameter: %s = %s\n", name, truncate(description, 50))
			}
		}
	}

	if debug && len(doc.Parameters) == 0 {
		// Try to show what table structure exists
		tableRe := regexp.MustCompile(`<table[^>]*>`)
		if tableMatches := tableRe.FindAllString(html, -1); len(tableMatches) > 0 {
			fmt.Fprintf(os.Stderr, "    Found %d <table> tags but no parameters extracted\n", len(tableMatches))
		} else {
			fmt.Fprintf(os.Stderr, "    No <table> tags found in HTML\n")
		}
	}
}

// extractResponseCodes extracts HTTP response codes from HTML
func extractResponseCodes(doc *ScrapedDocumentation, html string) {
	// Look for common patterns like "200 OK" or "201 Created"
	codeRe := regexp.MustCompile(`(?m)(\d{3})\s+([A-Za-z\s]+)`)
	matches := codeRe.FindAllStringSubmatch(html, -1)

	if debug {
		fmt.Fprintf(os.Stderr, "    Found %d potential HTTP status codes (pattern: \\d{3}\\s+[A-Za-z\\s]+)\n", len(matches))
	}

	seen := make(map[string]bool)
	validCount := 0
	for i, match := range matches {
		if len(match) > 2 {
			code := match[1]
			desc := strings.TrimSpace(match[2])

			// Only include valid HTTP codes
			if code >= "200" && code <= "599" && !seen[code] {
				doc.ResponseCodes = append(doc.ResponseCodes, ResponseCode{
					Code:        code,
					Description: desc,
				})
				seen[code] = true
				validCount++

				if debug && validCount <= 5 {
					fmt.Fprintf(os.Stderr, "    ✓ Response code: %s %s\n", code, truncate(desc, 50))
				}
			} else if debug && i < 10 && !seen[code] {
				fmt.Fprintf(os.Stderr, "    Skipping: %s %s (invalid or duplicate)\n", code, truncate(desc, 30))
			}
		}
	}

	if debug && len(doc.ResponseCodes) > 5 {
		fmt.Fprintf(os.Stderr, "    ... and %d more response codes\n", len(doc.ResponseCodes)-5)
	}
}

// stripHTML removes HTML tags from a string
func stripHTML(s string) string {
	tagRe := regexp.MustCompile(`<[^>]*>`)
	s = tagRe.ReplaceAllString(s, "")
	// Replace common HTML entities
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	return strings.TrimSpace(s)
}

// printText outputs metadata in human-readable text format
func printText(metadata []FunctionMetadata) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("FUNCTION METADATA EXTRACTION REPORT")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("\nTotal Functions Analyzed: %d\n\n", len(metadata))

	for i, meta := range metadata {
		fmt.Printf("\n[%d] %s\n", i+1, meta.FunctionName)
		fmt.Println(strings.Repeat("-", 80))

		if meta.Description != "" {
			fmt.Printf("Description:    %s\n", meta.Description)
		}
		if meta.HTTPMethod != "" {
			fmt.Printf("HTTP Method:    %s\n", meta.HTTPMethod)
		}
		if meta.Endpoint != "" {
			fmt.Printf("Endpoint:       %s\n", meta.Endpoint)
		}
		if meta.APIDocsURL != "" {
			fmt.Printf("API Docs:       %s\n", meta.APIDocsURL)
		}
		if meta.RequestType != "" {
			fmt.Printf("Request Type:   %s\n", meta.RequestType)
		}
		if meta.ResponseType != "" {
			fmt.Printf("Response Type:  %s\n", meta.ResponseType)
		}
		if meta.AcceptHeader != "" {
			fmt.Printf("Accept:         %s\n", meta.AcceptHeader)
		}
		if meta.ContentType != "" {
			fmt.Printf("Content-Type:   %s\n", meta.ContentType)
		}

		fmt.Printf("RSQL Support:   %v\n", meta.SupportsRSQL)
		fmt.Printf("Paginated:      %v\n", meta.IsPaginated)

		if len(meta.Parameters) > 0 {
			fmt.Println("Parameters:")
			for _, param := range meta.Parameters {
				fmt.Printf("  - %s: %s\n", param.Name, param.Type)
			}
		}

		// Display scraped documentation
		if meta.ScrapedDocs != nil {
			fmt.Println("\n" + strings.Repeat("~", 80))
			fmt.Println("SCRAPED DOCUMENTATION")
			fmt.Println(strings.Repeat("~", 80))

			if meta.ScrapedDocs.Success {
				if meta.ScrapedDocs.Title != "" {
					fmt.Printf("Title:          %s\n", meta.ScrapedDocs.Title)
				}
				if meta.ScrapedDocs.Method != "" {
					fmt.Printf("HTTP Method:    %s\n", meta.ScrapedDocs.Method)
				}
				if meta.ScrapedDocs.Endpoint != "" {
					fmt.Printf("Endpoint:       %s\n", meta.ScrapedDocs.Endpoint)
				}
				if meta.ScrapedDocs.Description != "" {
					fmt.Printf("Description:    %s\n", meta.ScrapedDocs.Description)
				}

				if len(meta.ScrapedDocs.Parameters) > 0 {
					fmt.Println("\nAPI Parameters:")
					for _, param := range meta.ScrapedDocs.Parameters {
						fmt.Printf("  - %s: %s\n", param.Name, param.Description)
					}
				}

				if len(meta.ScrapedDocs.ResponseCodes) > 0 {
					fmt.Println("\nResponse Codes:")
					for _, rc := range meta.ScrapedDocs.ResponseCodes {
						fmt.Printf("  - %s: %s\n", rc.Code, rc.Description)
					}
				}

				if meta.ScrapedDocs.RequestBody != "" {
					fmt.Printf("\nRequest Body:\n%s\n", truncate(meta.ScrapedDocs.RequestBody, 500))
				}

				if meta.ScrapedDocs.ResponseBody != "" {
					fmt.Printf("\nResponse Body:\n%s\n", truncate(meta.ScrapedDocs.ResponseBody, 500))
				}

				fmt.Printf("\nScraped at: %s\n", meta.ScrapedDocs.ScrapedAt.Format(time.RFC3339))
			} else {
				fmt.Printf("❌ Failed to scrape: %s\n", meta.ScrapedDocs.Error)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

// printJSON outputs metadata in JSON format
func printJSON(metadata []FunctionMetadata) {
	// Simple JSON output
	fmt.Println("[")
	for i, meta := range metadata {
		fmt.Printf("  {\n")
		fmt.Printf("    \"functionName\": %q,\n", meta.FunctionName)
		fmt.Printf("    \"httpMethod\": %q,\n", meta.HTTPMethod)
		fmt.Printf("    \"endpoint\": %q,\n", meta.Endpoint)
		fmt.Printf("    \"apiDocsURL\": %q,\n", meta.APIDocsURL)
		fmt.Printf("    \"description\": %q,\n", meta.Description)
		fmt.Printf("    \"requestType\": %q,\n", meta.RequestType)
		fmt.Printf("    \"responseType\": %q,\n", meta.ResponseType)
		fmt.Printf("    \"supportsRSQL\": %v,\n", meta.SupportsRSQL)
		fmt.Printf("    \"isPaginated\": %v", meta.IsPaginated)

		// Add scraped documentation if available
		if meta.ScrapedDocs != nil {
			fmt.Printf(",\n    \"scrapedDocs\": {\n")
			fmt.Printf("      \"success\": %v,\n", meta.ScrapedDocs.Success)
			if meta.ScrapedDocs.Success {
				fmt.Printf("      \"title\": %q,\n", meta.ScrapedDocs.Title)
				fmt.Printf("      \"method\": %q,\n", meta.ScrapedDocs.Method)
				fmt.Printf("      \"endpoint\": %q,\n", meta.ScrapedDocs.Endpoint)
				fmt.Printf("      \"description\": %q,\n", meta.ScrapedDocs.Description)
				fmt.Printf("      \"requestBody\": %q,\n", meta.ScrapedDocs.RequestBody)
				fmt.Printf("      \"responseBody\": %q,\n", meta.ScrapedDocs.ResponseBody)

				// Parameters
				fmt.Printf("      \"parameters\": [")
				for j, param := range meta.ScrapedDocs.Parameters {
					if j > 0 {
						fmt.Printf(",")
					}
					fmt.Printf("\n        {\"name\": %q, \"description\": %q}", param.Name, param.Description)
				}
				fmt.Printf("\n      ],\n")

				// Response codes
				fmt.Printf("      \"responseCodes\": [")
				for j, rc := range meta.ScrapedDocs.ResponseCodes {
					if j > 0 {
						fmt.Printf(",")
					}
					fmt.Printf("\n        {\"code\": %q, \"description\": %q}", rc.Code, rc.Description)
				}
				fmt.Printf("\n      ],\n")

				fmt.Printf("      \"scrapedAt\": %q\n", meta.ScrapedDocs.ScrapedAt.Format(time.RFC3339))
			} else {
				fmt.Printf("      \"error\": %q\n", meta.ScrapedDocs.Error)
			}
			fmt.Printf("    }\n")
		} else {
			fmt.Printf("\n")
		}

		if i < len(metadata)-1 {
			fmt.Printf("  },\n")
		} else {
			fmt.Printf("  }\n")
		}
	}
	fmt.Println("]")
}

// printCSV outputs metadata in CSV format
func printCSV(metadata []FunctionMetadata) {
	// CSV header
	fmt.Println("Function,HTTP Method,Endpoint,Request Type,Response Type,RSQL,Paginated,API Docs")

	// CSV rows
	for _, meta := range metadata {
		fmt.Printf("%s,%s,%s,%s,%s,%v,%v,%s\n",
			meta.FunctionName,
			meta.HTTPMethod,
			meta.Endpoint,
			meta.RequestType,
			meta.ResponseType,
			meta.SupportsRSQL,
			meta.IsPaginated,
			meta.APIDocsURL,
		)
	}
}

// printScrapedText outputs scraped documentation in human-readable format
func printScrapedText(docs []*ScrapedDocumentation) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("API DOCUMENTATION SCRAPER REPORT")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("\nTotal APIs Scraped: %d\n\n", len(docs))

	for i, doc := range docs {
		fmt.Printf("\n[%d] %s\n", i+1, doc.URL)
		fmt.Println(strings.Repeat("-", 80))

		if doc.Success {
			if doc.Title != "" {
				fmt.Printf("Title:          %s\n", doc.Title)
			}
			if doc.Method != "" {
				fmt.Printf("HTTP Method:    %s\n", doc.Method)
			}
			if doc.Endpoint != "" {
				fmt.Printf("Endpoint:       %s\n", doc.Endpoint)
			}
			if doc.Description != "" {
				fmt.Printf("Description:    %s\n", doc.Description)
			}

			if len(doc.Parameters) > 0 {
				fmt.Println("\nParameters:")
				for _, param := range doc.Parameters {
					fmt.Printf("  - %s: %s\n", param.Name, param.Description)
				}
			}

			if len(doc.ResponseCodes) > 0 {
				fmt.Println("\nResponse Codes:")
				for _, rc := range doc.ResponseCodes {
					fmt.Printf("  - %s: %s\n", rc.Code, rc.Description)
				}
			}

			if doc.RequestBody != "" {
				fmt.Printf("\nRequest Body:\n%s\n", truncate(doc.RequestBody, 500))
			}

			if doc.ResponseBody != "" {
				fmt.Printf("\nResponse Body:\n%s\n", truncate(doc.ResponseBody, 500))
			}

			fmt.Printf("\nScraped at: %s\n", doc.ScrapedAt.Format(time.RFC3339))
		} else {
			fmt.Printf("❌ Failed to scrape: %s\n", doc.Error)
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

// printScrapedJSON outputs scraped documentation in JSON format
func printScrapedJSON(docs []*ScrapedDocumentation) {
	fmt.Println("[")
	for i, doc := range docs {
		fmt.Printf("  {\n")
		fmt.Printf("    \"url\": %q,\n", doc.URL)
		fmt.Printf("    \"success\": %v,\n", doc.Success)

		if doc.Success {
			fmt.Printf("    \"title\": %q,\n", doc.Title)
			fmt.Printf("    \"method\": %q,\n", doc.Method)
			fmt.Printf("    \"endpoint\": %q,\n", doc.Endpoint)
			fmt.Printf("    \"description\": %q,\n", doc.Description)
			fmt.Printf("    \"requestBody\": %q,\n", doc.RequestBody)
			fmt.Printf("    \"responseBody\": %q,\n", doc.ResponseBody)

			// Parameters
			fmt.Printf("    \"parameters\": [")
			for j, param := range doc.Parameters {
				if j > 0 {
					fmt.Printf(",")
				}
				fmt.Printf("\n      {\"name\": %q, \"type\": %q, \"required\": %v, \"description\": %q, \"location\": %q}",
					param.Name, param.Type, param.Required, param.Description, param.Location)
			}
			fmt.Printf("\n    ],\n")

			// Response codes
			fmt.Printf("    \"responseCodes\": [")
			for j, rc := range doc.ResponseCodes {
				if j > 0 {
					fmt.Printf(",")
				}
				fmt.Printf("\n      {\"code\": %q, \"description\": %q}", rc.Code, rc.Description)
			}
			fmt.Printf("\n    ],\n")

			fmt.Printf("    \"scrapedAt\": %q\n", doc.ScrapedAt.Format(time.RFC3339))
		} else {
			fmt.Printf("    \"error\": %q\n", doc.Error)
		}

		if i < len(docs)-1 {
			fmt.Printf("  },\n")
		} else {
			fmt.Printf("  }\n")
		}
	}
	fmt.Println("]")
}

// printScrapedCSV outputs scraped documentation in CSV format
func printScrapedCSV(docs []*ScrapedDocumentation) {
	// CSV header
	fmt.Println("URL,Success,Title,Method,Endpoint,Description,Parameters Count,Response Codes Count,Scraped At")

	// CSV rows
	for _, doc := range docs {
		if doc.Success {
			// Escape description for CSV
			desc := strings.ReplaceAll(doc.Description, "\"", "\"\"")
			fmt.Printf("%s,%v,%s,%s,%s,\"%s\",%d,%d,%s\n",
				doc.URL,
				doc.Success,
				doc.Title,
				doc.Method,
				doc.Endpoint,
				desc,
				len(doc.Parameters),
				len(doc.ResponseCodes),
				doc.ScrapedAt.Format(time.RFC3339),
			)
		} else {
			fmt.Printf("%s,%v,,,,,,,\n", doc.URL, doc.Success)
		}
	}
}
