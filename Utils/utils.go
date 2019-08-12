package Utils

import (
	"caibeike-abtest/Components"
	"fmt"
)

func Log(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DeleteFromUint64Array(a []uint64, key uint64) []uint64 {
	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] == key {
			index = i
			break
		}
	}

	temp := a[index+1:]
	a = a[:index-1]
	for i := 0; i < len(temp); i++ {
		a = append(a, temp[i])
	}
	return a
}

func DeleteFromArray(arr interface{}, key interface{}) interface{} {
	domain, ok1 := arr.([]*Components.Domain)
	layer, ok2 := arr.([]*Components.Layer)
	expt, ok3 := arr.([]*Components.Experiment)
	index := 0
	if ok1 && !ok2 && !ok3 {
		for i := 0; i < len(domain); i++ {
			if domain[i] == key {
				index = i
				break
			}
		}

		temp := domain[index+1:]
		domain = domain[:index-1]
		for i := 0; i < len(temp); i++ {
			domain = append(domain, temp[i])
		}
		return domain
	} else if ok2 && !ok1 && !ok3 {
		for i := 0; i < len(layer); i++ {
			if layer[i] == key {
				index = i
				break
			}
		}

		temp := layer[index+1:]
		layer = layer[:index-1]
		for i := 0; i < len(temp); i++ {
			layer = append(layer, temp[i])
		}
		return layer
	} else if ok3 && !ok1 && ok2 {
		for i := 0; i < len(expt); i++ {
			if expt[i] == key {
				index = i
				break
			}
		}

		temp := expt[index+1:]
		expt = expt[:index-1]
		for i := 0; i < len(temp); i++ {
			expt = append(expt, temp[i])
		}
		return expt
	} else {
		fmt.Println("The type should be consistent! Please check the type")
		return nil
	}
}

func IsExist(arr []uint64, object uint64) bool {
	flag := false
	for x := range arr {
		if arr[x] == object {
			flag = true
			break
		}
	}
	return flag
}
