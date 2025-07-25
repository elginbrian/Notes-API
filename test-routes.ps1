# Route Testing Script for Notes API

Write-Host "🧪 Testing Notes API Routes" -ForegroundColor Green
Write-Host "==================================" -ForegroundColor Green

$BaseURL = "https://notes.elginbrian.com"

Write-Host "1. Testing root endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/" -Method Get
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
}

Write-Host "`n2. Testing health endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/health" -Method Get
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
}

Write-Host "`n3. Testing debug endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/debug/docs" -Method Get
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
}

Write-Host "`n4. Testing Swagger redirect..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/swagger" -Method Get -MaximumRedirection 0
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    if ($_.Exception.Response.StatusCode.value__ -eq 302) {
        Write-Host "Status: 302 (redirect - GOOD!)" -ForegroundColor Green
    } else {
        Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
    }
}

Write-Host "`n5. Testing Swagger UI..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/swagger/" -Method Get
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
}

Write-Host "`n6. Testing Swagger index page..." -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$BaseURL/swagger/index.html" -Method Get
    Write-Host "Status: $($response.StatusCode)" -ForegroundColor Green
} catch {
    Write-Host "Status: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
}

Write-Host "`n==================================" -ForegroundColor Green
Write-Host "Expected results:" -ForegroundColor Cyan
Write-Host "- Root, health, debug: 200" -ForegroundColor Cyan
Write-Host "- /swagger: 302 (redirect)" -ForegroundColor Cyan
Write-Host "- /swagger/: 200" -ForegroundColor Cyan
Write-Host "- /swagger/index.html: 200" -ForegroundColor Cyan
