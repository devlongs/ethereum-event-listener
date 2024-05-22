package database

import (
	"fmt"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
    cfg, err := config.LoadConfig()
    if err != nil {
        return err
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    err = DB.AutoMigrate(&models.Listener{}, &models.Event{})
    if err != nil {
        return err
    }

    return nil
}