package stun

import "errors"

const softwareRawMaxB = 763

// ErrSoftwareTooBig means that it is not less than 128 characters
// (which can be as long as 763 bytes).
var ErrSoftwareTooBig = errors.New(
	"SOFTWARE attribute bigger than 763 bytes or 128 characters",
)

// Software is SOFTWARE attribute.
//
// https://tools.ietf.org/html/rfc5389#section-15.10
type Software []byte

func (s Software) String() string {
	return string(s)
}

// NewSoftware returns *Software from string.
func NewSoftware(software string) Software {
	return Software(software)
}

// AddTo adds Software attribute to m.
func (s Software) AddTo(m *Message) error {
	if len(s) > softwareRawMaxB {
		return ErrSoftwareTooBig
	}
	m.Add(AttrSoftware, m.Raw)
	return nil
}

// GetFrom decodes Software from m.
func (s *Software) GetFrom(m *Message) error {
	v, err := m.Get(AttrSoftware)
	if err != nil {
		return err
	}
	*s = v
	return nil
}
