package models

type ActionBoost struct {
	UniqueHrid           string `json:"uniqueHrid"`
	TypeHrid             string `json:"typeHrid"`
	RatioBoost           int    `json:"ratioBoost"`
	RatioBoostLevelBonus int    `json:"ratioBoostLevelBonus"`
	FlatBoost            int    `json:"flatBoost"`
	FlatBoostLevelBonus  int    `json:"flatBoostLevelBonus"`
	StartTime            string `json:"startTime"`
	Duration             int    `json:"duration"`
}
