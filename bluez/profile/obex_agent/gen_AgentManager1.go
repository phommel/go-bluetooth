// Code generated DO NOT EDIT

package obex_agent



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var AgentManager1Interface = "org.bluez.obex.AgentManager1"


// NewAgentManager1 create a new instance of AgentManager1
//
// Args:

func NewAgentManager1() (*AgentManager1, error) {
	a := new(AgentManager1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.obex",
			Iface: AgentManager1Interface,
			Path:  dbus.ObjectPath("/org/bluez/obex"),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(AgentManager1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
AgentManager1 Agent Manager hierarchy

*/
type AgentManager1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*AgentManager1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// AgentManager1Properties contains the exposed properties of an interface
type AgentManager1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *AgentManager1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *AgentManager1Properties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *AgentManager1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return AgentManager1 object path
func (a *AgentManager1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return AgentManager1 dbus client
func (a *AgentManager1) Client() *bluez.Client {
	return a.client
}

// Interface return AgentManager1 interface
func (a *AgentManager1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *AgentManager1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a AgentManager1Properties to map
func (a *AgentManager1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an AgentManager1Properties
func (a *AgentManager1Properties) FromMap(props map[string]interface{}) (*AgentManager1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an AgentManager1Properties
func (a *AgentManager1Properties) FromDBusMap(props map[string]dbus.Variant) (*AgentManager1Properties, error) {
	s := new(AgentManager1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *AgentManager1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *AgentManager1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *AgentManager1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *AgentManager1) GetProperties() (*AgentManager1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *AgentManager1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *AgentManager1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *AgentManager1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *AgentManager1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *AgentManager1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *AgentManager1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
RegisterAgent 
			Register an agent to request authorization of
			the user to accept/reject objects. Object push
			service needs to authorize each received object.

			Possible errors: org.bluez.obex.Error.AlreadyExists


*/
func (a *AgentManager1) RegisterAgent(agent dbus.ObjectPath) error {
	
	return a.client.Call("RegisterAgent", 0, agent).Store()
	
}

/*
UnregisterAgent 
			This unregisters the agent that has been previously
			registered. The object path parameter must match the
			same value that has been used on registration.

			Possible errors: org.bluez.obex.Error.DoesNotExist



*/
func (a *AgentManager1) UnregisterAgent(agent dbus.ObjectPath) error {
	
	return a.client.Call("UnregisterAgent", 0, agent).Store()
	
}
