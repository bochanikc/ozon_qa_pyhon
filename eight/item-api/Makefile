.PHONY: version
version:
	go version
	protoc --version
	docker -v
	goose -version

.PHONY: generate
generate:
	protoc -I ./ \
    	--go_out ./pkg --go_opt paths=source_relative \
    	--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
    	--grpc-gateway_out ./pkg \
        --grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
    	--openapiv2_out ./pkg \
    	--openapiv2_opt logtostderr=true \
    	api/item.proto

.PHONY: database
database:
	docker run --name postgres -p 5432:5432 -e POSTGRES_DB=db -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -d postgres

.PHONY: migration
migration:
	cd migrations && goose postgres "user=test password=test dbname=db sslmode=disable" up

.PHONY: run
run:
	go run ./cmd/item/main.go

.PHONY: stop
stop:
	docker stop postgres

.PHONY: clean
clean:
	docker rm postgres

.PHONY: test
test:
	go test ./... -v

.PHONY: integration-test
integration-test:
	go test ./test/tests/... -tags integration -count=1 -v