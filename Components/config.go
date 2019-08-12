package Components

type ConfigOfComponents struct {
	Diversor *Diversor
}

//The definition of Container Manager's configuration
type ConfigOfContainerManager struct {
	UUID            string
	DomainContainer *DomainContainer
	ExptContainer   *ExperimentContainer
	LayerContainer  *LayerContainer
}
type ConfigOfContainer struct {
	UUID            string
	ComponentList   interface{}
	ComponentMapper interface{}
	/**
	Type == 1: Domain configuration
	Type == 2: Layer configuration
	Type == 3: Experiment configuration
	*/
	Type int
}

var ConfigurationOfComponents ConfigOfComponents
var ConfigurationOfContainer ConfigOfContainer
var ConfigurationOfContainerManager ConfigOfContainerManager
