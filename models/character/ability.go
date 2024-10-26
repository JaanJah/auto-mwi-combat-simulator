package models_character

type Ability struct {
	CharacterId int     `json:"characterId"`
	AbilityHrid string  `json:"abilityHrid"`
	Experience  float64 `json:"experience"`
	Level       int     `json:"level"`
	SlotNumber  int     `json:"slotNumber"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
