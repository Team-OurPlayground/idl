proto-test:
	protoc --go_out=. --go_opt=paths=source_relative proto/test.proto