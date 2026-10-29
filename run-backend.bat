@echo off
cd /d "%~dp0"
echo ==============================================
echo Starting Unotek Backend API Server...
echo ==============================================
echo Server will run immediately without migrations.
echo Interactive Swagger: http://localhost:8080/swagger/index.html
echo.
go run cmd/api/main.go
pause
