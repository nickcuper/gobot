package gpio

import (
	"strconv"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*DirectPinDriver)(nil)

// Represents a raw GPIO pin
type DirectPinDriver struct {
	name       string
	pin        string
	connection gobot.Connection
	gobot.Commander
}

// NewDirectPinDriver return a new DirectPinDriver given a DirectPin, name and pin.
//
// Adds the following API Commands:
// 	"DigitalRead" - See DirectPinDriver.DigitalRead
// 	"DigitalWrite" - See DirectPinDriver.DigitalWrite
// 	"AnalogRead" - See DirectPinDriver.AnalogRead
// 	"AnalogWrite" - See DirectPinDriver.AnalogWrite
// 	"PwmWrite" - See DirectPinDriver.PwmWrite
// 	"ServoWrite" - See DirectPinDriver.ServoWrite
func NewDirectPinDriver(a DirectPin, name string, pin string) *DirectPinDriver {
	d := &DirectPinDriver{
		name:       name,
		connection: a.(gobot.Connection),
		pin:        pin,
		Commander:  gobot.NewCommander(),
	}

	d.AddCommand("DigitalRead", func(params map[string]interface{}) interface{} {
		val, err := d.DigitalRead()
		return map[string]interface{}{"val": val, "err": err}
	})
	d.AddCommand("DigitalWrite", func(params map[string]interface{}) interface{} {
		level, _ := strconv.Atoi(params["level"].(string))
		return d.DigitalWrite(byte(level))
	})
	d.AddCommand("AnalogRead", func(params map[string]interface{}) interface{} {
		val, err := d.AnalogRead()
		return map[string]interface{}{"val": val, "err": err}
	})
	d.AddCommand("AnalogWrite", func(params map[string]interface{}) interface{} {
		level, _ := strconv.Atoi(params["level"].(string))
		return d.AnalogWrite(byte(level))
	})
	d.AddCommand("PwmWrite", func(params map[string]interface{}) interface{} {
		level, _ := strconv.Atoi(params["level"].(string))
		return d.PwmWrite(byte(level))
	})
	d.AddCommand("ServoWrite", func(params map[string]interface{}) interface{} {
		level, _ := strconv.Atoi(params["level"].(string))
		return d.ServoWrite(byte(level))
	})

	return d
}

func (d *DirectPinDriver) adaptor() DirectPin {
	return d.Connection().(DirectPin)
}

func (d *DirectPinDriver) Name() string                 { return d.name }
func (d *DirectPinDriver) Pin() string                  { return d.pin }
func (d *DirectPinDriver) Connection() gobot.Connection { return d.connection }

// Starts the DirectPinDriver. Returns true on successful start of the driver
func (d *DirectPinDriver) Start() (errs []error) { return }

// Halts the DirectPinDriver. Returns true on successful halt of the driver
func (d *DirectPinDriver) Halt() (errs []error) { return }

// DigitalRead returns the current digital state of the pin
func (d *DirectPinDriver) DigitalRead() (val int, err error) {
	return d.adaptor().DigitalRead(d.Pin())
}

// DigitalWrite writes to the pin
func (d *DirectPinDriver) DigitalWrite(level byte) (err error) {
	return d.adaptor().DigitalWrite(d.Pin(), level)
}

// AnalogRead reads the current analog reading of the pin
func (d *DirectPinDriver) AnalogRead() (val int, err error) {
	return d.adaptor().AnalogRead(d.Pin())
}

// AnalogWrite writes to the pin
func (d *DirectPinDriver) AnalogWrite(level byte) (err error) {
	return d.adaptor().AnalogWrite(d.Pin(), level)
}

// PwmWrite writes to the pin
func (d *DirectPinDriver) PwmWrite(level byte) (err error) {
	return d.adaptor().PwmWrite(d.Pin(), level)
}

// ServoWrite writes to the pin
func (d *DirectPinDriver) ServoWrite(level byte) (err error) {
	return d.adaptor().ServoWrite(d.Pin(), level)
}
