package model

import "time"

type Gateway struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Version string    `json:"version"` //用于版本同步？
	Created time.Time `json:"created" xorm:"created"`
}

type GatewayEx struct {
	Gateway
	Status map[string]interface{} `json:"status"`
}
