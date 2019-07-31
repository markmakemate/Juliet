package Components

/*
type AbstractExperiment interface {
	GetParameter() AbstractParameter
	GetId() uint64
	SetId(id uint64)
}
*/

type Experiment struct {
	id uint64
	parameter Parameter
	trafficChan chan Traffic
	name string
}

func (experiment *Experiment)GetParameter() Parameter {
	return experiment.GetParameter()
}
func (experiment *Experiment)GetId() uint64 {
	return experiment.id
}
func (experiment *Experiment)SetId(id uint64)  {
	experiment.id = id
}
func (experiment *Experiment)Send()  {
	t :=<- experiment.trafficChan
	t.Send(experiment.parameter)
}
func (experiment *Experiment)GetName()string  {
	return experiment.name
}
func (experiment *Experiment)SetName(name string)  {
	experiment.name = name
}