package main

import (
	"github.com/hangyan/ethtool"
	"log"
)

func main() {
	client, err := ethtool.New()
	if err != nil{
		log.Fatal(err)
	}

	result, err := client.DisableVlanCsumOffloading("ens160")
	if err != nil {
		log.Fatal(err)

	}

	log.Println(result)
}