```

To identify all the information about the scope's USB port, use this program - usbview.exe

https://learn.microsoft.com/en-us/windows-hardware/drivers/debugger/usbview

TBS1062 USB View 2025_11_03.txt

Device Descriptor:
bcdUSB:             0x0200
bDeviceClass:         0x00
bDeviceSubClass:      0x00
bDeviceProtocol:      0x00
bMaxPacketSize0:      0x40 (64)
idVendor:           0x0699 (Tektronix, Inc.)
idProduct:          0x03B0
bcdDevice:          0x0042
iManufacturer:        0x01
0x0409: "Tektronix, Inc."
iProduct:             0x02
0x0409: "Tektronix TBS1062"
iSerialNumber:        0x03
0x0409: "C030255"
bNumConfigurations:   0x01

ConnectionStatus: DeviceConnected
Current Config Value: 0x01
Device Bus Speed:     Full
Device Address:       0x03
Open Pipes:              3

Endpoint Descriptor:
bEndpointAddress:     0x85  IN
Transfer Type:        Bulk
wMaxPacketSize:     0x0040 (64)
bInterval:            0x00

Endpoint Descriptor:
bEndpointAddress:     0x06  OUT
Transfer Type:        Bulk
wMaxPacketSize:     0x0040 (64)
bInterval:            0x00

Endpoint Descriptor:
bEndpointAddress:     0x87  IN
Transfer Type:   Interrupt
wMaxPacketSize:     0x0040 (64)
bInterval:            0x08

Configuration Descriptor:
wTotalLength:       0x0027
bNumInterfaces:       0x01
bConfigurationValue:  0x01
iConfiguration:       0x00
bmAttributes:         0xC0 (Bus Powered Self Powered )
MaxPower:             0x32 (100 Ma)

Interface Descriptor:
bInterfaceNumber:     0x00
bAlternateSetting:    0x00
bNumEndpoints:        0x03
bInterfaceClass:      0xFE
bInterfaceSubClass:   0x03
bInterfaceProtocol:   0x01
iInterface:           0x00

Endpoint Descriptor:
bEndpointAddress:     0x85  IN
Transfer Type:        Bulk
wMaxPacketSize:     0x0040 (64)
bInterval:            0x00

Endpoint Descriptor:
bEndpointAddress:     0x06  OUT
Transfer Type:        Bulk
wMaxPacketSize:     0x0040 (64)
bInterval:            0x00

Endpoint Descriptor:
bEndpointAddress:     0x87  IN
Transfer Type:   Interrupt
wMaxPacketSize:     0x0040 (64)
bInterval:            0x08


```
