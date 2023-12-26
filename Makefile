db: 
	docker start scylla

killdb:
	docker kill scylla

cqlsh:
	docker exec -it scylla cqlsh

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

.PHONY: db killdb cqlsh proto
