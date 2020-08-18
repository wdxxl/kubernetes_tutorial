APP_NAME=sample

.PHONY: sonar-with-test-coverage
sonar-with-test-coverage: sonar-scanner
  $(shell rm -rf reports)
  $(shell mkdir -p reports)
  $(shell go test ./... -coverprofile=reports/coverage.out -json > reports/test.json)

## 还有一些问题
.PHONY: sonar-scanner
sonar-scanner:
  $(shell /Users/wdxxl/Mac_Software/sonar-scanner-4.4.0.2170-macosx/bin/sonar-scanner)
