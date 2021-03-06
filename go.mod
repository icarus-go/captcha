module github.com/icarus-go/captcha

go 1.16

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc v1.40.1 => google.golang.org/grpc v1.26.0

require (
	github.com/icarus-go/component v0.0.73
	github.com/mojocn/base64Captcha v1.3.5
)
