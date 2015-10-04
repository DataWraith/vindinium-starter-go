package vindinium

import (
	"encoding/json"
	"fmt"
)

// Board represents the current game board
type Board struct {
	Size int

	HeroID    map[Position]int
	MineOwner map[Position]int
	Taverns   map[Position]struct{}

	tiles []Tile
}

// TileAt returns the tile at the given board position. If the given position
// is outside of the board, WallTile is returned.
func (b Board) TileAt(pos Position) Tile {
	if pos.X < 0 || pos.X >= b.Size {
		return WallTile
	}

	if pos.Y < 0 || pos.Y >= b.Size {
		return WallTile
	}

	// The positions sent from the server follow a different convention, so
	// we can't use pos.Y*b.Size + pos.X here
	return b.tiles[pos.X*b.Size+pos.Y]
}

// To returns the position that lies in the direction `dir` of the tile at
// Position `pos`.
func (b Board) To(pos Position, dir Direction) Position {
	switch dir {
	case North:
		return Position{pos.X, pos.Y - 1}
	case East:
		return Position{pos.X + 1, pos.Y}
	case South:
		return Position{pos.X, pos.Y + 1}
	case West:
		return Position{pos.X - 1, pos.Y}
	default:
		return pos
	}
}

// Neighbors returns an array with the positions that lie adjacent to the given
// position.
func (b Board) Neighbors(pos Position) [4]Position {
	return [4]Position{
		b.To(pos, North),
		b.To(pos, East),
		b.To(pos, South),
		b.To(pos, West),
	}
}

// Passable returns whether the given position on the board is passable
func (b Board) Passable(pos Position) bool {
	return b.TileAt(pos) == AirTile
}

func newBoard(size int, tiles string) (Board, error) {
	b := Board{
		Size: size,

		HeroID:    make(map[Position]int),
		MineOwner: make(map[Position]int),
		Taverns:   make(map[Position]struct{}),

		tiles: make([]Tile, size*size),
	}

	if len(tiles) != size*size*2 {
		return Board{}, fmt.Errorf("Board: couldn't parse raw tiles, raw tiles string is of wrong size. Expected %v, got %v.", size*size*2, len(tiles))
	}

	for x := 0; x < b.Size; x++ {
		for y := 0; y < b.Size; y++ {
			idx := y*b.Size + x

			switch tiles[2*idx : 2*idx+2] {
			case "  ":
				b.tiles[idx] = AirTile

			case "##":
				b.tiles[idx] = WallTile

			case "[]":
				b.tiles[idx] = TavernTile
				b.Taverns[Position{x, y}] = struct{}{}

			case "$-":
				b.tiles[idx] = MineTile

			case "$1":
				b.tiles[idx] = MineTile
				b.MineOwner[Position{x, y}] = 1

			case "$2":
				b.tiles[idx] = MineTile
				b.MineOwner[Position{x, y}] = 2

			case "$3":
				b.tiles[idx] = MineTile
				b.MineOwner[Position{x, y}] = 3

			case "$4":
				b.tiles[idx] = MineTile
				b.MineOwner[Position{x, y}] = 4

			case "@1":
				b.tiles[idx] = HeroTile
				b.HeroID[Position{x, y}] = 1

			case "@2":
				b.tiles[idx] = HeroTile
				b.HeroID[Position{x, y}] = 2

			case "@3":
				b.tiles[idx] = HeroTile
				b.HeroID[Position{x, y}] = 3

			case "@4":
				b.tiles[idx] = HeroTile
				b.HeroID[Position{x, y}] = 4

			default:
				return Board{}, fmt.Errorf("Board: Could not parse tiles, unknown tile found: %q", tiles[2*idx:2*idx+2])
			}
		}
	}

	return b, nil
}

// jsonBoard is used to unmarshal the board sent by the server
type jsonBoard struct {
	Size  int
	Tiles string
}

// UnmarshalJSON is called by the JSON unmarshaller. It takes care to parse the
// tile-string sent by the server into the useable Board representation.
func (b *Board) UnmarshalJSON(text []byte) error {
	var jb jsonBoard

	err := json.Unmarshal(text, &jb)
	if err != nil {
		return err
	}

	*b, err = newBoard(jb.Size, jb.Tiles)

	return err
}
