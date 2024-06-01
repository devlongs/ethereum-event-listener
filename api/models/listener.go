package models

type Listener struct {
	ContractAddress string `json:"contract_address"`
	EventName       string `json:"event_name"`
}
