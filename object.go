package goq

import (
	"bytes"
	"encoding/json"
	"io"
)

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

// NewObjectFromReader also does the unmarshaling
func NewObjectFromReader(in io.Reader) (*Object, error) {
	var i interface{}
	r := json.NewDecoder(in)
	err := r.Decode(&i)
	if err != nil {
		return nil, err
	}
	return &Object{
		i:    i,
		err:  nil,
		step: 0,
	}, nil
}

// NewObjectFromReader also does the unmarshaling
func NewObjectFromBytes(b []byte) (*Object, error) {
	var i interface{}
	in := bytes.NewReader(b)
	r := json.NewDecoder(in)
	err := r.Decode(&i)
	if err != nil {
		return nil, err
	}
	return &Object{
		i:    i,
		err:  nil,
		step: 0,
	}, nil
}

func (po *Object) Error() error {
	return po.err
}

func (po *Object) Interface() interface{} {
	return po.i
}
