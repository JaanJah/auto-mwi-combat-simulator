package models_mappings

type DefaultCombatTriggers struct {
	DependencyHrid string `json:"dependencyHrid"`
	ConditionHrid  string `json:"conditionHrid"`
	ComparatorHrid string `json:"comparatorHrid"`
	Value          int    `json:"value"`
}
