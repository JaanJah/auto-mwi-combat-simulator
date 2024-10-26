package models_mappings

type Ability struct {
	AquaArrow         AbilityInfo `json:"/abilities/aqua_arrow"`
	AquaAura          AbilityInfo `json:"/abilities/aqua_aura"`
	ArcaneReflection  AbilityInfo `json:"/abilities/arcane_reflection"`
	Berserk           AbilityInfo `json:"/abilities/berserk"`
	Cleave            AbilityInfo `json:"/abilities/cleave"`
	CripplingSlash    AbilityInfo `json:"/abilities/crippling_slash"`
	CriticalAura      AbilityInfo `json:"/abilities/critical_aura"`
	ElementalAffinity AbilityInfo `json:"/abilities/elemental_affinity"`
	Elusiveness       AbilityInfo `json:"/abilities/elusiveness"`
	Entangle          AbilityInfo `json:"/abilities/entangle"`
	FierceAura        AbilityInfo `json:"/abilities/fierce_aura"`
	Fireball          AbilityInfo `json:"/abilities/fireball"`
	Firestorm         AbilityInfo `json:"/abilities/firestorm"`
	FlameArrow        AbilityInfo `json:"/abilities/flame_arrow"`
	FlameAura         AbilityInfo `json:"/abilities/flame_aura"`
	FlameBlast        AbilityInfo `json:"/abilities/flame_blast"`
	Frenzy            AbilityInfo `json:"/abilities/frenzy"`
	FrostSurge        AbilityInfo `json:"/abilities/frost_surge"`
	Heal              AbilityInfo `json:"/abilities/heal"`
	IceSpear          AbilityInfo `json:"/abilities/ice_spear"`
	Impale            AbilityInfo `json:"/abilities/impale"`
	Insanity          AbilityInfo `json:"/abilities/insanity"`
	Invincible        AbilityInfo `json:"/abilities/invincible"`
	Maim              AbilityInfo `json:"/abilities/maim"`
	ManaSpring        AbilityInfo `json:"/abilities/mana_spring"`
	MinorHeal         AbilityInfo `json:"/abilities/minor_heal"`
	NaturesVeil       AbilityInfo `json:"/abilities/natures_veil"`
	PenetratingShot   AbilityInfo `json:"/abilities/penetrating_shot"`
	PenetratingStrike AbilityInfo `json:"/abilities/penetrating_strike"`
	PestilentShot     AbilityInfo `json:"/abilities/pestilent_shot"`
	Poke              AbilityInfo `json:"/abilities/poke"`
	Precision         AbilityInfo `json:"/abilities/precision"`
	Promote           AbilityInfo `json:"/abilities/promote"`
	Provoke           AbilityInfo `json:"/abilities/provoke"`
	Puncture          AbilityInfo `json:"/abilities/puncture"`
	QuickAid          AbilityInfo `json:"/abilities/quick_aid"`
	QuickShot         AbilityInfo `json:"/abilities/quick_shot"`
	RainOfArrows      AbilityInfo `json:"/abilities/rain_of_arrows"`
	Rejuvenate        AbilityInfo `json:"/abilities/rejuvenate"`
	Revive            AbilityInfo `json:"/abilities/revive"`
	Scratch           AbilityInfo `json:"/abilities/scratch"`
	SilencingShot     AbilityInfo `json:"/abilities/silencing_shot"`
	Smack             AbilityInfo `json:"/abilities/smack"`
	SmokeBurst        AbilityInfo `json:"/abilities/smoke_burst"`
	SpeedAura         AbilityInfo `json:"/abilities/speed_aura"`
	SpikeShell        AbilityInfo `json:"/abilities/spike_shell"`
	SteadyShot        AbilityInfo `json:"/abilities/steady_shot"`
	StunningBlow      AbilityInfo `json:"/abilities/stunning_blow"`
	Sweep             AbilityInfo `json:"/abilities/sweep"`
	SylvanAura        AbilityInfo `json:"/abilities/sylvan_aura"`
	Taunt             AbilityInfo `json:"/abilities/taunt"`
	Toughness         AbilityInfo `json:"/abilities/toughness"`
	ToxicPollen       AbilityInfo `json:"/abilities/toxic_pollen"`
	Vampirism         AbilityInfo `json:"/abilities/vampirism"`
	WaterStrike       AbilityInfo `json:"/abilities/water_strike"`
}

type AbilityInfo struct {
	Hrid                  string                  `json:"hrid"`
	Name                  string                  `json:"name"`
	Description           string                  `json:"description"`
	IsSpecialAbility      bool                    `json:"isSpecialAbility"`
	ManaCost              int                     `json:"manaCost"`
	CooldownDuration      int                     `json:"cooldownDuration"`
	CastDuration          int                     `json:"castDuration"`
	AbilityEffects        []AbilityEffect         `json:"abilityEffects"`
	DefaultCombatTriggers []DefaultCombatTriggers `json:"defaultCombatTriggers"`
	SortIndex             int                     `json:"sortIndex"`
}

type AbilityEffect struct {
	TargetType                   string        `json:"targetType"`
	EffectType                   string        `json:"effectType"`
	CombatStyleHrid              string        `json:"combatStyleHrid"`
	DamageType                   string        `json:"damageType"`
	BaseDamageFlat               int           `json:"baseDamageFlat"`
	BaseDamageFlatLevelBonus     float64       `json:"baseDamageFlatLevelBonus"`
	BaseDamageRatio              float64       `json:"baseDamageRatio"`
	BaseDamageRatioLevelBonus    float64       `json:"baseDamageRatioLevelBonus"`
	BonusAccuracyRatio           float64       `json:"bonusAccuracyRatio"`
	BonusAccuracyRatioLevelBonus float64       `json:"bonusAccuracyRatioLevelBonus"`
	DamageOverTimeRatio          float64       `json:"damageOverTimeRatio"`
	DamageOverTimeDuration       float64       `json:"damageOverTimeDuration"`
	PierceChance                 float64       `json:"pierceChance"`
	BlindChange                  float64       `json:"blindChance"`
	BlindDuration                float64       `json:"blindDuration"`
	SilenceChance                float64       `json:"silenceChance"`
	SilenceDuration              float64       `json:"silenceDuration"`
	StunChance                   float64       `json:"stunChance"`
	StunDuration                 float64       `json:"stunDuration"`
	SpendHpRatio                 float64       `json:"spendHpRatio"`
	Buffs                        []AbilityBuff `json:"buffs"`
}

type AbilityBuff struct {
	UniqueHrid           string  `json:"unique_Hrid"`
	TypeHrid             string  `json:"typeHrid"`
	ratioBoost           float64 `json:"ratioBoost"`
	ratioBoostLevelBonus float64 `json:"ratioBoostLevelBonus"`
	flatBoost            float64 `json:"flatBoost"`
	flatBoostLevelBonus  float64 `json:"flatBoostLevelBonus"`
	startTime            string  `json:"startTime"`
	duration             float64 `json:"duration"`
}
