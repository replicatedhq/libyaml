.PHONY: test shell deps

test:
	go test -v -covermode=count .

shell:
	docker run --rm -it --name libyaml \
	  -v "`pwd`:/go/src/github.com/replicatedhq/libyaml" \
	  libyaml

deps:
	go get -t .
