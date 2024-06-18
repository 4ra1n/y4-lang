@echo off
setlocal enabledelayedexpansion
set arg=%1

if "%arg%"=="clean" (
    del build-url.txt
    go run .\delete-runs\main.go
    go run .\delete-caches\main.go
) else if "%arg%"=="build" (
    go run .\build\main.go
    set /p url=<build-url.txt
    "C:\Program Files\Google\Chrome\Application\chrome.exe" !url!
)
