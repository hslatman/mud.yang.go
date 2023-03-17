module github.com/hslatman/mud.yang.go

go 1.19

require (
	github.com/openconfig/goyang v1.2.0
	github.com/openconfig/ygot v0.26.1
	github.com/spf13/cobra v1.6.1
)

require (
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/openconfig/gnmi v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

//replace github.com/openconfig/ygot => ./../ygot
