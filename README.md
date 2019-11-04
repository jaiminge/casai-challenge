# Posada (CPMS)
Golang back-end for the Posada system aka CPMS (Casai Property Management System)

## Get, Build, RUN!
First download the code
 `git clone https://github.com/hotel-casai/posada.git`

then cd into it
`cd posada`

then build the code!
`go build`

finally run the code with
`./posada`

alternatively you can also do
`go run main.go`
to build & run.

Config is specified through *dev.cfg* for dev and *prod.cfg* for production.
The config is a text protobuffer.

To rebuild all protobuffers for the Posada server, run
`protoc --proto_path=proto -I=proto --go_out=plugins=grpc:proto proto/*.proto `
# casai-challenge
