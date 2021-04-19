package main

import (
	"fmt"

	"github.com/echenim/corelightfx/repository"
	"github.com/echenim/corelightfx/utilities"
)

func main() {
	cron()
}

func cron() {
	DbStoreLinker := utilities.Instanciate()
	c, _ := DbStoreLinker.DB()
	defer c.Close()

	providernasdaq := repository.ProviderNasdaqRespository(DbStoreLinker)
	nasdaq := providernasdaq.FindAll()

	for i := 0; i < len(nasdaq); i++ {
		
	}

	//cron()
}
