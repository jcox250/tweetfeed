build:
	go build *.go

install: 
	go install
	cp $(GOBIN)/tweets /usr/local/bin

