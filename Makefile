all: clean scanner

clean:
	rm -fr *.go

scanner:
	protoc --go_out=plugins=grpc:. scanner/scanner.proto
