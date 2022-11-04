# compile list
files = ocp/*.go

test:
	go test -v $(files)

clean:
	echo "usage: make test"

.PHONY: test clean
