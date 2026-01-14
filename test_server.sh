#!/bin/bash

# HTTP Server Test Script
# This script tests all the endpoints of the HTTP server

echo "==================================="
echo "HTTP Server Test Suite"
echo "==================================="
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test 1: Root endpoint
echo -e "${BLUE}Test 1: GET / (200 OK)${NC}"
curl -s -o /dev/null -w "Status: %{http_code}\n" http://localhost:4221/
echo ""

# Test 2: Echo endpoint
echo -e "${BLUE}Test 2: GET /echo/hello (Echo message)${NC}"
curl -s http://localhost:4221/echo/hello
echo ""
echo ""

# Test 3: 404 for unknown path
echo -e "${BLUE}Test 3: GET /unknown (404 Not Found)${NC}"
curl -s -o /dev/null -w "Status: %{http_code}\n" http://localhost:4221/unknown
echo ""

# Test 4: User-Agent header
echo -e "${BLUE}Test 4: GET /user-agent (Return User-Agent)${NC}"
curl -s http://localhost:4221/user-agent
echo ""
echo ""

# Test 5: Echo with gzip compression
echo -e "${BLUE}Test 5: GET /echo/compressed with gzip${NC}"
curl -s -H "Accept-Encoding: gzip" http://localhost:4221/echo/compressed --compressed
echo ""
echo ""

# Test 6: File operations (if directory is specified)
echo -e "${YELLOW}File operations tests require --directory flag${NC}"
echo -e "${YELLOW}Start server with: go run . --directory /tmp${NC}"
echo ""

echo -e "${GREEN}==================================="
echo -e "Test suite completed!"
echo -e "===================================${NC}"
