package ethereum

import (
	"context"
	"log"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/config"
	"github.com/devlongs/ethereum-event-listener/database"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenForEvents(listener models.Listener) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    client, err := ethclient.Dial(cfg.EthereumURL)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    contractAddress := common.HexToAddress(listener.ContractAddress)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            event := models.Event{
                ContractAddress: vLog.Address.Hex(),
                EventName:       listener.EventName,
                BlockNumber:     vLog.BlockNumber,
                TransactionHash: vLog.TxHash.Hex(),
                Data:            common.Bytes2Hex(vLog.Data),
            }
            if err := database.DB.Create(&event).Error; err != nil {
                log.Println(err)
            }
        }
    }
}