package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Sharpness is an enchantment applied to a sword or axe that increases melee damage.
type Sharpness struct{}

// Addend returns the additional damage when attacking with sharpness.
func (Sharpness) Addend(level int) float64 {
	return float64(level) * 1.25
}

// Name ...
func (Sharpness) Name() string {
	return "Sharpness"
}

// MaxLevel ...
func (Sharpness) MaxLevel() int {
	return 5
}

// MinCost ...
func (Sharpness) MinCost(level int) int {
	return 1 + (level-1)*11
}

// MaxCost ...
func (s Sharpness) MaxCost(level int) int {
	return s.MinCost(level) + 20
}

// Rarity ...
func (Sharpness) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityCommon
}

// CompatibleWithOther ...
func (Sharpness) CompatibleWithOther(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (Sharpness) CompatibleWithItem(i world.Item) bool {
	t, ok := i.(item.Tool)
	return ok && (t.ToolType() == item.TypeSword || t.ToolType() == item.TypeAxe)
}
