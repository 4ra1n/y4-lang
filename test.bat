@ECHO OFF
SETLOCAL ENABLEDELAYEDEXPANSION

FOR %%G IN (examples\*.y4) DO (
    ECHO running script for %%G
    go run cmd\main.go -f %%G
    IF ERRORLEVEL 1 (
        ECHO error occurred with %%G
        EXIT /B 1
    )
)

ECHO all scripts executed successfully
ENDLOCAL
