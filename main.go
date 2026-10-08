package main

import (
	"device/nrf"
	"log"
	"machine"
	"machine/usb"
	"time"

	keyboard "github.com/sago35/tinygo-keyboard"
	jp "github.com/sago35/tinygo-keyboard/keycodes/japanese"
)

func init() {
	// Hide from the USB host until d.Init() has configured VID/PID and HID.
	nrf.USBD.USBPULLUP.Set(0) // temporary workaround: hide USB until d.Init()
}

func main() {
	// BLE name is usb.Product + " ble"; keep usb.Product <= 13 chars to fit the advertisement.
	usb.Product = "conf2025badge"

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	d := keyboard.New()

	// Two keys and the rotary encoder push.
	gpioPins := []machine.Pin{machine.D2, machine.D3, machine.D8} // for conf2025badge
	//gpioPins := []machine.Pin{machine.P1_06, machine.P1_04, machine.P0_22} // for nicebadge
	for _, p := range gpioPins {
		p.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
	d.AddGpioKeyboard(gpioPins, [][]keyboard.Keycode{
		{
			jp.KeyA, jp.KeyB, jp.KeyBluetoothUnpair,
		},
	})

	d.Keyboard = keyboard.NewBLEKeyboard()

	err := d.Init()
	if err != nil {
		return err
	}
	nrf.USBD.USBPULLUP.Set(1) // temporary workaround: show USB after d.Init()

	// Reset if the main loop stops (known SoftDevice idle hang, see sg48key).
	machine.Watchdog.Configure(machine.WatchdogConfig{TimeoutMillis: 1000})
	machine.Watchdog.Start()

	for {
		time.Sleep(1 * time.Millisecond)
		machine.Watchdog.Update()
		err := d.Tick()
		if err != nil {
			return err
		}
	}
}
