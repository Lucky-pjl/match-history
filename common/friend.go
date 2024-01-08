package common

type Friend struct {
	Availability            string `json:"availability"`
	DisplayGroupID          int    `json:"displayGroupId"`
	DisplayGroupName        string `json:"displayGroupName"`
	GameName                string `json:"gameName"`
	GameTag                 string `json:"gameTag"`
	GroupID                 int    `json:"groupId"`
	GroupName               string `json:"groupName"`
	Icon                    int    `json:"icon"`
	ID                      string `json:"id"`
	IsP2PConversationMuted  bool   `json:"isP2PConversationMuted"`
	LastSeenOnlineTimestamp string `json:"lastSeenOnlineTimestamp"`
	Lol                     Lol    `json:"lol"`
	Name                    string `json:"name"`
	Note                    string `json:"note"`
	Patchline               string `json:"patchline"`
	Pid                     string `json:"pid"`
	PlatformID              string `json:"platformId"`
	Product                 string `json:"product"`
	ProductName             string `json:"productName"`
	Puuid                   string `json:"puuid"`
	StatusMessage           string `json:"statusMessage"`
	Summary                 string `json:"summary"`
	SummonerID              int64  `json:"summonerId"`
	Time                    int64  `json:"time"`
}
type Lol struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type FriendGroup struct {
	Collapsed   bool   `json:"collapsed"`
	ID          int    `json:"id"`
	IsLocalized bool   `json:"isLocalized"`
	IsMetaGroup bool   `json:"isMetaGroup"`
	Name        string `json:"name"`
	Priority    int    `json:"priority"`
}
