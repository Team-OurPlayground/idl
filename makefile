proto-data:
	protoc --go_out=. proto/data.proto

flat-buffer-chat:
	flatc --go --gen-object-api flatbuffers/chat.fbs