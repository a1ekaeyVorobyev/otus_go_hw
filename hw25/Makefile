.PHONY: build
build:
	go build -o ./bin/calendar_api ./cmd/calendar_api.go
	go build -o ./bin/calendar_scheduler ./cmd/calendar_scheduler.go
	go build -o ./bin/client ./cmd/client.go
	go build -o ./bin/calendar_sender ./internal/calendar_sender/calendar_sender.go

test:
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/calendar
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/schedulerer

go_test:
	go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar && \
    go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage && \
    go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/grpc && \
    go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq && \
    go  test -v --count=1 github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/schedulerer