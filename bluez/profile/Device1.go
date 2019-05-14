package profile

import (
	"gitlab.intern-enexoma.de/homeserver/libraries/go-bluetooth.git/bluez"
	"sync"

	"github.com/godbus/dbus"
)

// NewDevice1 create a new Device1 client
func NewDevice1(path string) *Device1 {
	a := new(Device1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: "org.bluez.Device1",
			Path:  path,
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(Device1Properties)
	a.GetProperties()
	return a
}

// Device1 client
type Device1 struct {
	client     *bluez.Client
	Properties *Device1Properties
}

// Device1Properties exposed properties for Device1
type Device1Properties struct {
	Lock             sync.RWMutex
	AdvertisingFlags []byte
	UUIDs            []string
	Blocked          bool
	Connected        bool
	LegacyPairing    bool
	Paired           bool
	ServicesResolved bool
	Trusted          bool
	ServiceData      map[string]dbus.Variant
	ManufacturerData map[uint16]dbus.Variant
	RSSI             int16
	TxPower          int16
	Adapter          dbus.ObjectPath
	Address          string
	AddressType      string
	Alias            string
	Icon             string
	Modalias         string
	Name             string
	Appearance       uint16
	Class            uint32
}

// Close the connection
func (d *Device1) Close() {
	d.client.Disconnect()
}

//Register for changes signalling
func (d *Device1) Register() (chan *dbus.Signal, error) {
	return d.client.Register(d.client.Config.Path, bluez.PropertiesInterface)
}

//Unregister for changes signalling
func (d *Device1) Unregister(signal chan *dbus.Signal) error {
	return d.client.Unregister(d.client.Config.Path, bluez.PropertiesInterface, signal)
}

//GetProperties load all available properties
func (d *Device1) GetProperties() (*Device1Properties, error) {
	d.Properties.Lock.Lock()
	err := d.client.GetProperties(d.Properties)
	d.Properties.Lock.Unlock()
	return d.Properties, err
}

//GetProperty get a property
func (d *Device1) GetProperty(name string) (dbus.Variant, error) {
	return d.client.GetProperty(name)
}

//SetProperty set a property
func (d *Device1) SetProperty(name string, v interface{}) error {
	return d.client.SetProperty(name, v)
}

//CancelParing stop the pairing process
func (d *Device1) CancelPairing() error {
	return d.client.Call("CancelPairing", 0).Store()
}

//Connect to the device
func (d *Device1) Connect() error {
	return d.client.Call("Connect", 0).Store()
}

//ConnectProfile connect to the specific profile
func (d *Device1) ConnectProfile(uuid string) error {
	return d.client.Call("ConnectProfile", 0, uuid).Store()
}

//Disconnect from the device
func (d *Device1) Disconnect() error {
	return d.client.Call("Disconnect", 0).Store()
}

//DisconnectProfile from the device
func (d *Device1) DisconnectProfile(uuid string) error {
	return d.client.Call("DisconnectProfile", 0, uuid).Store()
}

//Pair with the device
func (d *Device1) Pair() error {
	return d.client.Call("Pair", 0).Store()
}
