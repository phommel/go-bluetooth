package service

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/godbus/dbus/prop"
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/bluez"
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/bluez/profile"
)

// NewGattCharacteristic1 create a new GattCharacteristic1 client
func NewGattCharacteristic1(config *GattCharacteristic1Config, props *profile.GattCharacteristic1Properties) (*GattCharacteristic1, error) {

	propInterface, err := NewProperties(config.conn)
	if err != nil {
		return nil, err
	}

	s := &GattCharacteristic1{
		config:              config,
		properties:          props,
		PropertiesInterface: propInterface,
		descriptors:         make(map[dbus.ObjectPath]*GattDescriptor1),
	}

	err = propInterface.AddProperties(s.Interface(), props)
	if err != nil {
		return nil, err
	}

	return s, nil
}

//GattCharacteristic1Config GattCharacteristic1 configuration
type GattCharacteristic1Config struct {
	objectPath dbus.ObjectPath
	service    *GattService1
	ID         int
	conn       *dbus.Conn
}

// GattCharacteristic1 client
type GattCharacteristic1 struct {
	config              *GattCharacteristic1Config
	properties          *profile.GattCharacteristic1Properties
	PropertiesInterface *Properties
	descriptors         map[dbus.ObjectPath]*GattDescriptor1
	descIndex           int
	notifying           bool
}

//Interface return the dbus interface name
func (s *GattCharacteristic1) Interface() string {
	return bluez.GattCharacteristic1Interface
}

//Path return the object path
func (s *GattCharacteristic1) Path() dbus.ObjectPath {
	return s.config.objectPath
}

//Properties return the properties of the service
func (s *GattCharacteristic1) Properties() map[string]bluez.Properties {
	p := make(map[string]bluez.Properties)
	s.properties.Descriptors = s.GetDescriptorPaths()
	p[s.Interface()] = s.properties
	return p
}

//GetDescriptors return the characteristics of the service
func (s *GattCharacteristic1) GetDescriptors() map[dbus.ObjectPath]*GattDescriptor1 {
	return s.descriptors
}

//GetDescriptorPaths return the characteristics object paths
func (s *GattCharacteristic1) GetDescriptorPaths() []dbus.ObjectPath {
	paths := make([]dbus.ObjectPath, 0)
	for path := range s.descriptors {
		paths = append(paths, path)
	}
	return paths
}

//CreateDescriptor create a new characteristic
func (s *GattCharacteristic1) CreateDescriptor(props *profile.GattDescriptor1Properties) (*GattDescriptor1, error) {
	s.descIndex++
	path := string(s.config.objectPath) + "/desc" + strconv.Itoa(s.descIndex)
	config := &GattDescriptor1Config{
		ID:             s.descIndex,
		objectPath:     dbus.ObjectPath(path),
		conn:           s.config.conn,
		characteristic: s,
	}

	props.Characteristic = config.objectPath

	desc, err := NewGattDescriptor1(config, props)
	return desc, err
}

//AddDescriptor add a characteristic
func (s *GattCharacteristic1) AddDescriptor(desc *GattDescriptor1) error {

	s.descriptors[desc.Path()] = desc

	err := desc.Expose()
	if err != nil {
		return err
	}

	err = s.config.service.GetApp().exportTree()
	if err != nil {
		return err
	}

	om := s.config.service.GetApp().GetObjectManager()
	return om.AddObject(desc.Path(), desc.Properties())
}

//RemoveDescriptor remove a characteristic
func (s *GattCharacteristic1) RemoveDescriptor(char *GattDescriptor1) error {
	if _, ok := s.descriptors[char.Path()]; ok {
		delete(s.descriptors, char.Path())
		om := s.config.service.GetApp().GetObjectManager()
		return om.RemoveObject(char.Path())
	}
	return nil
}

//ReadValue read a value
func (s *GattCharacteristic1) ReadValue(options map[string]interface{}) []byte {
	log.Debug("Characteristic.ReadValue")
	b := make([]byte, 0)
	return b
}

//WriteValue write a value
func (s *GattCharacteristic1) WriteValue(value []byte, options map[string]interface{}) {
	log.Debug("Characteristic.WriteValue")
}

//StartNotify start notification
func (s *GattCharacteristic1) StartNotify() error {
	log.Debug("Characteristic.StartNotify")
	s.notifying = true
	return nil
}

//StopNotify stop notification
func (s *GattCharacteristic1) StopNotify() error {
	log.Debug("Characteristic.StopNotify")
	s.notifying = false
	return nil
}

//Expose the char to dbus
func (s *GattCharacteristic1) Expose() error {

	conn := s.config.conn

	err := conn.Export(s, s.Path(), s.Interface())
	if err != nil {
		return err
	}

	for iface, props := range s.Properties() {
		s.PropertiesInterface.AddProperties(iface, props)
	}

	s.PropertiesInterface.Expose(s.Path())

	node := &introspect.Node{
		Interfaces: []introspect.Interface{
			//Introspect
			introspect.IntrospectData,
			//Properties
			prop.IntrospectData,
			//GattCharacteristic1
			{
				Name:       s.Interface(),
				Methods:    introspect.Methods(s),
				Properties: s.PropertiesInterface.Introspection(s.Interface()),
			},
		},
	}

	err = conn.Export(
		introspect.NewIntrospectable(node),
		s.Path(),
		"org.freedesktop.DBus.Introspectable")
	if err != nil {
		return err
	}

	return nil
}
