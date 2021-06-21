package GolangTechTask

//go:generate protoc -I api/ --go_out=plugins=grpc:./api api/service.proto
