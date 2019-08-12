package Components

import "fmt"

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

func PrintLog(status int, err error, traffic *Traffic) {
	fmt.Printf("Client's url path:" + traffic.Request.URL.Path)
	fmt.Printf("The status is: " + string(status))
	fmt.Printf("The error info: " + err.Error())
}

//判定Map的类型
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

//判定Array的类型
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
	flag := Components.JudgeArrayType(a)
	switch flag {
	case 1:
		p, _ := a.([]*Components.Domain)
		return p
	case 2:
		p, _ := a.([]*Components.Layer)
		return p
	case 3:
		p, _ := a.([]*Components.Experiment)
		return p
	case 4:
		p, _ := a.([]*Components.Parameter)
		return p
	default:
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}
}

//Transform abstract type to map
func transformMapperType(a interface{}) interface{} {
	flag := Components.JudgeMapType(a)
	switch flag {
	case 1:
		p, _ := a.(map[uint64]*Components.Domain)
		return p
	case 2:
		p, _ := a.(map[uint64]*Components.Layer)
		return p
	case 3:
		p, _ := a.(map[uint64]*Components.Experiment)
		return p
	case 4:
		p, _ := a.(map[string]*Components.Parameter)
		return p
	default:
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}

	/**

	_, ok1 := a.(map[uint64]*Components.Domain)
	_, ok2 := a.(map[uint64]*Components.Layer)
	_, ok3 := a.(map[uint64]*Components.Experiment)
	_, ok4 := a.(map[string]*Components.Parameter)
	if ok1 && !ok2 && !ok3 && !ok4 {
		p, _ := a.(map[uint64]*Components.Domain)
		return p
	} else if !ok1 && ok2 && !ok3 && !ok4 {
		p, _ := a.(map[uint64]*Components.Layer)
		return p
	} else if ok3 && !ok1 && !ok2 &&!ok4 {
		p, _ := a.(map[uint64]*Components.Experiment)
		return p
	} else if !ok3 && !ok1 && !ok2 && ok4{
		p, _ := a.(map[string]*Components.Parameter)
		return p
	} else {
		fmt.Println("The type of container's mapper is wrong. Please check the container's mapper type")
		return nil
	}

	*/
}
