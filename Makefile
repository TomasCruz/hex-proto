all: clean scanner

.PHONY: clean
clean:
	rm -fr scanner/*.pb.go

.PHONY: scanner
scanner:
	protoc --go_out=plugins=grpc:. scanner/scanner.proto
