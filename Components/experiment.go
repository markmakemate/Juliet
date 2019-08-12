package Components

/**
Experiment实现
*/
func (experiment *Experiment) Serialize() {

}

//RPC method
func (experiment *Experiment) Init(name string, exptId uint64) {

}

func (experiment *Experiment) SetExptConfig(config ExptConfig) {
	experiment.Type = config.Type
	experiment.Ratio = config.Ratio
	experiment.Timestamp = config.Timestamp
}
