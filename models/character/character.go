package models

type Character struct {
	Skills         []Skill   `json:"skills"`
	Abilities      []Ability `json:"abilities"`
	Items          []Item    `json:"items"`
	HouseBuffs     []Buff    `json:"houseBuffs"`
	EquipmentBuffs []Buff    `json:"equipmentBuffs"`
}
