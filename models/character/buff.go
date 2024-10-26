package models

type Buff struct {
	Brewing        []ActionBoost `json:"/action_types/brewing"`
	Cheesesmithing []ActionBoost `json:"/action_types/cheesesmithing"`
	Combat         []ActionBoost `json:"/action_types/combat"`
	Cooking        []ActionBoost `json:"/action_types/cooking"`
	Crafting       []ActionBoost `json:"/action_types/crafting"`
	Enhancing      []ActionBoost `json:"/action_types/enhancing"`
	Foraging       []ActionBoost `json:"/action_types/foraging"`
	Milking        []ActionBoost `json:"/action_types/milking"`
	Tailoring      []ActionBoost `json:"/action_types/tailoring"`
	Woodcutting    []ActionBoost `json:"/action_types/woodcutting"`
}
