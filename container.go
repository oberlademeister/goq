package goq

import "fmt"

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
