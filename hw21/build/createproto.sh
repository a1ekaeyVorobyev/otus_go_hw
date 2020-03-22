#!/bin/bash
protoc  --proto_path=api calendar.proto --go_out=plugins=grpc:pkg/calendar