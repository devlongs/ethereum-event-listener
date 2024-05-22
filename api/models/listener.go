package models

import "gorm.io/gorm"

type Listener struct {
    gorm.Model
    ContractAddress string `json:"contract_address"`
    EventName       string `json:"event_name"`
}