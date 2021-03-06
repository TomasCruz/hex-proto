all: clean scanner translator

.PHONY: clean
clean:
	rm -fr scanner/*.pb.go
	rm -fr translator/*.pb.go

.PHONY: scanner
scanner:
	protoc --go_out=plugins=grpc:. scanner/scanner.proto

.PHONY: translator
translator:
	protoc --go_opt=Mscanner/scanner.proto=github.com/TomasCruz/hex-proto/scanner --go_out=plugins=grpc:. translator/translator.proto
