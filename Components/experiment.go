package Components

/*
type AbstractExperiment interface {
	GetParameter() AbstractParameter
	GetId() uint64
	SetId(id uint64)
}
*/

type Experiment struct {
	Id    uint64
	Param Parameter
	Name  string
}
