package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/echenim/corelightfx/repository"
	"github.com/echenim/corelightfx/utilities"
	"gorm.io/gorm"
)

func main() {
	config := GetConfig()
	intervals, _ := strconv.Atoi(config["intervalInMinute"])

	fmt.Println("Timer inter of ", intervals, " minute has been set")

	DbStoreLinker := utilities.Instanciate()
	c, _ := DbStoreLinker.DB()
	defer c.Close()
	cron(intervals, DbStoreLinker)
}

func cron(intervalsInMinute int, DbStoreLinker *gorm.DB) {

	providernasdaq := repository.ProviderNasdaqRespository(DbStoreLinker)
	providerStock := repository.ProviderRespository(DbStoreLinker)

	nasdaq := providernasdaq.FindAll()

	for i := 0; i < len(nasdaq); i++ {
		providerStock.GetStockDataFromProvider(nasdaq[i].Symbol, "Tpk_5367671f3afb40a78a02a8925e461a14")
	}
	time.Sleep(time.Duration(intervalsInMinute) * time.Minute)
	cron(intervalsInMinute, DbStoreLinker)
}

func GetConfig() map[string]string {
	j, _ := os.Open("config.json")
	defer j.Close()
	b, _ := io.ReadAll(j)
	var settings map[string]string
	json.Unmarshal([]byte(b), &settings)
	return settings
}
