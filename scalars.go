package goq

import "fmt"

func (po *Object) ToString() (string, error) {
	if po.err != nil {
		return "", po.err
	}
	m, ok := po.i.(string)
	if !ok {
		return "", fmt.Errorf("ToMap(): no string")
	}
	return m, nil
}

func (po *Object) ToBool() (bool, error) {
	if po.err != nil {
		return false, po.err
	}
	m, ok := po.i.(bool)
	if !ok {
		return false, fmt.Errorf("ToMap(): no string")
	}
	return m, nil
}

func (po *Object) ToFloat64() (float64, error) {
	if po.err != nil {
		return 0.0, po.err
	}
	m, ok := po.i.(float64)
	if !ok {
		return 0.0, fmt.Errorf("ToMap(): no string")
	}
	return m, nil
}
