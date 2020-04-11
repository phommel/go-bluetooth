// Code generated DO NOT EDIT

package media



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var MediaFolder1Interface = "org.bluez.MediaFolder1"


// NewMediaFolder1 create a new instance of MediaFolder1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewMediaFolder1(servicePath string, objectPath dbus.ObjectPath) (*MediaFolder1, error) {
	a := new(MediaFolder1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: MediaFolder1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaFolder1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}

// NewMediaFolder1Controller create a new instance of MediaFolder1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX/playerX
func NewMediaFolder1Controller(objectPath dbus.ObjectPath) (*MediaFolder1, error) {
	a := new(MediaFolder1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaFolder1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaFolder1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
MediaFolder1 MediaFolder1 hierarchy

*/
type MediaFolder1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*MediaFolder1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// MediaFolder1Properties contains the exposed properties of an interface
type MediaFolder1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Attributes Item properties that should be included in the list.

			Possible Values:

				"title", "artist", "album", "genre",
				"number-of-tracks", "number", "duration"

			Default Value: All
	*/
	Attributes []string

	/*
	NumberOfItems Number of items in the folder
	*/
	NumberOfItems uint32

	/*
	Name Folder name:

			Possible values:
				"/Filesystem/...": Filesystem scope
				"/NowPlaying/...": NowPlaying scope

			Note: /NowPlaying folder might not be listed if player
			is stopped, folders created by Search are virtual so
			once another Search is perform or the folder is
			changed using ChangeFolder it will no longer be listed.

Filters
	*/
	Name string

	/*
	Start Offset of the first item.

			Default value: 0
	*/
	Start uint32

	/*
	End Offset of the last item.

			Default value: NumbeOfItems
	*/
	End uint32

}

//Lock access to properties
func (p *MediaFolder1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *MediaFolder1Properties) Unlock() {
	p.lock.Unlock()
}




// SetAttributes set Attributes value
func (a *MediaFolder1) SetAttributes(v []string) error {
	return a.SetProperty("Attributes", v)
}



// GetAttributes get Attributes value
func (a *MediaFolder1) GetAttributes() ([]string, error) {
	v, err := a.GetProperty("Attributes")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}






// GetNumberOfItems get NumberOfItems value
func (a *MediaFolder1) GetNumberOfItems() (uint32, error) {
	v, err := a.GetProperty("NumberOfItems")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}






// GetName get Name value
func (a *MediaFolder1) GetName() (string, error) {
	v, err := a.GetProperty("Name")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}




// SetStart set Start value
func (a *MediaFolder1) SetStart(v uint32) error {
	return a.SetProperty("Start", v)
}



// GetStart get Start value
func (a *MediaFolder1) GetStart() (uint32, error) {
	v, err := a.GetProperty("Start")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetEnd set End value
func (a *MediaFolder1) SetEnd(v uint32) error {
	return a.SetProperty("End", v)
}



// GetEnd get End value
func (a *MediaFolder1) GetEnd() (uint32, error) {
	v, err := a.GetProperty("End")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}



// Close the connection
func (a *MediaFolder1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return MediaFolder1 object path
func (a *MediaFolder1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return MediaFolder1 dbus client
func (a *MediaFolder1) Client() *bluez.Client {
	return a.client
}

// Interface return MediaFolder1 interface
func (a *MediaFolder1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *MediaFolder1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a MediaFolder1Properties to map
func (a *MediaFolder1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an MediaFolder1Properties
func (a *MediaFolder1Properties) FromMap(props map[string]interface{}) (*MediaFolder1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an MediaFolder1Properties
func (a *MediaFolder1Properties) FromDBusMap(props map[string]dbus.Variant) (*MediaFolder1Properties, error) {
	s := new(MediaFolder1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *MediaFolder1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *MediaFolder1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *MediaFolder1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *MediaFolder1) GetProperties() (*MediaFolder1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *MediaFolder1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *MediaFolder1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *MediaFolder1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *MediaFolder1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *MediaFolder1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *MediaFolder1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Search 
			Return a folder object containing the search result.

			To list the items found use the folder object returned
			and pass to ChangeFolder.

			Possible Errors: org.bluez.Error.NotSupported
					 org.bluez.Error.Failed


*/
func (a *MediaFolder1) Search(value string, filter map[string]interface{}) (dbus.ObjectPath, error) {
	
	var val0 dbus.ObjectPath
	err := a.client.Call("Search", 0, value, filter).Store(&val0)
	return val0, err	
}

/*
ListItems 
			Return a list of items found

			Possible Errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.NotSupported
					 org.bluez.Error.Failed


*/
func (a *MediaFolder1) ListItems(filter map[string]interface{}) ([]dbus.ObjectPath, string, error) {
	
	var val0 []dbus.ObjectPath
  var val1 string
	err := a.client.Call("ListItems", 0, filter).Store(&val0, &val1)
	return val0, val1, err	
}

/*
ChangeFolder 
			Change current folder.

			Note: By changing folder the items of previous folder
			might be destroyed and have to be listed again, the
			exception is NowPlaying folder which should be always
			present while the player is active.

			Possible Errors: org.bluez.Error.InvalidArguments
					 org.bluez.Error.NotSupported
					 org.bluez.Error.Failed


*/
func (a *MediaFolder1) ChangeFolder(folder dbus.ObjectPath) error {
	
	return a.client.Call("ChangeFolder", 0, folder).Store()
	
}
