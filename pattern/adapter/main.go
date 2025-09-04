package main

import "fmt"

// The Adapter pattern is a structural design pattern that allows objects with incompatible interfaces to collaborate.
// The Adapter pattern works as a bridge between two incompatible interfaces.

type OldPhone struct{}

func (o *OldPhone) chargeWithMicroUsb() {
	fmt.Println("charge with micro usb")
}

type NewCharger struct{}

func (o *NewCharger) chargeWithUsbC() {
	fmt.Println("charge with usb c")
}

type ChargerAdapter struct {
	charger *NewCharger
}

func (c *ChargerAdapter) chargeWithMicroUsb() {
	// new charger use usb c to charge. old phone need micro usb to charge
	// charger adapter will convert usb c to micro usb and can charge for old phone

	fmt.Println("adapter convert usb c to micro usb")
	c.charger.chargeWithUsbC()
}

func main() {
	oldPhone := OldPhone{}
	newCharger := NewCharger{}
	adapter := ChargerAdapter{
		charger: &newCharger,
	}

	fmt.Println("old phone charge process")
	oldPhone.chargeWithMicroUsb()

	fmt.Println("old phone charge with charger adapter")
	adapter.chargeWithMicroUsb()
}
