module github.com/hslatman/mud.yang.go

go 1.16

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/openconfig/gnmi v0.0.0-20200617225440-d2b4e6a45802 // indirect
	github.com/openconfig/goyang v0.3.1
	github.com/openconfig/ygot v0.12.5
	github.com/spf13/cobra v1.2.1
)

replace github.com/openconfig/ygot => ./../ygot
