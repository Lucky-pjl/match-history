package common

type Summoner struct {
	AccountID                   int64        `json:"accountId"`
	DisplayName                 string       `json:"displayName"`
	GameName                    string       `json:"gameName"`
	InternalName                string       `json:"internalName"`
	NameChangeFlag              bool         `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int          `json:"percentCompleteForNextLevel"`
	Pr                          string       `json:"pr"`
	ProfileIconID               int          `json:"profileIconId"`
	Puuid                       string       `json:"puuid"`
	RerollPoints                RerollPoints `json:"rerollPoints"`
	SummonerID                  int64        `json:"summonerId"`
	SummonerLevel               int          `json:"summonerLevel"`
	TagLine                     string       `json:"tagLine"`
	Unnamed                     bool         `json:"unnamed"`
	XpSinceLastLevel            int          `json:"xpSinceLastLevel"`
	XpUntilNextLevel            int          `json:"xpUntilNextLevel"`
}
type RerollPoints struct {
	CurrentPoints    int `json:"currentPoints"`
	MaxRolls         int `json:"maxRolls"`
	NumberOfRolls    int `json:"numberOfRolls"`
	PointsCostToRoll int `json:"pointsCostToRoll"`
	PointsToReroll   int `json:"pointsToReroll"`
}
