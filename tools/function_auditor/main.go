package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FunctionIssue struct {
	FilePath           string   `json:"file_path"`
	FunctionName       string   `json:"function_name"`
	HasComment         bool     `json:"has_comment"`
	CommentStartsRight bool     `json:"comment_starts_right"`
	Comment            string   `json:"comment"`
	DocURLs            []string `json:"doc_urls"`
	Issues             []string `json:"issues"`
}

type AuditReport struct {
	TotalFunctions                int              `json:"total_functions"`
	FunctionsWithMissingComments  int              `json:"functions_with_missing_comments"`
	FunctionsWithWrongFormat      int              `json:"functions_with_wrong_format"`
	FunctionsWithMissingDocURLs   int              `json:"functions_with_missing_doc_urls"`
	AllDocURLs                    []string         `json:"all_doc_urls"`
	Issues                        []FunctionIssue  `json:"issues"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory_pattern>")
		os.Exit(1)
	}

	pattern := os.Args[1]
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding files: %v\n", err)
		os.Exit(1)
	}

	report := AuditReport{
		Issues: []FunctionIssue{},
		AllDocURLs: []string{},
	}

	urlMap := make(map[string]bool)

	for _, file := range files {
		analyzeFile(file, &report, urlMap)
	}

	// Convert URL map to slice
	for url := range urlMap {
		report.AllDocURLs = append(report.AllDocURLs, url)
	}

	// Output as JSON
	output, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func analyzeFile(filePath string, report *AuditReport, urlMap map[string]bool) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file %s: %v\n", filePath, err)
		return
	}

	// Iterate through all declarations
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Only analyze exported functions (receiver methods)
		if funcDecl.Name == nil || !funcDecl.Name.IsExported() {
			continue
		}

		// Skip if it doesn't have a receiver (we want methods only)
		if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
			continue
		}

		report.TotalFunctions++

		issue := FunctionIssue{
			FilePath:     filePath,
			FunctionName: funcDecl.Name.Name,
			DocURLs:      []string{},
			Issues:       []string{},
		}

		// Get the function's doc comment
		if funcDecl.Doc == nil || len(funcDecl.Doc.List) == 0 {
			issue.HasComment = false
			issue.Issues = append(issue.Issues, "Missing comment")
			report.FunctionsWithMissingComments++
		} else {
			issue.HasComment = true
			fullComment := ""
			for _, comment := range funcDecl.Doc.List {
				fullComment += strings.TrimPrefix(comment.Text, "// ") + "\n"
			}
			issue.Comment = strings.TrimSpace(fullComment)

			// Check if comment starts with function name
			firstLine := funcDecl.Doc.List[0].Text
			firstLine = strings.TrimPrefix(firstLine, "//")
			firstLine = strings.TrimSpace(firstLine)

			if !strings.HasPrefix(firstLine, funcDecl.Name.Name) {
				issue.CommentStartsRight = false
				issue.Issues = append(issue.Issues, fmt.Sprintf("Comment should start with '%s' but starts with '%s'", funcDecl.Name.Name, getFirstWords(firstLine, 3)))
				report.FunctionsWithWrongFormat++
			} else {
				issue.CommentStartsRight = true
			}

			// Extract doc URLs
			urls := extractDocURLs(issue.Comment)
			if len(urls) == 0 {
				issue.Issues = append(issue.Issues, "Missing doc URL")
				report.FunctionsWithMissingDocURLs++
			} else {
				issue.DocURLs = urls
				for _, url := range urls {
					urlMap[url] = true
				}
			}
		}

		if len(issue.Issues) > 0 {
			report.Issues = append(report.Issues, issue)
		}
	}
}

func extractDocURLs(comment string) []string {
	// Look for URLs in comments
	urlRegex := regexp.MustCompile(`https?://[^\s\)]+`)
	matches := urlRegex.FindAllString(comment, -1)

	// Filter to only include Jamf documentation URLs
	var docURLs []string
	for _, url := range matches {
		if strings.Contains(url, "developer.jamf.com") ||
		   strings.Contains(url, "docs.jamf.com") ||
		   strings.Contains(url, "jamf.com") {
			docURLs = append(docURLs, url)
		}
	}

	return docURLs
}

func getFirstWords(s string, n int) string {
	words := strings.Fields(s)
	if len(words) <= n {
		return s
	}
	return strings.Join(words[:n], " ") + "..."
}
