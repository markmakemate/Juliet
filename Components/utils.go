package Components

import "fmt"

//Delete uint64 object from []uint64 and return []uint64
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

//Print log on shell
func PrintLog(status int, err error, traffic *Traffic) {
	if err != nil {
		fmt.Printf("Client's url path:" + traffic.Request.URL.Path)
		fmt.Printf("The status is: " + string(status))
		fmt.Printf("The error info: " + err.Error())
	}
}

/**
判定Map的类型，return 整型
:return 1 为 map[uint64]DomainPtr
:return 2 为 map[uint64]LayerPtr
:return 3 为 map[uint64]ExperimentPtr
:return 4 为 map[string]ParamPtr
*/
func JudgeMapType(a interface{}) int {
	_, ok1 := a.(map[uint64]*Domain)
	_, ok2 := a.(map[uint64]*Layer)
	_, ok3 := a.(map[uint64]*Experiment)
	_, ok4 := a.(map[string]*Parameter)
	if ok1 && !ok2 && !ok3 && !ok4 {
		return 1
	} else if !ok1 && ok2 && !ok3 && !ok4 {
		return 2
	} else if ok3 && !ok1 && !ok2 && !ok4 {
		return 3
	} else if !ok3 && !ok1 && !ok2 && ok4 {
		return 4
	} else {
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return 0
	}
}

/**
判定Array的类型, return 整型
:return 1 为 []DomainPtr
:return 2 为 []LayerPtr
:return 3 为 []ExperimentPtr
:return 4 为 []ParamPtr
*/
func JudgeArrayType(a interface{}) int {
	_, ok1 := a.([]*Domain)
	_, ok2 := a.([]*Layer)
	_, ok3 := a.([]*Experiment)
	_, ok4 := a.([]*Parameter)
	if ok1 && !ok2 && !ok3 && !ok4 {
		return 1
	} else if !ok1 && ok2 && !ok3 && !ok4 {
		return 2
	} else if ok3 && !ok1 && !ok2 && !ok4 {
		return 3
	} else if !ok3 && !ok1 && !ok2 && ok4 {
		return 4
	} else {
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return 0
	}
}

//Transform abstract type to array
func transformListType(a interface{}) interface{} {
	flag := JudgeArrayType(a)
	switch flag {
	case 1:
		p, _ := a.([]*Domain)
		return p
	case 2:
		p, _ := a.([]*Layer)
		return p
	case 3:
		p, _ := a.([]*Experiment)
		return p
	case 4:
		p, _ := a.([]*Parameter)
		return p
	default:
		fmt.Println("The type of container's list is wrong. Please check the container's list type")
		return nil
	}
}

//Transform abstract type to map
func transformMapperType(a interface{}) interface{} {
	flag := JudgeMapType(a)
	switch flag {
	case 1:
		p, _ := a.(map[uint64]*Domain)
		return p
	case 2:
		p, _ := a.(map[uint64]*Layer)
		return p
	case 3:
		p, _ := a.(map[uint64]*Experiment)
		return p
	case 4:
		p, _ := a.(map[string]*Parameter)
		return p
	default:
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}

	/**

	_, ok1 := a.(map[uint64]*Domain)
	_, ok2 := a.(map[uint64]*Layer)
	_, ok3 := a.(map[uint64]*Experiment)
	_, ok4 := a.(map[string]*Parameter)
	if ok1 && !ok2 && !ok3 && !ok4 {
		p, _ := a.(map[uint64]*Domain)
		return p
	} else if !ok1 && ok2 && !ok3 && !ok4 {
		p, _ := a.(map[uint64]*Layer)
		return p
	} else if ok3 && !ok1 && !ok2 &&!ok4 {
		p, _ := a.(map[uint64]*Experiment)
		return p
	} else if !ok3 && !ok1 && !ok2 && ok4{
		p, _ := a.(map[string]*Parameter)
		return p
	} else {
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}

	*/
}

/**
Delete any object from its corresponding array and return the array has been processed
*/
func DeleteFromArray(arr interface{}, key interface{}) interface{} {
	domain, ok1 := arr.([]*Domain)
	layer, ok2 := arr.([]*Layer)
	expt, ok3 := arr.([]*Experiment)
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
