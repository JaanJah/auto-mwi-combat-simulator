package models_character

type Item struct {
	Id               int    `json:"id"`
	CharacterId      int    `json:"characterId"`
	ItemLocationHrid string `json:"itemLocationHrid"`
	ItemHrid         string `json:"itemHrid"`
	EnchantmentHrid  string `json:"enchantmentHrid"`
	Count            int    `json:"count"`
	OfflineCount     int    `json:"offlineCount"`
	Hash             string `json:"hash"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
