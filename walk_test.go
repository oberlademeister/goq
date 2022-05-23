package goq

import (
	"errors"
	"testing"

	_ "embed"

	"github.com/davecgh/go-spew/spew"
)

//go:embed testinput/t1.json
var t1 []byte

func TestObject_Walk(t *testing.T) {
	t.Log("creating object")
	o, err := NewObjectFromBytes([]byte(t1))
	if err != nil {
		t.Error(err)
	}
	var container interface{}
	err = o.WalkLeaves(false, func(key string, index int, thisI, parentI interface{}) error {
		if key != "title" {
			return nil
		}
		switch parentI.(type) {
		case map[string]interface{}:
			container = parentI
			return &WalkAbortError{}
		}
		return nil
	})
	var aerr *WalkAbortError
	if err != nil && !errors.As(err, &aerr) {
		t.Error(err)
	}
	spew.Dump(container)
}
