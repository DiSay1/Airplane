package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// AquaAffinity is a helmet enchantment that increases underwater mining speed.
type AquaAffinity struct{}

// Name ...
func (AquaAffinity) Name() string {
	return "Aqua Affinity"
}

// MaxLevel ...
func (AquaAffinity) MaxLevel() int {
	return 1
}

// MinCost ...
func (e AquaAffinity) MinCost(int) int {
	return 1
}

// MaxCost ...
func (e AquaAffinity) MaxCost(int) int {
	return 41
}

// Rarity ...
func (AquaAffinity) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithOther ...
func (AquaAffinity) CompatibleWithOther(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (AquaAffinity) CompatibleWithItem(i world.Item) bool {
	h, ok := i.(item.HelmetType)
	return ok && h.Helmet()
}
