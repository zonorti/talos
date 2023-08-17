// Code generated by "enumer -type=OutputKind -linecomment -text"; DO NOT EDIT.

package profile

import (
	"fmt"
)

const _OutputKindName = "unknownisoimageinstallerkernelinitramfsukicmdline"

var _OutputKindIndex = [...]uint8{0, 7, 10, 15, 24, 30, 39, 42, 49}

func (i OutputKind) String() string {
	if i < 0 || i >= OutputKind(len(_OutputKindIndex)-1) {
		return fmt.Sprintf("OutputKind(%d)", i)
	}
	return _OutputKindName[_OutputKindIndex[i]:_OutputKindIndex[i+1]]
}

var _OutputKindValues = []OutputKind{0, 1, 2, 3, 4, 5, 6, 7}

var _OutputKindNameToValueMap = map[string]OutputKind{
	_OutputKindName[0:7]:   0,
	_OutputKindName[7:10]:  1,
	_OutputKindName[10:15]: 2,
	_OutputKindName[15:24]: 3,
	_OutputKindName[24:30]: 4,
	_OutputKindName[30:39]: 5,
	_OutputKindName[39:42]: 6,
	_OutputKindName[42:49]: 7,
}

// OutputKindString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OutputKindString(s string) (OutputKind, error) {
	if val, ok := _OutputKindNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to OutputKind values", s)
}

// OutputKindValues returns all values of the enum
func OutputKindValues() []OutputKind {
	return _OutputKindValues
}

// IsAOutputKind returns "true" if the value is listed in the enum definition. "false" otherwise
func (i OutputKind) IsAOutputKind() bool {
	for _, v := range _OutputKindValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for OutputKind
func (i OutputKind) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for OutputKind
func (i *OutputKind) UnmarshalText(text []byte) error {
	var err error
	*i, err = OutputKindString(string(text))
	return err
}
