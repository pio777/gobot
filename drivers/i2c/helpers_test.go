package i2c

var rgb = map[string]interface{}{
	"red":   1.0,
	"green": 1.0,
	"blue":  1.0,
}

func castColor(color string) byte {
	return byte(rgb[color].(float64))
}

var red = castColor("red")
var green = castColor("green")
var blue = castColor("blue")

type i2cTestAdaptor struct {
	name         string
	written      []byte
	i2cReadImpl  func() ([]byte, error)
	i2cWriteImpl func() error
	i2cStartImpl func() error
}

func (t *i2cTestAdaptor) I2cStart(int) (err error) {
	return t.i2cStartImpl()
}
func (t *i2cTestAdaptor) I2cRead(int, int) (data []byte, err error) {
	return t.i2cReadImpl()
}
func (t *i2cTestAdaptor) I2cWrite(address int, b []byte) (err error) {
	t.written = append(t.written, b...)
	return t.i2cWriteImpl()
}
func (t *i2cTestAdaptor) Name() string          { return t.name }
func (t *i2cTestAdaptor) SetName(n string)      { t.name = n }
func (t *i2cTestAdaptor) Connect() (err error)  { return }
func (t *i2cTestAdaptor) Finalize() (err error) { return }

func newI2cTestAdaptor() *i2cTestAdaptor {
	return &i2cTestAdaptor{
		i2cReadImpl: func() ([]byte, error) {
			return []byte{}, nil
		},
		i2cWriteImpl: func() error {
			return nil
		},
		i2cStartImpl: func() error {
			return nil
		},
	}
}
