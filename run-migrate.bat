@echo off
cd /d "%~dp0"
echo ==============================================
echo Unotek Database Migration CLI
echo ==============================================
echo [1] Run Migrations only (up)
echo [2] Run Seeder only (seed)
echo [3] Run Migration + Seeder (all)
echo [4] Drop Tables (drop)
echo.
set /p opt="Pilih opsi (1/2/3/4) [default 3]: "
if "%opt%"=="1" goto UP
if "%opt%"=="2" goto SEED
if "%opt%"=="4" goto DROP
goto ALL

:UP
go run cmd/migrate/main.go up
goto END

:SEED
go run cmd/migrate/main.go seed
goto END

:ALL
go run cmd/migrate/main.go all
goto END

:DROP
go run cmd/migrate/main.go drop
goto END

:END
pause
