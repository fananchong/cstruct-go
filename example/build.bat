set GOPATH=%~dp0\\..\\..\\..\\..\\..\\
set GOBIN=%~dp0\\
echo %GOPATH%
echo %GOBIN%
go install -race main.go
pause