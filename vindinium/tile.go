package vindinium

// Tile represents the type of a tile on the board
type Tile byte

const (
	// AirTile represents a walkable space
	AirTile Tile = iota

	// HeroTile represents a tile a hero is standing on
	HeroTile

	// MineTile represents a Mine
	MineTile

	// TavernTile represents a Tavern
	TavernTile

	// WallTile represents an impassable space
	WallTile
)
