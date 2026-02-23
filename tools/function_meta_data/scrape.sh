#!/bin/bash

# All-in-One API Documentation Scraper
# This script handles everything: setup, installation, and execution

set -e  # Exit on error

echo "╔════════════════════════════════════════════════════════════════════╗"
echo "║     Jamf Pro API Documentation Scraper - Automated Setup & Run    ║"
echo "╚════════════════════════════════════════════════════════════════════╝"
echo ""

# Change to script directory
cd "$(dirname "$0")"

# ============================================================================
# Step 1: Find appropriate Python version
# ============================================================================
echo "🐍 Step 1: Checking Python installation..."

PYTHON_CMD=""
if command -v python3.13 &> /dev/null; then
    PYTHON_CMD="python3.13"
elif command -v python3.12 &> /dev/null; then
    PYTHON_CMD="python3.12"
elif command -v python3.11 &> /dev/null; then
    PYTHON_CMD="python3.11"
elif command -v python3 &> /dev/null; then
    PYTHON_CMD="python3"
else
    echo "❌ ERROR: Python 3 is not installed."
    echo "   Install with: brew install python@3.13"
    exit 1
fi

PYTHON_VERSION=$($PYTHON_CMD --version)
echo "   ✅ Found: $PYTHON_VERSION ($PYTHON_CMD)"

# Warn about Python 3.14+
if [[ "$PYTHON_VERSION" == *"3.14"* ]] || [[ "$PYTHON_VERSION" == *"3.15"* ]]; then
    echo "   ⚠️  WARNING: Python 3.14+ may have compatibility issues"
    echo "   Recommended: brew install python@3.13"
fi
echo ""

# ============================================================================
# Step 2: Setup virtual environment
# ============================================================================
echo "📦 Step 2: Setting up virtual environment..."

if [ -d "venv" ]; then
    echo "   🧹 Removing old virtual environment..."
    rm -rf venv
fi

echo "   📂 Creating new virtual environment..."
$PYTHON_CMD -m venv venv

echo "   ✅ Virtual environment created"
echo ""

# ============================================================================
# Step 3: Install Python dependencies
# ============================================================================
echo "📥 Step 3: Installing Python dependencies..."

# Activate virtual environment
source venv/bin/activate

echo "   📌 Upgrading pip..."
pip install --upgrade pip --quiet

echo "   📦 Installing Playwright $(cat requirements.txt | grep playwright | cut -d'=' -f3)..."
pip install -r requirements.txt --quiet

echo "   ✅ Dependencies installed"
echo ""

# ============================================================================
# Step 4: Install Playwright browsers
# ============================================================================
echo "🎭 Step 4: Installing Playwright browsers..."

playwright install chromium

echo "   ✅ Browser installed"
echo ""

# ============================================================================
# Step 5: Run the scraper
# ============================================================================
echo "🚀 Step 5: Running API documentation scraper..."
echo ""

# Default file path (can be overridden with $1)
FILE_PATH="${1:-/Users/dafyddwatkins/GitHub/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory/crud.go}"

if [ ! -f "$FILE_PATH" ]; then
    echo "❌ ERROR: File not found: $FILE_PATH"
    echo "   Usage: ./scrape.sh [path/to/crud.go]"
    exit 1
fi

echo "📂 Input file: $FILE_PATH"
echo "📄 Output: output.json"
echo "📋 Debug log: debug.log"
echo "💾 HTML dumps: html_dumps/"
echo ""

# Run the scraper
python scraper.py \
  -file "$FILE_PATH" \
  -format json \
  -debug \
  -save-html \
  > output.json 2> debug.log

# ============================================================================
# Step 6: Report results
# ============================================================================
echo ""
echo "╔════════════════════════════════════════════════════════════════════╗"
echo "║                         ✅ SCRAPING COMPLETE!                       ║"
echo "╚════════════════════════════════════════════════════════════════════╝"
echo ""

# Count results
if [ -f "output.json" ]; then
    API_COUNT=$(grep -o '"url"' output.json | wc -l | tr -d ' ')
    echo "📊 Results:"
    echo "   • APIs scraped: $API_COUNT"
    echo "   • JSON output: output.json"
    echo "   • Debug log: debug.log"

    if [ -d "html_dumps" ]; then
        HTML_COUNT=$(ls -1 html_dumps/*.html 2>/dev/null | wc -l | tr -d ' ')
        echo "   • HTML files: $HTML_COUNT in html_dumps/"
    fi

    echo ""
    echo "📖 View results:"
    echo "   cat output.json | jq '.'"
    echo "   cat debug.log | less"
else
    echo "⚠️  WARNING: output.json not created"
    echo "   Check debug.log for errors:"
    echo "   cat debug.log"
fi

echo ""
echo "🔄 To run again:"
echo "   ./scrape.sh [path/to/crud.go]"
echo ""
