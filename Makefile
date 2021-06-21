build:
	go build  -o bin/voding_service ./cmd/voting

test:
	AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY go test -v