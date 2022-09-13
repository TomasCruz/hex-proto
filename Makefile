all: clean scanner translator

.PHONY: clean
clean:
	rm -fr scanner/*.pb.go
	rm -fr translator/*.pb.go

.PHONY: scanner
scanner:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative scanner/scanner.proto
.PHONY: translator
translator:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative translator/translator.proto
