# compile list
files = ocp/*.go

run: 
	go run main.go

test:
	go test -v $(files)

clean:
	echo "usage: make test"

.PHONY: test clean
