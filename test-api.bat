@echo off
setlocal enabledelayedexpansion
echo ===================================================
echo [TEST] Testing Unotek Backend API Endpoints
echo ===================================================
echo.

echo 1. Testing Public Endpoint: GET /api/v1/company/profile
curl -s -X GET http://localhost:8080/api/v1/company/profile
echo.
echo.

echo 2. Testing Super Admin Login: POST /api/v1/auth/login
curl -s -X POST http://localhost:8080/api/v1/auth/login ^
  -H "Content-Type: application/json" ^
  -d "{\"identifier\":\"admin@company.com\",\"password\":\"admin12345\"}"
echo.
echo.

echo 3. Testing Public Clients: GET /api/v1/clients
curl -s -X GET http://localhost:8080/api/v1/clients
echo.
echo.

echo 4. Testing Public Solutions (WMS/TMS): GET /api/v1/solutions
curl -s -X GET http://localhost:8080/api/v1/solutions
echo.
echo.

echo 5. Testing Swagger Documentation:
echo Open in browser: http://localhost:8080/swagger/index.html
echo.
echo ===================================================
echo [TEST] Completed!
echo ===================================================
pause
