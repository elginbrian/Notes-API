# Notes API Deployment Script (PowerShell)

Write-Host "🚀 Starting Notes API deployment..." -ForegroundColor Green

# Step 1: Clean up old containers and images
Write-Host "🧹 Cleaning up old containers..." -ForegroundColor Yellow
docker-compose down --remove-orphans
docker system prune -f

# Step 2: Generate Swagger documentation
Write-Host "📖 Generating Swagger documentation..." -ForegroundColor Yellow
swag init

# Step 3: Build and start services
Write-Host "🏗️ Building and starting services..." -ForegroundColor Yellow
docker-compose up --build --no-cache -d

# Step 4: Wait for services to be ready
Write-Host "⏳ Waiting for services to start..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

# Step 5: Test endpoints
Write-Host "🧪 Testing API endpoints..." -ForegroundColor Yellow

Write-Host "Testing health endpoint..."
try { 
    $response = Invoke-RestMethod -Uri "https://notes.elginbrian.com/health" -Method Get
    Write-Host "✅ Health check passed" -ForegroundColor Green
} catch {
    Write-Host "❌ Health check failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "Testing root endpoint..."
try { 
    $response = Invoke-RestMethod -Uri "https://notes.elginbrian.com/" -Method Get
    Write-Host "✅ Root endpoint passed" -ForegroundColor Green
} catch {
    Write-Host "❌ Root endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "Testing Swagger debug endpoint..."
try { 
    $response = Invoke-RestMethod -Uri "https://notes.elginbrian.com/debug/docs" -Method Get
    Write-Host "✅ Debug endpoint passed" -ForegroundColor Green
} catch {
    Write-Host "❌ Debug endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}

# Step 6: Show logs
Write-Host "📋 Recent logs:" -ForegroundColor Yellow
docker-compose logs --tail=20 app

Write-Host "✅ Deployment complete!" -ForegroundColor Green
Write-Host "🌐 API URL: https://notes.elginbrian.com/" -ForegroundColor Cyan
Write-Host "📖 Swagger UI: https://notes.elginbrian.com/swagger/" -ForegroundColor Cyan
Write-Host "🏥 Health Check: https://notes.elginbrian.com/health" -ForegroundColor Cyan
