build:
	cd sum; \
	GOOS=linux go build main.go; \
	cd ..; \
	mkdir -p out; \
	mv sum/main out;

deploy: build
	sam validate
	sam deploy --guided

invoke:
	sam local invoke -e events/event.json SumAppFunction

test:
	cd sum; AWS_REGION='us-east2' go test;
