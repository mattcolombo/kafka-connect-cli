$Env:GOARCH = "amd64"
$Env:GOOS = "linux"
go build ..

$Env:GOOS = "windows"
go build ..

$Env:GOOS = ""
$Env:GOARCH = ""