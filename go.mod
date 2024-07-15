module go.mau.fi/whatsmeow

go 1.21

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.0
	github.com/rs/zerolog v1.33.0
	go.mau.fi/libsignal v0.1.1-0.20240705162345-47e713a595ab
	go.mau.fi/util v0.5.0
	golang.org/x/crypto v0.25.0
	golang.org/x/net v0.27.0
	google.golang.org/protobuf v1.34.2
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/sys v0.22.0 // indirect
)

replace go.mau.fi/libsignal => github.com/Valdenirmezadri/libsignal-protocol-go v0.1.1
