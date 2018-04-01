set GOPATH=%~dp0\\..\\..\\..\\..
set GOBIN=%~dp0\\bin
echo %GOPATH%
echo %GOBIN%

set CURPATH=%~dp0

cd benchmarks
call go test -test.bench=".*" -count=1

cd %CURPATH%