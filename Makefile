build:
	go build  -o bin/voding_service ./cmd/voting
	go build  -o bin/voding_client ./cmd/voting_client

test:
	AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY go test -v