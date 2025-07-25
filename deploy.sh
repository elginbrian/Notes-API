#!/bin/bash

# Notes API Deployment Script

echo "🚀 Starting Notes API deployment..."

# Step 1: Clean up old containers and images
echo "🧹 Cleaning up old containers..."
docker-compose down --remove-orphans
docker system prune -f

# Step 2: Generate Swagger documentation
echo "📖 Generating Swagger documentation..."
swag init

# Step 3: Build and start services
echo "🏗️ Building and starting services..."
docker-compose up --build --no-cache -d

# Step 4: Wait for services to be ready
echo "⏳ Waiting for services to start..."
sleep 10

# Step 5: Test endpoints
echo "🧪 Testing API endpoints..."

echo "Testing health endpoint..."
curl -s https://notes.elginbrian.com/health || echo "❌ Health check failed"

echo "Testing root endpoint..."
curl -s https://notes.elginbrian.com/ || echo "❌ Root endpoint failed"

echo "Testing Swagger debug endpoint..."
curl -s https://notes.elginbrian.com/debug/docs || echo "❌ Debug endpoint failed"

echo "Testing Swagger UI..."
curl -s https://notes.elginbrian.com/swagger/ || echo "❌ Swagger UI failed"

# Step 6: Show logs
echo "📋 Recent logs:"
docker-compose logs --tail=20 app

echo "✅ Deployment complete!"
echo "🌐 API URL: https://notes.elginbrian.com/"
echo "📖 Swagger UI: https://notes.elginbrian.com/swagger/"
echo "🏥 Health Check: https://notes.elginbrian.com/health"
