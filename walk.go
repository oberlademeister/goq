package goq

import "log"

func (o *Object) WalkLeaves(fn func(key string, index int, i interface{})) {
	type todoT struct {
		i     interface{}
		key   string
		index int
	}

	todo := []todoT{
		{
			i:     o.i,
			key:   "",
			index: 0,
		},
	}
	for len(todo) > 0 {
		this := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		switch v := this.i.(type) {
		case string, float64, bool:
			fn(this.key, this.index, this)
			continue
		case map[string]interface{}:
			for k, v2 := range v {
				switch v3 := v2.(type) {
				case string, float64, bool:
					fn(k, 0, v2)
					continue
				case []interface{}:
					todo = append(todo, todoT{
						i:     v2,
						key:   k,
						index: 0,
					})
					continue
				case map[string]interface{}:
					todo = append(todo, todoT{
						i:     v2,
						key:   k,
						index: 0,
					})
				default:
					log.Printf("unknown type %t", v3)
				}
			}
		case []interface{}:
			for i, v2 := range v {
				switch v3 := v2.(type) {
				case string, float64, bool:
					fn("", i, v2)
					continue
				case []interface{}:
					todo = append(todo, todoT{
						i:     v2,
						key:   "",
						index: i,
					})
					continue
				case map[string]interface{}:
					todo = append(todo, todoT{
						i:     v2,
						key:   "",
						index: i,
					})
				default:
					log.Printf("unknown type %t", v3)
				}
			}
		}
	}

}
