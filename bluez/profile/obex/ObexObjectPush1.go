package obex

import (
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/bluez"
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/util"
	"github.com/godbus/dbus"
)

// NewObjectPush1 create a new ObjectPush1 client
func NewObjectPush1(sessionPath string) *ObjectPush1 {
	a := new(ObjectPush1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.obex",
			Iface: "org.bluez.obex.ObjectPush1",
			Path:  sessionPath,
			Bus:   bluez.SessionBus,
		},
	)
	return a
}

// ObjectPush1 client
type ObjectPush1 struct {
	client *bluez.Client
}

// Close the connection
func (d *ObjectPush1) Close() {
	d.client.Disconnect()
}

//
// Send one local file to the remote device.
//
// The returned path represents the newly created transfer,
// which should be used to find out if the content has been
// successfully transferred or if the operation fails.
//
// The properties of this transfer are also returned along
// with the object path, to avoid a call to GetProperties.
//
// Possible errors:
//  - org.bluez.obex.Error.InvalidArguments
//  - org.bluez.obex.Error.Failed
//
func (a *ObjectPush1) SendFile(sourcefile string) (string, *ObexTransfer1Properties, error) {

	result := make(map[string]dbus.Variant)
	var sessionPath string
	err := a.client.Call("SendFile", 0, sourcefile).Store(&sessionPath, &result)

	transportProps := new(ObexTransfer1Properties)
	util.MapToStruct(transportProps, result)

	return sessionPath, transportProps, err
}
