package main

import (
	"fmt"
	"sort"
	"v2/common"
)

type ScoreCalculate struct {
	GameList   []*common.GameDetail
	FriendList []*common.Friend
	ScoreList  []*common.FriendScore
}

func NewScoreCalculate(gameList []*common.GameDetail, friendList []*common.Friend) *ScoreCalculate {
	return &ScoreCalculate{
		GameList:   gameList,
		FriendList: friendList,
		ScoreList:  make([]*common.FriendScore, 0),
	}
}

func (s *ScoreCalculate) FriendToMap() map[string]*common.Friend {
	fMap := make(map[string]*common.Friend)
	for _, friend := range s.FriendList {
		fMap[friend.Name] = friend
	}
	return fMap
}

func (s *ScoreCalculate) ParticipantsToMap(participants []common.Participants) map[int]*common.Participants {
	pMap := make(map[int]*common.Participants)
	for _, p := range participants {
		newP := &common.Participants{
			Stats: common.Stats{
				Kills:   p.Stats.Kills,
				Deaths:  p.Stats.Deaths,
				Assists: p.Stats.Assists,
				Win:     p.Stats.Win,
			},
		}
		pMap[p.ParticipantID] = newP
	}
	return pMap
}

func (s *ScoreCalculate) Calculate() {
	fMap := s.FriendToMap()
	friendScoreMap := make(map[string]*common.FriendScore)

	games := s.GameList
	for _, game := range games {
		pMap := s.ParticipantsToMap(game.Participants)
		//ss, _ := json.MarshalIndent(pMap, "", "\t")
		//fmt.Printf("pMap:%s\n", ss)
		for _, partIdent := range game.ParticipantIdentities {

			name := partIdent.Player.SummonerName
			// 判断是否是好友
			if _, ok := fMap[name]; ok {
				friendScore, ok := friendScoreMap[name]
				if !ok {
					friendScore = &common.FriendScore{FriendName: name}
					friendScoreMap[name] = friendScore
				}
				participants := pMap[partIdent.ParticipantID]
				if participants.Stats.Win {
					friendScore.GameWin++
				}
				friendScore.Kill += participants.Stats.Kills
				friendScore.Death += participants.Stats.Deaths
				friendScore.Assists += participants.Stats.Assists
				friendScore.GameCount++
			}
		}
	}

	//ss, _ := json.MarshalIndent(friendScoreMap, "", "\t")
	//fmt.Printf("%s\n", ss)

	friendScoreList := make([]*common.FriendScore, 0)
	// 计算平均值
	for _, fs := range friendScoreMap {
		//if fs.GameCount < 10 {
		//	continue
		//}
		gameCount := float32(fs.GameCount)
		gameWin := float32(fs.GameWin)
		kill := float32(fs.Kill)
		death := float32(fs.Death)
		assists := float32(fs.Assists)

		fs.WinRate = gameWin / gameCount
		fs.KillAvg = kill / gameCount
		fs.DeathAvg = death / gameCount
		fs.AssistsAvg = assists / gameCount
		friendScoreList = append(friendScoreList, fs)
	}

	// 排序
	sort.Slice(friendScoreList, func(i, j int) bool {
		return friendScoreList[i].WinRate > friendScoreList[j].WinRate
	})
	s.ScoreList = friendScoreList
}

func (s *ScoreCalculate) Output() {
	fmt.Printf("---------------------\n")
	for _, fs := range s.ScoreList {
		fmt.Printf("%s  总场次:%d  胜场:%d  胜率:%.2f  ", fs.FriendName,
			fs.GameCount, fs.GameWin, fs.WinRate)
		fmt.Printf("场均KDA:%.1f/%.1f/%.1f\t\n", fs.KillAvg, fs.DeathAvg, fs.AssistsAvg)
	}
	fmt.Printf("---------------------\n")
}
