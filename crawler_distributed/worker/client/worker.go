package client

import (
	"My_Go_Study/crawler/engine"
	"My_Go_Study/crawler_distributed/config"
	"My_Go_Study/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor)  {
	return func(request engine.Request) (engine.ParseResult, error) {
		var sReq = worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
