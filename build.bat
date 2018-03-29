set GOPATH=%~dp0\\..\\..\\..\\..
set GOBIN=%~dp0\\bin
echo %GOPATH%
echo %GOBIN%
go install -race ./...
pause