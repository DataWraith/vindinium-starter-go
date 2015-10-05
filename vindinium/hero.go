package vindinium

// Hero contains information about a single hero in the game.
type Hero struct {
	ID        HeroID
	Name      string
	UserID    string
	Elo       int
	Pos       Position
	LastDir   Direction // May be empty if the bot playing that hero has crashed
	Life      int
	Gold      int
	MineCount int
	SpawnPos  Position
	Crashed   bool
}
