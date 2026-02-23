.PHONY: build clean test run-example run-policies run-computer-inventory help

# Build the tool
build:
	@echo "🔨 Building function_meta_data tool..."
	go build -o function_meta_data .
	@echo "✅ Build complete: ./function_meta_data"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -f function_meta_data
	rm -f *.json *.csv
	@echo "✅ Clean complete"

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Run example on computer_inventory service
run-computer-inventory: build
	@echo "📊 Analyzing computer_inventory service..."
	./function_meta_data \
		-file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
		-debug

# Run on policies service
run-policies: build
	@echo "📊 Analyzing policies service..."
	./function_meta_data \
		-file ../../jamfpro/services/classic_api/policies/crud.go \
		-debug

# Generate JSON report for computer_inventory
json-report: build
	@echo "📄 Generating JSON report..."
	./function_meta_data \
		-file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
		-format json > computer_inventory_metadata.json
	@echo "✅ Report saved to: computer_inventory_metadata.json"

# Generate CSV report for computer_inventory
csv-report: build
	@echo "📊 Generating CSV report..."
	./function_meta_data \
		-file ../../jamfpro/services/jamf_pro_api/computer_inventory/crud.go \
		-format csv > computer_inventory_metadata.csv
	@echo "✅ Report saved to: computer_inventory_metadata.csv"

# Show help
help:
	@echo "Available targets:"
	@echo "  build                    - Build the function_meta_data tool"
	@echo "  clean                    - Remove build artifacts"
	@echo "  test                     - Run tests"
	@echo "  run-computer-inventory   - Analyze computer_inventory service with debug"
	@echo "  run-policies             - Analyze policies service with debug"
	@echo "  json-report              - Generate JSON metadata report"
	@echo "  csv-report               - Generate CSV metadata report"
	@echo "  help                     - Show this help message"
