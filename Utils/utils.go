package Utils

import (
	"fmt"
	"reflect"
)

//Del uint64 object from []uint64 and return []uint64
func DelFromUint64Array(a []uint64, key uint64) []uint64 {
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

func Log(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

type TypeError struct {
	Info string
}

func (e *TypeError) Error() string {
	return e.Info
}

func GetRecover()  {
	if err := recover(); err != nil{
		fmt.Println(err)
	}
}

func Join(arr1 interface{}, arr2 interface{}) {
	if reflect.TypeOf(arr1).Kind() == reflect.Slice &&
		reflect.TypeOf(arr2).Kind() == reflect.Slice &&
		reflect.ValueOf(arr1).Index(0).Type() == reflect.ValueOf(arr2).Index(0).Type(){
		arrValue := reflect.AppendSlice(reflect.ValueOf(arr1), reflect.ValueOf(arr2))
		reflect.ValueOf(arr1).Set(arrValue)
	}else{
		panic("Two arrays' type are not equivalent or they are not array type")
	}
}

func JoinByIndex(temp reflect.Value, index int)  {
	if index == temp.Len() - 1{
		temp.Set(temp.Slice(0, index - 1))
	}else{
		a := temp.Slice(index + 1, temp.Len() - 1)
		if index > 0{
			b := temp.Slice(0, index - 1)
			temp.Set(reflect.AppendSlice(b, a))
		}else{
			temp.Set(a)
		}
	}
}

func DelFromArray(arr interface{}, key interface{}) error {
	defer GetRecover()
	temp := reflect.ValueOf(arr).Elem()
	Key := reflect.ValueOf(key).Elem()
	if reflect.TypeOf(arr).Kind() == reflect.Slice &&
		temp.Index(0).Kind() == Key.Kind(){
		index := 0
		for i := 0; i< temp.Len(); i++ {
			if temp.Index(i) == Key{
				index = i
				break
			}
		}
		JoinByIndex(temp, index)
		return nil
	} else {
		return &TypeError{
			Info:"The DelFromArray method only process array type and the key's" +
				"type should be the same! Please check the type",
		}
	}
}


func IsExist(arr interface{}, object interface{}) (int, bool) {
	flag := false
	index := 0
	ArrType := reflect.TypeOf(arr)
	ArrValue := reflect.ValueOf(arr).Elem()
	if ArrType.Kind() == reflect.Slice &&
		ArrValue.Kind() == reflect.TypeOf(object).Kind(){
		for i := 0; i < ArrValue.Len() ; i++  {
			if ArrValue.Index(i) == reflect.ValueOf(object){
				flag = true
				index = i
			}
		}
	}
	return index, flag
}

func TransformString2Uint64(key string) uint64 {
	return 0
}
