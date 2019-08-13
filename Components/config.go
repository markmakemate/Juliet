package Components

import "github.com/golang/protobuf/ptypes/timestamp"

type ConfigOfComponents struct {
	Diversor *Diversor
}

//The definition of Container Manager's configuration
type ConfigOfContainerManager struct {
	UUID               string
	DomainContainer    *DomainContainer
	ExptContainer      *ExperimentContainer
	LayerContainer     *LayerContainer
	ParameterContainer *ParameterContainer
}
type ConfigOfContainer struct {
	UUID            string
	ComponentList   interface{}
	ComponentMapper interface{}
	/**
	Type == 0: nil
	Type == 1: Domain configuration
	Type == 2: Layer configuration
	Type == 3: Experiment configuration
	*/
	Type int
}

type ExptConfig struct {
	Type      string
	Ratio     float32
	Param     Parameter
	Timestamp timestamp.Timestamp
}

/**
三个必要的全局配置
*/
var ConfigurationOfComponents ConfigOfComponents
var ConfigurationOfContainer ConfigOfContainer
var ConfigurationOfContainerManager ConfigOfContainerManager
