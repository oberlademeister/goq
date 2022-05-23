package goq

import "log"

type WalkAbortError struct{}

func (e *WalkAbortError) Error() string {
	return "WalkAbortError"
}

func (o *Object) WalkLeaves(bfs bool, fn func(key string, index int, thisI, parentI interface{}) error) error {
	type todoT struct {
		i       interface{}
		parentI interface{}
		key     string
		index   int
	}

	todo := []todoT{
		{
			i:       o.i,
			parentI: nil,
			key:     "",
			index:   0,
		},
	}
	for len(todo) > 0 {
		var this todoT
		// depth first search
		if bfs {
			this = todo[0]
			todo = todo[1:]
		} else {
			this = todo[len(todo)-1]
			todo = todo[:len(todo)-1]
		}
		switch v := this.i.(type) {
		case string, float64, bool:
			err := fn(this.key, this.index, this.i, this.parentI)
			if err != nil {
				return err
			}
			continue
		case map[string]interface{}:
			for k, v2 := range v {
				switch v3 := v2.(type) {
				case string, float64, bool:
					todo = append(todo, todoT{
						i:       v2,
						parentI: this.i,
						key:     k,
						index:   0,
					})
					continue
				case []interface{}:
					todo = append(todo, todoT{
						i:       v2,
						parentI: this.i,
						key:     k,
						index:   0,
					})
					continue
				case map[string]interface{}:
					todo = append(todo, todoT{
						i:       v2,
						parentI: this.i,
						key:     k,
						index:   0,
					})
				default:
					log.Printf("unknown type %t", v3)
				}
			}
		case []interface{}:
			for i, v2 := range v {
				switch v3 := v2.(type) {
				case string, float64, bool:
					todo = append(todo, todoT{
						i:       v2,
						parentI: this.i,
						key:     "",
						index:   i,
					})
					continue
				case []interface{}:
					todo = append(todo, todoT{
						i:       v2,
						parentI: this.i,
						key:     "",
						index:   i,
					})
					continue
				case map[string]interface{}:
					todo = append(todo, todoT{
						i:       v2,
						key:     "",
						parentI: this.i,
						index:   i,
					})
				default:
					log.Printf("unknown type %t", v3)
				}
			}
		}
	}
	return nil
}
