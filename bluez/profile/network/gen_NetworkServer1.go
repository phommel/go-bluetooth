// Code generated DO NOT EDIT

package network



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var NetworkServer1Interface = "org.bluez.NetworkServer1"


// NewNetworkServer1 create a new instance of NetworkServer1
//
// Args:
// - objectPath: /org/bluez/{hci0,hci1,...}
func NewNetworkServer1(objectPath dbus.ObjectPath) (*NetworkServer1, error) {
	a := new(NetworkServer1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: NetworkServer1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(NetworkServer1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
NetworkServer1 Network server hierarchy

*/
type NetworkServer1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*NetworkServer1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// NetworkServer1Properties contains the exposed properties of an interface
type NetworkServer1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *NetworkServer1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *NetworkServer1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *NetworkServer1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return NetworkServer1 object path
func (a *NetworkServer1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return NetworkServer1 dbus client
func (a *NetworkServer1) Client() *bluez.Client {
	return a.client
}

// Interface return NetworkServer1 interface
func (a *NetworkServer1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *NetworkServer1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a NetworkServer1Properties to map
func (a *NetworkServer1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an NetworkServer1Properties
func (a *NetworkServer1Properties) FromMap(props map[string]interface{}) (*NetworkServer1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an NetworkServer1Properties
func (a *NetworkServer1Properties) FromDBusMap(props map[string]dbus.Variant) (*NetworkServer1Properties, error) {
	s := new(NetworkServer1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *NetworkServer1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *NetworkServer1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *NetworkServer1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *NetworkServer1) GetProperties() (*NetworkServer1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *NetworkServer1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *NetworkServer1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *NetworkServer1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *NetworkServer1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *NetworkServer1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *NetworkServer1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Register 
			Register server for the provided UUID. Every new
			connection to this server will be added the bridge
			interface.

			Valid UUIDs are "gn", "panu" or "nap".

			Initially no network server SDP is provided. Only
			after this method a SDP record will be available
			and the BNEP server will be ready for incoming
			connections.


*/
func (a *NetworkServer1) Register(uuid string, bridge string) error {
	
	return a.client.Call("Register", 0, uuid, bridge).Store()
	
}

/*
Unregister 
			Unregister the server for provided UUID.

			All servers will be automatically unregistered when
			the calling application terminates.

*/
func (a *NetworkServer1) Unregister(uuid string) error {
	
	return a.client.Call("Unregister", 0, uuid).Store()
	
}
