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
