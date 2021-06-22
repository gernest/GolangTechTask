build:
	go build  -o bin/voting_service ./cmd/voting
	go build  -o bin/voting_client ./cmd/voting_client

test:
	AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY go test -v

service: build
	AWS_ACCESS_KEY_ID=DUMMYIDEXAMPLE AWS_SECRET_ACCESS_KEY=DUMMYEXAMPLEKEY ./bin/voting_service

client: build
	./bin/voting_client