$Env:GOARCH = "amd64"
$Env:GOOS = "linux"
go build -o kconnect-cli

$Env:GOOS = "windows"
go build -o kconnect-cli.exe

$Env:GOOS = ""
$Env:GOARCH = ""