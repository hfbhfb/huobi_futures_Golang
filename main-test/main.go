package main

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"os"
)

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	AccountId int64
	SubUid    int64
}

var config *Config

func init() {

	filePtr, err := os.Open("main-test/config.json")
	if err != nil {
		// 以下是以命令行方式gowatch
		filePtr, err = os.Open("config.json")
		if err != nil {
			return
		}
		//return
	}
	defer filePtr.Close()

	config = new(Config)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(config)

	fmt.Println(config)

	acClient = restful.AccountClient{}
	acClient.Init(config.AccessKey, config.SecretKey, config.Host)
}

var acClient restful.AccountClient

func main() {

	data := make(chan account.GetBalanceValuationResponse)

	go acClient.GetBalanceValuationAsync(data, "USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		fmt.Printf("%d:%s", x.ErrorCode, x.ErrorMessage)
	} else {
		byt, _ := json.Marshal(x)
		fmt.Printf("%v", string(byt))
	}

}
