
go mod init sample

go test ./... -coverprofile=reports/coverage.out -json > reports/test.json