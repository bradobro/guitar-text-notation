test:
	go test -v

testw:
	watchexec -f '*.go' 'go test -v'
