package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"v2/common"
	"v2/utils"
)

const (
	baseUrl        = "https://riot:%s@127.0.0.1:%s"
	getSummoner    = "/lol-summoner/v1/current-summoner"
	getGameHistory = "/lol-match-history/v1/products/lol/%s/matches"
	getGameDetail  = "/lol-match-history/v1/games/%s"
	getFriendGroup = "/lol-chat/v1/friend-groups"
	getFriend      = "/lol-chat/v1/friend-groups/%s/friends"
)

type GameHistoryClient struct {
	BaseUrl    string
	Summoner   *common.Summoner
	Games      []common.Games
	GameIdList []int
	GameList   []*common.GameDetail
	FriendList []*common.Friend
}

func NewGameHistoryClient(port string, token string) *GameHistoryClient {
	url := fmt.Sprintf(baseUrl, token, port)
	return &GameHistoryClient{
		BaseUrl: url,
	}
}

func (g *GameHistoryClient) GetSummoner() error {
	url := g.BaseUrl + "/lol-summoner/v1/current-summoner"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	data, err := utils.Do(request)
	if err != nil {
		return err
	}
	summoner := &common.Summoner{}
	err = json.Unmarshal(data, summoner)
	if err != nil {
		return err
	}
	g.Summoner = summoner
	return nil
}

func (g *GameHistoryClient) GetFriend() error {
	url := g.BaseUrl + "/lol-chat/v1/friends"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	data, err := utils.Do(request)
	if err != nil {
		return err
	}
	friends := make([]*common.Friend, 0)
	err = json.Unmarshal(data, &friends)
	if err != nil {
		return err
	}
	g.FriendList = friends
	//PrintFriends(g.FriendList)
	return nil
}

func PrintFriends(friends []*common.Friend) {
	fmt.Println("---------------friend------------")
	for _, friend := range friends {
		fmt.Printf("1:%s 2:%s 3:%s\n", friend.Summary, friend.ProductName, friend.Name)
	}
	fmt.Println("---------------friend------------")
}

func (g *GameHistoryClient) GetGameDetail() error {
	gameDetails := make([]*common.GameDetail, 0)
	for _, id := range g.GameIdList {
		url := g.BaseUrl + fmt.Sprintf("/lol-match-history/v1/games/%s", strconv.Itoa(id))
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("NewRequest err:%s\n", err)
			continue
		}

		data, err := utils.Do(request)
		if err != nil {
			fmt.Printf(" utils.Do err:%s\n", err)
			continue
		}
		detail := &common.GameDetail{}
		err = json.Unmarshal(data, detail)
		if err != nil {
			fmt.Printf("Unmarshal err:%s\n", err)
			continue
		}
		gameDetails = append(gameDetails, detail)
	}
	g.GameList = gameDetails
	return nil
}

func (g *GameHistoryClient) GetGameHistory() error {
	gameIdList := make([]int, 0)
	for begin, end := 0, 199; end < 2000; {
		//fmt.Printf("begin=%d,end=%d\n", begin, end)
		url := g.BaseUrl + fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%s&endIndex=%s",
			g.Summoner.Puuid, strconv.Itoa(begin), strconv.Itoa(end))
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		data, err := utils.Do(request)
		if err != nil {
			fmt.Printf("GetGameHistory err:%s\n", err)
			continue
		}
		gameHistory := &common.GameHistory{}
		err = json.Unmarshal(data, gameHistory)
		if err != nil {
			continue
		}

		for _, game := range gameHistory.Game.Games {
			gameIdList = append(gameIdList, game.GameID)
		}
		begin = end
		end += 200
	}
	g.GameIdList = gameIdList
	fmt.Println("统计场次:", len(gameIdList))
	return nil
}

func (g *GameHistoryClient) Start() error {
	err := g.GetSummoner()
	if err != nil {
		fmt.Printf("GetSummoner err:%s\n", err)
		return err
	}
	err = g.GetFriend()
	if err != nil {
		fmt.Printf("GetFriend err:%s\n", err)
		return err
	}
	err = g.GetGameHistory()
	if err != nil {
		fmt.Printf("GetGameHistory err:%s\n", err)
		return err
	}
	err = g.GetGameDetail()
	if err != nil {
		fmt.Printf("GetGameDetail err:%s\n", err)
		return err
	}
	return nil
}
