package Components

import (
	"fmt"
	"log"
	"net/http"
)

type Traffic struct {
	Writer   http.ResponseWriter
	Request  http.Request
	DomainId uint64
	LayerId  uint64
	Expt     Experiment
}

func (traffic *Traffic) Send() {
	status, err := traffic.Writer.Write(traffic.Expt.Param.toString())
	fmt.Printf("Client's url path:" + traffic.Request.URL.Path)
	fmt.Printf("The status is: " + string(status))
	if err != nil {
		log.Fatal(err)
	}
}
