module kcpgo

go 1.16

require (
	github.com/xtaci/kcp-go v4.3.4+incompatible
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
)

replace (
	github.com/xtaci/kcp-go v4.3.4+incompatible => ../../base/kcp-go
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 => ../../base/crypto
)
