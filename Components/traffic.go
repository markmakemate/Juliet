package Components

import (
	"net/http"
)

/**
Traffic实现
*/
func (traffic *Traffic) Init(writer http.ResponseWriter, r *http.Request,
	domainId uint64, layerId uint64) {
	traffic.Request = r
	traffic.DomainId = domainId
	traffic.Writer = writer
	traffic.LayerId = layerId
}

func (traffic *Traffic) Send(signal chan int) {
	go func() {
		x := <-signal
		if x == 1 {
			status, err := traffic.Writer.Write(traffic.Expt.Param.toString())
			if err != nil {
				PrintLog(status, err, traffic)
			}
		}
	}()
}

func (traffic *Traffic) SetExpt(experiment Experiment) {
	traffic.Expt = experiment
}
