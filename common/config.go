package common


/**
Global Configuration of Juliet A/B test platfo

Version 0.0.1-SNAPSHOT

Data: 2019.8

Author: Song Yi

Contact: yisong@caibeike.com

Copyright: CaiBeiKe
*/

import (
	"encoding/json"
	"time"
)


/**
必要的全局配置
*/
var(
	DomainEntity    *Domain
	Diversion *Diversor
)

type ParamType []byte
type ExpIdConfig struct {
	Name string
	Objective string
	Param     []ParamType
	Timestamp time.Time
}

func (config *ExpIdConfig) Set(data []byte) error {
	err := json.Unmarshal(data, config)
	return err
}


type Arg struct {
	paramId string
	data []byte
}


