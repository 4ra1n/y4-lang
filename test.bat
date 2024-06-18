@ECHO OFF
SETLOCAL ENABLEDELAYEDEXPANSION

FOR %%G IN (y4-examples\*.y4) DO (
    ECHO running script for %%G
    go run cmd\main.go -quiet %%G
    IF ERRORLEVEL 1 (
        ECHO error occurred with %%G
        EXIT /B 1
    )
)

ECHO all scripts executed successfully
ENDLOCAL
