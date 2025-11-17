// Tektronix - TBS1062 - get settings Version 01 - TESTS OK
// Gustavo Murta - 2025_11_16
// go version go1.25.1 windows/amd64
// https://pkg.go.dev/github.com/google/gousb#hdr-A_Short_Tutorial
// https://github.com/google/gousb/blob/master/rawread/main.go
// https://github.com/google/gousb/blob/master/device_test.go
// C:\Users\jgust\go\programas\Tektronix\tek_get_settings\Version01

package main

import (
	"log"
	"os"
	"time"

	"github.com/google/gousb"
)

const (
	vid = gousb.ID(0x0699)
	pid = gousb.ID(0x03B0)
)

var settingsBuffer = make([]byte, 0) // settings buffer

////////////////////////////////////////////////////////////////////////////
// Reading data from SCOPE - URB_BULK in

func readScope(epInR *gousb.InEndpoint, dataLenght int) []byte {
	readBuffer := make([]byte, dataLenght)   // read buffer = X bytes
	readBytes, err := epInR.Read(readBuffer) // reading packets
	if err != nil {
		log.Printf("Read returned an error: %s", err)
	}
	if readBuffer == nil {
		log.Fatalf("Reading is not possible")
		//return
	}
	log.Printf("Read data HEXA (%d bytes): % X", readBytes, readBuffer)                      // print read buffer Hexadecimal
	log.Printf("Read data ASCII: (%d bytes): %s", readBytes, string(readBuffer[:readBytes])) // print read buffer ASCII
	return readBuffer[12:]                                                                   // return data without header (first 12 bytes)
}

