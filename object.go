package smij

import "fmt"

type Object struct {
	i    interface{}
	err  error
	step int
}

// NewObject expects the results of unmarshaling a JSON blob to an interface{}
func NewObject(i interface{}) *Object {
	return &Object{
		i:    i,
		err:  nil,
		step: 0,
	}
}

func (po *Object) Index(i int) *Object {
	ret := &Object{
		i:    po.i,
		err:  po.err,
		step: po.step + 1,
	}
	if po.err != nil {
		return ret
	}
	m, ok := po.i.([]interface{})
	if !ok {
		ret.err = fmt.Errorf("Index(): no slice at step %d", ret.step)
		return ret
	}
	if i < 0 || i >= len(m) {
		ret.err = fmt.Errorf("Index(): index (%d) out of range at step %d", i, ret.step)
		return ret
	}
	ret.i = m[i]
	return ret
}

func (po *Object) Key(key string) *Object {
	ret := &Object{
		i:    po.i,
		err:  po.err,
		step: po.step + 1,
	}
	if po.err != nil {
		return ret
	}
	m, ok := po.i.(map[string]interface{})
	if !ok {
		ret.err = fmt.Errorf("Key(): no map at step %d", ret.step)
		return ret
	}
	i, ok := m[key]
	if !ok {
		ret.err = fmt.Errorf("Key(): key (%s) not found at step %d", key, ret.step)
		return ret
	}
	ret.i = i
	return ret
}

func (po *Object) Error() error {
	return po.err
}

func (po *Object) Interface() interface{} {
	return po.i
}

func (po *Object) ToMap() (map[string]interface{}, error) {
	if po.err != nil {
		return nil, po.err
	}
	m, ok := po.i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ToMap(): no map")
	}
	return m, nil
}

func (po *Object) ToSlice() ([]interface{}, error) {
	if po.err != nil {
		return nil, po.err
	}
	m, ok := po.i.([]interface{})
	if !ok {
		return nil, fmt.Errorf("ToMap(): no slice")
	}
	return m, nil
}

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
