package common

type GameDetail struct {
	GameCreation          int64                   `json:"gameCreation"`
	GameCreationDate      string                  `json:"gameCreationDate"`
	GameDuration          int                     `json:"gameDuration"`
	GameID                int                     `json:"gameId"`
	GameMode              string                  `json:"gameMode"`
	GameType              string                  `json:"gameType"`
	GameVersion           string                  `json:"gameVersion"`
	MapID                 int                     `json:"mapId"`
	ParticipantIdentities []ParticipantIdentities `json:"participantIdentities"`
	Participants          []Participants          `json:"participants"`
	PlatformID            string                  `json:"platformId"`
	QueueID               int                     `json:"queueId"`
	SeasonID              int                     `json:"seasonId"`
	Teams                 []Teams                 `json:"teams"`
}

type FriendScore struct {
	FriendName string  `json:"friend_name"`
	GameCount  int     `json:"game_count"`
	GameWin    int     `json:"game_win"`
	WinRate    float32 `json:"win_rate"`

	// K/D/A
	Kill       int     `json:"kill"`
	KillAvg    float32 `json:"kill_avg"`
	Death      int     `json:"death"`
	DeathAvg   float32 `json:"death_avg"`
	Assists    int     `json:"assists"`
	AssistsAvg float32 `json:"assists_avg"`
}
