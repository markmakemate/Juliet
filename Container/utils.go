package Container

import (
	"caibeike-abtest/Components"
	"fmt"
)

func transformListType(a interface{}) interface{} {
	_, ok1 := a.([]*Components.Domain)
	_, ok2 := a.([]*Components.Layer)
	_, ok3 := a.([]*Components.Experiment)
	if ok1 && !ok2 && !ok3 {
		p, _ := a.([]*Components.Domain)
		return p
	} else if !ok1 && ok2 && !ok3 {
		p, _ := a.([]*Components.Layer)
		return p
	} else if ok3 && !ok1 && !ok2 {
		p, _ := a.([]*Components.Experiment)
		return p
	} else {
		fmt.Println("The type of container's list is wrong. Please check the container's list type")
		return nil
	}
}

func transformMapperType(a interface{}) interface{} {
	_, ok1 := a.(map[uint64]*Components.Domain)
	_, ok2 := a.(map[uint64]*Components.Layer)
	_, ok3 := a.(map[uint64]*Components.Experiment)
	if ok1 && !ok2 && !ok3 {
		p, _ := a.(map[uint64]*Components.Domain)
		return p
	} else if !ok1 && ok2 && !ok3 {
		p, _ := a.(map[uint64]*Components.Layer)
		return p
	} else if ok3 && !ok1 && !ok2 {
		p, _ := a.(map[uint64]*Components.Experiment)
		return p
	} else {
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}
}
