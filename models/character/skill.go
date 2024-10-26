package models_character

type Skill struct {
	CharacterId       int     `json:"characterId"`
	SkillHrid         string  `json:"skillHrid"`
	Experience        float64 `json:"experience"`
	Level             int     `json:"level"`
	OfflineExperience float64 `json:"offlineExperience"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}
