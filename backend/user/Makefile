.PHONY:run
run:
	make -j2 api rpc;
.PHONY: api
api:
	cd ./api && go run user.go
.PHONY: rpc
rpc:
	cd ./rpc && go run user.go