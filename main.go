package main

import (
	"fmt"
	"time"
	"v2/utils"
)

func main() {
	info, err := utils.FindLCUConnectInfo()
	if err != nil {
		fmt.Printf("获取lol进程信息失败:%s\n", err)
		time.Sleep(time.Minute * 3)
		return
	}

	client := NewGameHistoryClient(info.Port, info.AuthToken)
	err = client.Start()
	if err != nil {
		fmt.Printf("client.Start err:%s\n", err)
		time.Sleep(time.Minute * 3)
		return
	}
	calculate := NewScoreCalculate(client.Summoner.DisplayName, client.GameList, client.FriendList)
	calculate.Calculate()
	calculate.Output()

	time.Sleep(time.Minute * 3)
}
