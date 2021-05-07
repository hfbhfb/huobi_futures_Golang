package main

import (
	"fmt"
	"huobi_futures_Golang/sdk/linearswap/ws"
)

func main() {
	fmt.Println("client")

	//var wsnfClient *ws.WSNotifyClient
	//wsnfClient =

	new(ws.WSNotifyClient).Init("", "", "")

}
