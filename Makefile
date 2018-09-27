.PHONY: pb

doc1:
	cd gateway && swag init

doc2:
	swag init -g doc.go

pb:
	@rm -f pb/*.pb.go
	protoc -I=.:pb/extra/src --go_out=plugins=grpc:. pb/*.proto

