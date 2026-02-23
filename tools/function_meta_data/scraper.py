#!/usr/bin/env python3
"""
Jamf Pro API Documentation Scraper

This tool extracts API documentation URLs from Go source files and scrapes
the Jamf Developer Portal to gather comprehensive API metadata.
"""

import argparse
import json
import re
import sys
import time
from datetime import datetime
from pathlib import Path
from typing import List, Dict, Optional
from urllib.parse import urlparse

from playwright.sync_api import sync_playwright, Page, Browser


class APIDocumentation:
    """Holds scraped API documentation data"""

    def __init__(self, url: str):
        self.url = url
        self.success = False
        self.error = None
        self.scraped_at = datetime.now().isoformat()

        # Scraped data
        self.title = ""
        self.method = ""
        self.endpoint = ""
        self.description = ""
        self.request_body = ""
        self.response_body = ""
        self.parameters = []
        self.response_codes = []
        self.raw_html = ""


def extract_doc_urls(file_path: str, debug: bool = False) -> List[str]:
    """Extract Jamf Pro API documentation URLs from a Go source file"""

    if debug:
        print(f"🔍 Extracting API documentation URLs from: {file_path}", file=sys.stderr)

    try:
        with open(file_path, 'r') as f:
            content = f.read()
    except Exception as e:
        print(f"❌ Error reading file: {e}", file=sys.stderr)
        return []

    # Extract URLs using regex
    url_pattern = r'https://developer\.jamf\.com/jamf-pro/reference/[^\s\)]+'
    urls = re.findall(url_pattern, content)

    # Remove duplicates while preserving order
    seen = set()
    unique_urls = []
    for url in urls:
        if url not in seen:
            seen.add(url)
            unique_urls.append(url)
            if debug:
                print(f"  📎 Found URL: {url}", file=sys.stderr)

    if debug:
        print(f"🔍 Found {len(unique_urls)} unique API documentation URLs", file=sys.stderr)

    return unique_urls


def scrape_documentation(page: Page, url: str, debug: bool = False) -> APIDocumentation:
    """Scrape API documentation from a URL using Playwright"""

    doc = APIDocumentation(url)

    if debug:
        print(f"\n🌐 Scraping documentation from: {url}", file=sys.stderr)

    try:
        # Navigate to the page
        if debug:
            print("  🔄 Loading page...", file=sys.stderr)

        response = page.goto(url, wait_until="networkidle", timeout=60000)

        if response.status != 200:
            doc.error = f"HTTP {response.status}"
            if debug:
                print(f"  ❌ Error: {doc.error}", file=sys.stderr)
            return doc

        # Wait for content to render
        time.sleep(2)

        # Get the rendered HTML
        doc.raw_html = page.content()

        if debug:
            print(f"  📄 Rendered HTML Length: {len(doc.raw_html)} bytes", file=sys.stderr)

        # Extract title
        title_elem = page.query_selector('h1')
        if title_elem:
            doc.title = title_elem.inner_text().strip()
            if debug:
                print(f"  ✓ Title: {doc.title}", file=sys.stderr)

        # Extract endpoint and method
        # Look for code blocks or specific elements that show the endpoint
        code_blocks = page.query_selector_all('code')
        for code in code_blocks:
            text = code.inner_text().strip()
            # Check if it looks like an HTTP method + endpoint
            if re.match(r'^(GET|POST|PUT|PATCH|DELETE)\s+/api/', text, re.IGNORECASE):
                parts = text.split(None, 1)
                if len(parts) == 2:
                    doc.method = parts[0].upper()
                    doc.endpoint = parts[1]
                    if debug:
                        print(f"  ✓ Endpoint: {doc.method} {doc.endpoint}", file=sys.stderr)
                    break

        # Extract description
        desc_elem = page.query_selector('meta[name="description"]')
        if desc_elem:
            doc.description = desc_elem.get_attribute('content') or ""

        if not doc.description:
            # Try to find description in the page
            paragraphs = page.query_selector_all('p')
            if paragraphs:
                doc.description = paragraphs[0].inner_text().strip()

        if debug and doc.description:
            print(f"  ✓ Description: {doc.description[:100]}...", file=sys.stderr)

        # Extract parameters from tables
        tables = page.query_selector_all('table')
        for table in tables:
            rows = table.query_selector_all('tr')
            for row in rows:
                cells = row.query_selector_all('td')
                if len(cells) >= 2:
                    param_name = cells[0].inner_text().strip()
                    param_desc = cells[1].inner_text().strip()

                    # Skip header rows
                    if param_name.lower() in ['name', 'parameter', 'field']:
                        continue

                    if param_name:
                        doc.parameters.append({
                            'name': param_name,
                            'description': param_desc
                        })

        if debug:
            print(f"  📊 Parameters: {len(doc.parameters)}", file=sys.stderr)

        # Extract response codes
        # Look for patterns like "200 OK", "201 Created", etc.
        response_code_pattern = r'(\d{3})\s+([A-Za-z\s]+)'
        matches = re.findall(response_code_pattern, doc.raw_html)

        seen_codes = set()
        for code, desc in matches:
            if 200 <= int(code) <= 599 and code not in seen_codes:
                doc.response_codes.append({
                    'code': code,
                    'description': desc.strip()
                })
                seen_codes.add(code)

        if debug:
            print(f"  📨 Response codes: {len(doc.response_codes)}", file=sys.stderr)

        # Extract request/response schemas
        # Look for JSON code blocks
        pre_elements = page.query_selector_all('pre')
        for pre in pre_elements:
            code = pre.inner_text().strip()
            if code.startswith('{') and len(code) > 50:
                # This might be a JSON schema
                if not doc.request_body and 'request' in pre.get_attribute('class') or '':
                    doc.request_body = code
                elif not doc.response_body:
                    doc.response_body = code

        doc.success = True

        if debug:
            print("  ✅ Successfully scraped documentation", file=sys.stderr)

    except Exception as e:
        doc.error = str(e)
        if debug:
            print(f"  ❌ Error: {doc.error}", file=sys.stderr)

    return doc


