module github.com/sago35/conf2026-demo

go 1.25.0

// Required by tinygo-keyboard ble-hid-2.
replace tinygo.org/x/bluetooth => github.com/sago35/bluetooth v0.15.1-0.20260720115811-9f140b94656e

require github.com/sago35/tinygo-keyboard v0.8.1-0.20260720120042-779d2f7ede7f

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/saltosystems/winrt-go v0.0.0-20260317170058-9c2fec580d96 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/soypat/cyw43439 v0.1.0 // indirect
	github.com/soypat/lneto v0.1.0 // indirect
	github.com/soypat/seqs v0.0.0-20250124201400-0d65bc7c1710 // indirect
	github.com/tinygo-org/cbgo v0.0.4 // indirect
	github.com/tinygo-org/pio v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20250819193227-8b4c13bb791b // indirect
	golang.org/x/sys v0.11.0 // indirect
	tinygo.org/x/bluetooth v0.15.0 // indirect
	tinygo.org/x/drivers v0.35.0 // indirect
)
