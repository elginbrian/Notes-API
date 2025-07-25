#!/bin/bash

# Route Testing Script for Notes API

echo "🧪 Testing Notes API Routes"
echo "=================================="

BASE_URL="https://notes.elginbrian.com"

echo "1. Testing root endpoint..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/" | head -1

echo -e "\n2. Testing health endpoint..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/health" | head -1

echo -e "\n3. Testing debug endpoint..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/debug/docs" | head -1

echo -e "\n4. Testing Swagger redirect..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/swagger" | head -1

echo -e "\n5. Testing Swagger UI..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/swagger/" | head -1

echo -e "\n6. Testing Swagger index page..."
curl -s -w "Status: %{http_code}\n" "$BASE_URL/swagger/index.html" | head -1

echo -e "\n=================================="
echo "Expected results:"
echo "- Root, health, debug: 200"
echo "- /swagger: 302 (redirect)"
echo "- /swagger/: 200" 
echo "- /swagger/index.html: 200"
