package common

import (
	"Juliet/Utils"
	"encoding/json"
)

type Recommend struct {
	Type string      `json:"type"`
	Link string      `json:"link"`
	Id string        `json:"id"`
}

func (re *Recommend) String() []byte {
	return []byte(re.Link)
}

func (re *Recommend) Init(param []byte, result *string) error {
	if err := json.Unmarshal(param, re); err != nil{
		*result = "initialization failed"
		return err
	}
	return nil
}

func (re *Recommend) GetType() string {
	return re.Type
}

func (re *Recommend) Receive(data []byte) {
	Utils.Log(json.Unmarshal(data, re))
}

func (re *Recommend) Reset() {
	re.Type = ""
	re.Link = ""
}

func (re *Recommend) GetId() string  {
	return re.Id
}