@echo off
SET GOCMD=go
SET GOBUILD=%GOCMD% build 
SET GOCLEAN=%GOCMD% clean
SET GOGET=%GOCMD% get
SET BINARY_NAME=app.exe

:all
call :clean
call :test
call :build
goto :eof

:build
echo Building executable
%GOBUILD% -o %BINARY_NAME% -v ./cmd/rest_api/
goto :eof

:test
echo Running designated tests
%GOCMD% test -v ./...
goto :eof

:clean
echo Removing previous configurations
%GOCLEAN%
del %BINARY_NAME%
goto :eof


