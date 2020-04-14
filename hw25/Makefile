.PHONY: build
build:
	go build -o ./bin/calendar_api ./cmd/calendar_api.go
	go build -o ./bin/calendar_scheduler ./cmd/calendar_scheduler.go
	go build -o ./bin/client ./cmd/client.go
	go build -o ./bin/calendar_sender ./internal/calendar_sender/calendar_sender.go

test:
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage 2>&1 | go-junit-report > ./report/storage.xml
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq 2>&1 | go-junit-report > ./report/rabbitmq.xml
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/calendar 2>&1 | go-junit-report > ./report/calendar.xml
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/scheduler 2>&1 | go-junit-report > ./report/scheduler.xml
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/grpc 2>&1 | go-junit-report > ./report/grpc.xml

build_grpc_unix:
	./build/createproto.sh

build_grpc_win:
	./build/createproto.bat

get_junit-report:
	go	get -u github.com/jstemmer/go-junit-report

help:
	@echo build  - create programm
	@echo test - run unit test
	@echo get_junit-report - get unit test
	@echo build_grpc_unix - protobuf in unix
	@echo build_grpc_win -  protobuf in windows