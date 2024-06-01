package models

type Event struct {
	ContractAddress string `json:"contract_address"`
	EventName       string `json:"event_name"`
	BlockNumber     uint64 `json:"block_number"`
	TransactionHash string `json:"transaction_hash"`
	Data            string `json:"data"`
}
