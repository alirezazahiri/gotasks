@echo off
setlocal enabledelayedexpansion

if "%~1"=="" goto help
if "%~1"=="help" goto help
if "%~1"=="up" goto up
if "%~1"=="down" goto down
if "%~1"=="version" goto version
if "%~1"=="steps" goto steps
if "%~1"=="force" goto force
if "%~1"=="create" goto create

:help
echo Available migration commands:
echo   migrate.bat up                  - Run all pending migrations
echo   migrate.bat down                - Rollback all migrations
echo   migrate.bat version             - Show current migration version
echo   migrate.bat steps N             - Run N migration steps (use negative for rollback)
echo   migrate.bat force V             - Force database to specific version
echo   migrate.bat create NAME         - Create new migration files
goto end

:up
echo Running migrations up...
go run cmd/migrate/main.go -cmd=up
goto end

:down
echo Rolling back migrations...
go run cmd/migrate/main.go -cmd=down
goto end

:version
go run cmd/migrate/main.go -cmd=version
goto end

:steps
if "%~2"=="" (
    echo Error: Steps number required. Usage: migrate.bat steps N
    exit /b 1
)
go run cmd/migrate/main.go -cmd=steps -steps=%~2
goto end

:force
if "%~2"=="" (
    echo Error: Version number required. Usage: migrate.bat force V
    exit /b 1
)
go run cmd/migrate/main.go -cmd=force -version=%~2
goto end

:create
if "%~2"=="" (
    echo Error: Migration name required. Usage: migrate.bat create NAME
    exit /b 1
)
set timestamp=%date:~-4%%date:~-10,2%%date:~-7,2%_%time:~0,2%%time:~3,2%%time:~6,2%
set timestamp=%timestamp: =0%
set up_file=migrations\%timestamp%_%~2.up.sql
set down_file=migrations\%timestamp%_%~2.down.sql
type nul > "%up_file%"
type nul > "%down_file%"
echo Created migration files:
echo   %up_file%
echo   %down_file%
goto end

:end
endlocal

