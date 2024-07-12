package lego

// Commands come from <https://lego.github.io/lego-ble-wireless-protocol-docs/index.html>
const (

// TODO commands
)

/*
Manufacturer Data
Name	Description	Offset	Format	Min	Max
Length	Length of data	0	UINT8	0x09	0x09
Data Type Name	Manufacturer specific data to follow	1	UINT8	0xFF	0xFF
Manufacturer ID	Manufacturer identifier	2	UINT16	0x0397	0x0397
Button State	Reflects the state of the device button	4	UINT8	0	1
System Type and Device Number	Identifies the system type of the Hub	5	UINT8	0x00	0xFF
Device Capabilities	Describes the capabilities of the device	6	UINT8	0x01	0xFF
Last network	Network ID of the last network	7	UINT8	0	255
Status	Actual status	8	UINT8	0	255
Option	Future use	9	UINT8	0	255
*/
type Manufacturer struct {
	Length         uint8  // Length
	DataTypeName   uint8  // Data type
	ManufacturerID uint16 // ManufacturerID
	ButtonState    uint8  // Button State
	SystemType     uint8  // System Type and device number
	Capabilities   uint8  // Device Capabilities
	LastNetwork    uint8  // Last network
	Status         uint8  // Actual status
	Option         uint8  // Future use
}

// System Type <https://lego.github.io/lego-ble-wireless-protocol-docs/index.html#system-type-and-device-number>
const (
	SystemTypeSystemMask = 0b11100000
	SystemTypeDeviceMask = 0b00011111
)

const (
	SystemType_LegoWedo2   uint8 = 1
	SystemType_LegoDuplo   uint8 = 2
	SystemType_LegoSystem  uint8 = 3
	SystemType_LegoSystem2 uint8 = 4
)

/*
The Device Number in the rightmost bits represents the unique identifier for a product within a System Type. The currently available LEGO BLE compatible devices are:

SSS	DDDDD	Description
000	00000	WeDo Hub
001	00000	Duplo Train
010	00000	Boost Hub
010	00001	2 Port Hub
010	00010	2 Port Handset
*/

// Device Capabilities
/*
A LEGO BLE compatible device has capabilities, but not all devices have the same capabilities.

CCCC CCCC	Description
0000 0001	Supports Central Role
0000 0010	Supports Peripheral Role
0000 0100	Supports LPF2 devices (H/W connectors)
0000 1000	Act as a Remote Controller (R/C)
0001 0000	TBD
0010 0000	TBD
0100 0000	TBD
1000 0000	TBD
*/

const (
	DeviceCapability_Central    = 1 // Supports Central Role
	DeviceCapability_Preipheral = 2 // Supports Peripheral Role
	DeviceCapability_LPF2       = 4 // Supports LPF2 devices (H/W connectors)
	DeviceCapability_Remot      = 8 // Act as a Remote Controller (R/C)
)

// Last Network ID <https://lego.github.io/lego-ble-wireless-protocol-docs/index.html#last-network-id>
/*
The Last Network ID is used in the LEGO BLE device pairing process.

000:	NONE (unknown)
001..250:	The ID’s used for “Last Connection” Network ID’s. Used in H/W network “auto connect”
251:	DEFAULT 1, Locked
252:	DEFAULT 2, NOT Locked
253:	DEFAULT 3, RSSI Dependent
254:	DEFAULT 4, DISABLE H/W Network
255:	DON’T CARE - NOT Implemented
*/

// Status <https://lego.github.io/lego-ble-wireless-protocol-docs/index.html#status>
/*
0000 0001	“I can be Peripheral”
0000 0010	“I can be Central”
0000 0100	TBD
0000 1000	TBD
0001 0000	TBD
0010 0000	“Request Window” A stretching of the Button Pressed (Adding 1 sec. after release) [part of connection process]. See Request Window
0100 0000	“Request Connect”. Hardcoded request (i.e. CONSTANT flag)
1000 0000	TBD
*/

// Message Types <https://lego.github.io/lego-ble-wireless-protocol-docs/index.html#message-types>

/*
Name	Value	Comm.	Reply to	Notes
Hub Properties	0x01	Down + Up	0x01	Set or retrieve standard Hub Property information
Hub Actions	0x02	Down + Up	0x02	Perform actions on connected hub
Hub Alerts	0x03	Down + Up	0x03	Subscribe or retrieve Hub alerts
Hub Attached I/O	0x04	Up	N/A	Transmitted upon Hub detection of attached I/O
Generic Error Messages	0x05	Up	N/A	Generic Error Messages from the Hub
H/W NetWork Commands	0x08	Down + Up	0x08	Commands used for H/W Networks
F/W Update - Go Into Boot Mode	0x10	Down	N/A	Set the Hub in a special Boot Loader mode
F/W Update Lock memory	0x11	Down	N/A	Locks the memory
F/W Update Lock Status Request	0x12	Down	N/A	Request the Memory Locking State
F/W Lock Status	0x13	Up	0x12	Answer to the F/W Lock Status Request
*/