func main() {

	cmdHeaderOnLRN := []byte{0x01, 0x01, 0xfe, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48, 0x45, 0x41, 0x44,
		0x45, 0x52, 0x20, 0x4f, 0x4e, 0x3b, 0x56, 0x45, 0x52, 0x42, 0x4f, 0x53, 0x45, 0x20, 0x4f, 0x4e,
		0x3b, 0x2a, 0x4c, 0x52, 0x4e, 0x3f, 0x00, 0x00} // SCPI command HEADER ON / VERBOSE ON / *LRN?

	cmdLRN0 := []byte{0x02, 0x02, 0xfd, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00} // SCPI command LEARN0 ?
	cmdLRN1 := []byte{0x02, 0x03, 0xfc, 0x00, 0x00, 0xfc, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00} // SCPI command LEARN1 ?

	////////////////////////////////////////////////////////////////////////////////////

	fileSettings, err := os.Create("TBS1062_settings.set") // create a bin format file

	if err != nil { // if any error
		log.Fatal(err) // cancel and print error message
	}

	defer fileSettings.Close() // close file at end

	ctx := gousb.NewContext() // Only one context should be needed for an application.
	defer ctx.Close()         // Close the context when done.

	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid) // Conecta ao aparelho VID/PID = 0699/03B0
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}

	if dev == nil {
		log.Printf("Device not found") //
	}
	defer dev.Close() // Close the device when done.

	// When autodetach is enabled gousb will automatically detach the kernel driver
	// on the interface and reattach it when releasing the interface

	err = dev.SetAutoDetach(true) // SetAutoDetach enables/disables automatic kernel driver detachment
	if err != nil {               // Se erro for diferente de nada
		log.Println("ERROR: falhou para detachar") // Mensagem de erro
		return
	}
	log.Print("Enabling autodetach")

	////////////////////////////////////////////////////////////////////////////
	log.Printf("Setting configuration 1")
	cfg, err := dev.Config(1) // Config returns a USB device set to use a particular config
	if err != nil {
		log.Fatalf("dev.Config(%d): %v", 1, err)
	}
	defer cfg.Close()

	////////////////////////////////////////////////////////////////////////////
	log.Printf("Claiming interface 0 (alt setting 0)")
	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("cfg.Interface(%d, %d): %v", 0, 0, err)
	}
	defer intf.Close()

	////////////////////////////////////////////////////////////////////////////
	epIn05, err := intf.InEndpoint(0x85) // In this interface open endpoint 0x85 for reading.
	if err != nil {
		log.Fatalf("%s.InEndpoint(0x85): %v", intf, err)
	}
	log.Printf("Found Input endpoint: %s", epIn05)

	////////////////////////////////////////////////////////////////////////////
	epOut06, err := intf.OutEndpoint(0x06) // And in the same interface open endpoint 0x06 for writing.
	if err != nil {
		log.Fatalf("%s.OutEndpoint(0x06): %v", intf, err)
	}
	log.Printf("Found Output endpoint: %s", epOut06)

	////////////////////////////////////////////////////////////////////////////
	epIn07, err := intf.InEndpoint(0x87) // In this interface open endpoint 0x87 for reading.
	if err != nil {
		log.Fatalf("%s.InEndpoint(0x87): %v", intf, err)
	}
	log.Printf("Found Input endpoint: %s", epIn07)

	////////////////////////////////////////////////////////////////////////////
	indexCfg := 1
	descCfg, err := dev.ConfigDescription(indexCfg) // ConfigDescription returns the description of the selected device
	if err != nil {
		log.Fatalf("Error on ConfigDescription(%s): %v", descCfg, err)
	}
	log.Printf("Device configuration [%d] %s...", indexCfg, descCfg)

	////////////////////////////////////////////////////////////////////////////
	manufacturer, err := dev.Manufacturer() // Identifica o nome do Fabricante do Osciloscopio
	if err != nil {
		log.Fatalf("dev.Manufacturer(%s): %v", manufacturer, err)
	}
	log.Printf("Manufacturer is %s...", manufacturer)

	////////////////////////////////////////////////////////////////////////////
	productName, err := dev.Product() // Identifica o nome do Osciloscopio
	if err != nil {
		log.Fatalf("dev.Product(%s): %v", productName, err)
	}
	log.Printf("Product name is %s...", productName)

	////////////////////////////////////////////////////////////////////////////
	serialNumber, err := dev.SerialNumber() // identifica o serial number do Osciloscopio
	if err != nil {
		log.Fatalf("dev.SerialNumber(%s): %v", serialNumber, err)
	}
	log.Printf("Oscilloscope serial number is %s...", serialNumber)

	//////////////////////////////////////////////////////////////////////////// Teste OK !!!!
	// URB_CONTROL in - Endpoint: 0x80, Direction: IN
	// Control (bmRequestType, bRequest, wValue, wIndex, commandBuffer) (error)
	// bmRequestType= 0xA1 -  Data direction=Device to host, Type: Class (0x1), Recipient: Interface (0x01)
	// bRequest= 0x07 (Class request?)
	// wValue= 0x0000  wIndex='0' (0x0000) Interface='0' (0x00)
	// wLength= 0x18 (24)
	commandBuffer := make([]byte, 24)                                        // Command buffer with 24 bytes
	controlCommand, err := dev.Control(0xA1, 0x07, 0x0000, 0, commandBuffer) // Function Control sends a control request to the device.
	if err != nil {
		log.Fatalf(" dev.Control %v : %v", controlCommand, err)
	}
	log.Printf("URB Control response: % X", commandBuffer)

	////////////////////////////////////////////////////////////////////////////
	// Send Command SCPI: HEADER ON / VERBOSE ON / *LRN? - URB_BULK out
	writeBytesHeaderOnLRN, err := epOut06.Write(cmdHeaderOnLRN)
	if err != nil {
		log.Fatalf("Write returned an error: %v %v", writeBytesHeaderOnLRN, err)
	}
	log.Printf("Send command HEADER ON / VERBOSE ON / *LRN?") // print command

	////////////////////////////////////////////////////////////////////////////
	// Send Command SCPI: LEARN0 - URB_BULK out
	println("")
	writeBytesLRN0, err := epOut06.Write(cmdLRN0)
	if err != nil {
		log.Fatalf("Write returned an error: %v %v", writeBytesLRN0, err)
	}
	log.Printf("Send command LEARN0: % X", cmdLRN0) // print command

	////////////////////////////////////////////////////////////////////////////
	// Reading data from SCOPE - URB_BULK in
	settingsBuffer = append(settingsBuffer, readScope(epIn05, 1036)...)
	time.Sleep(50 * time.Millisecond)

	////////////////////////////////////////////////////////////////////////////
	// Send Command SCPI: LEARN1 - URB_BULK out
	println("")
	writeBytesLRN1, err := epOut06.Write(cmdLRN1)
	if err != nil {
		log.Fatalf("Write returned an error: %v %v", writeBytesLRN1, err)
	}
	log.Printf("Send command LEARN1: % X", cmdLRN1) // print command

	////////////////////////////////////////////////////////////////////////////
	// Reading data from SCOPE - URB_BULK in
	settingsBuffer = append(settingsBuffer, readScope(epIn05, 904)...)
	time.Sleep(50 * time.Millisecond)
	println("")

	////////////////////////////////////////////////////////////////////////////
	log.Printf("Save TBS1062 settings to file")        // print
	_, err = fileSettings.Write(settingsBuffer[:1915]) // write settings buffer to file (1915 bytes)
	// ou ?
	// err = os.WriteFile("TBS1062_settings.set", settingsBuffer, 0644)
	if err != nil {
		log.Fatal(err) // cancel and print error message
	}
	log.Printf("Save file OK! ")
}