def save_html_dump(doc: APIDocumentation, output_dir: Path, debug: bool = False):
    """Save the raw HTML to a file"""

    # Create filename from URL
    filename = re.sub(r'[^a-zA-Z0-9-]', '_', doc.url) + '.html'
    filepath = output_dir / filename

    try:
        with open(filepath, 'w', encoding='utf-8') as f:
            f.write(doc.raw_html)
        if debug:
            print(f"  💾 Saved HTML to: {filepath}", file=sys.stderr)
    except Exception as e:
        if debug:
            print(f"  ⚠️  Failed to save HTML: {e}", file=sys.stderr)


def main():
    parser = argparse.ArgumentParser(
        description='Scrape Jamf Pro API documentation from developer portal'
    )
    parser.add_argument(
        '-file',
        '--file',
        required=True,
        help='Path to the Go source file to parse'
    )
    parser.add_argument(
        '-format',
        '--format',
        choices=['text', 'json', 'csv'],
        default='json',
        help='Output format (default: json)'
    )
    parser.add_argument(
        '-debug',
        '--debug',
        action='store_true',
        help='Enable debug output'
    )
    parser.add_argument(
        '-save-html',
        '--save-html',
        action='store_true',
        help='Save raw HTML files for inspection'
    )
    parser.add_argument(
        '-headless',
        '--headless',
        action='store_true',
        default=True,
        help='Run browser in headless mode (default: True)'
    )

    args = parser.parse_args()

    # Extract URLs from the Go file
    urls = extract_doc_urls(args.file, args.debug)

    if not urls:
        print("❌ No documentation URLs found", file=sys.stderr)
        sys.exit(1)

    # Prepare HTML dumps directory if needed
    if args.save_html:
        html_dumps_dir = Path('html_dumps')
        html_dumps_dir.mkdir(exist_ok=True)
        if args.debug:
            print(f"💾 Will save raw HTML files to {html_dumps_dir}/", file=sys.stderr)

    # Initialize Playwright and scrape
    if args.debug:
        print("🎭 Initializing Playwright...", file=sys.stderr)

    scraped_docs = []

    with sync_playwright() as p:
        # Launch browser
        browser = p.chromium.launch(headless=args.headless)

        if args.debug:
            print("✅ Browser launched successfully", file=sys.stderr)

        # Create a new page
        page = browser.new_page()

        # Scrape each URL
        for url in urls:
            doc = scrape_documentation(page, url, args.debug)
            scraped_docs.append(doc)

            if args.save_html and doc.raw_html:
                save_html_dump(doc, html_dumps_dir, args.debug)

        # Close browser
        browser.close()

    # Output results
    if args.format == 'json':
        output = []
        for doc in scraped_docs:
            output.append({
                'url': doc.url,
                'success': doc.success,
                'error': doc.error,
                'title': doc.title,
                'method': doc.method,
                'endpoint': doc.endpoint,
                'description': doc.description,
                'requestBody': doc.request_body,
                'responseBody': doc.response_body,
                'parameters': doc.parameters,
                'responseCodes': doc.response_codes,
                'scrapedAt': doc.scraped_at
            })
        print(json.dumps(output, indent=2))

    elif args.format == 'csv':
        print("URL,Success,Title,Method,Endpoint,Description,Parameters,Response Codes")
        for doc in scraped_docs:
            desc = doc.description.replace('"', '""')[:100]
            print(f'{doc.url},{doc.success},"{doc.title}",{doc.method},{doc.endpoint},"{desc}",{len(doc.parameters)},{len(doc.response_codes)}')

    else:  # text
        print("=" * 80)
        print("API DOCUMENTATION SCRAPER REPORT")
        print("=" * 80)
        print(f"\nTotal APIs Scraped: {len(scraped_docs)}\n")

        for i, doc in enumerate(scraped_docs, 1):
            print(f"\n[{i}] {doc.url}")
            print("-" * 80)
            if doc.success:
                if doc.title:
                    print(f"Title:       {doc.title}")
                if doc.method and doc.endpoint:
                    print(f"Endpoint:    {doc.method} {doc.endpoint}")
                if doc.description:
                    print(f"Description: {doc.description[:200]}")
                if doc.parameters:
                    print(f"\nParameters: {len(doc.parameters)}")
                    for param in doc.parameters[:5]:
                        print(f"  - {param['name']}: {param['description'][:50]}")
                if doc.response_codes:
                    print(f"\nResponse Codes: {len(doc.response_codes)}")
                    for rc in doc.response_codes[:5]:
                        print(f"  - {rc['code']}: {rc['description']}")
            else:
                print(f"❌ Failed: {doc.error}")

        print("\n" + "=" * 80)


if __name__ == '__main__':
    main()
