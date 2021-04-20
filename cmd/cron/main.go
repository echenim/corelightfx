package main

import (
	"fmt"
	"time"

	"github.com/echenim/corelightfx/repository"
	"github.com/echenim/corelightfx/utilities"
)

func main() {
	var intervalsInMinute int
	fmt.Print("Enter how often you want to retrieve stock data from provider in munites : ")
	fmt.Scanf("%d", &intervalsInMinute)

	fmt.Print("Servic has started running ..... : ")
	cron(intervalsInMinute)
}

func cron(intervalsInMinute int) {
	DbStoreLinker := utilities.Instanciate()
	c, _ := DbStoreLinker.DB()
	defer c.Close()

	providernasdaq := repository.ProviderNasdaqRespository(DbStoreLinker)
	providerStock := repository.ProviderRespository(DbStoreLinker)

	nasdaq := providernasdaq.FindAll()

	for i := 0; i < len(nasdaq); i++ {
		providerStock.GetStockDataFromProvider(nasdaq[i].Symbol, "Tpk_5367671f3afb40a78a02a8925e461a14")
	}
	time.Sleep(time.Duration(intervalsInMinute) * time.Minute)
	cron(intervalsInMinute)
}
