package vindinium

import (
	"encoding/json"
	"testing"
)

const sampleStateJSON = `
{
   "game":{
      "id":"s2xh3aig",
      "turn":1100,
      "maxTurns":1200,
      "heroes":[
         {
            "id":1,
            "name":"vjousse",
            "userId":"j07ws669",
            "elo":1200,
            "pos":{
               "x":5,
               "y":6
            },
            "life":60,
            "gold":0,
            "mineCount":0,
            "spawnPos":{
               "x":5,
               "y":6
            },
            "crashed":true
         },
         {
            "id":2,
            "name":"vjousse",
            "userId":"j07ws669",
            "elo":1200,
            "pos":{
               "x":12,
               "y":6
            },
            "life":100,
            "gold":0,
            "mineCount":0,
            "spawnPos":{
               "x":12,
               "y":6
            },
            "crashed":true
         },
         {
            "id":3,
            "name":"vjousse",
            "userId":"j07ws669",
            "elo":1200,
            "pos":{
               "x":12,
               "y":11
            },
            "life":80,
            "gold":0,
            "mineCount":0,
            "spawnPos":{
               "x":12,
               "y":11
            },
            "crashed":true
         },
         {
            "id":4,
            "name":"vjousse",
            "userId":"j07ws669",
            "elo":1200,
            "pos":{
               "x":4,
               "y":8
            },
            "lastDir": "South",
            "life":38,
            "gold":1078,
            "mineCount":6,
            "spawnPos":{
               "x":5,
               "y":11
            },
            "crashed":false
         }
      ],
      "board":{
         "size":18,
         "tiles":"##############        ############################        ##############################    ##############################$4    $4############################  @4    ########################  @1##    ##    ####################  []        []  ##################        ####        ####################  $4####$4  ########################  $4####$4  ####################        ####        ##################  []        []  ####################  @2##    ##@3  ########################        ############################$-    $-##############################    ##############################        ############################        ##############"
      },
      "finished":true
   },
   "hero":{
      "id":4,
      "name":"vjousse",
      "userId":"j07ws669",
      "elo":1200,
      "pos":{
         "x":4,
         "y":8
      },
      "lastDir": "South",
      "life":38,
      "gold":1078,
      "mineCount":6,
      "spawnPos":{
         "x":5,
         "y":11
      },
      "crashed":false
   },
   "token":"lte0",
   "viewUrl":"http://localhost:9000/s2xh3aig",
   "playUrl":"http://localhost:9000/api/s2xh3aig/lte0/play"
}
`

func assertUnmarshalling(t *testing.T, name string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("%q was incorrectly unmarshalled. Expected '%v', got '%v'.", name, expected, actual)
	}
}

func TestUnmarshalling(t *testing.T) {
	var s State

	err := json.Unmarshal([]byte(sampleStateJSON), &s)
	if err != nil {
		t.Errorf("Error while unmarshalling: %v", err.Error())
	}

	assertUnmarshalling(t, "Game.ID", "s2xh3aig", s.Game.ID)
	assertUnmarshalling(t, "Game.Turn", 1100, s.Game.Turn)
	assertUnmarshalling(t, "Game.MaxTurns", 1200, s.Game.MaxTurns)

	assertUnmarshalling(t, "Game.Heroes[0].ID", 1, s.Game.Heroes[0].ID)
	assertUnmarshalling(t, "Game.Heroes[0].Name", "vjousse", s.Game.Heroes[0].Name)
	assertUnmarshalling(t, "Game.Heroes[0].UserID", "j07ws669", s.Game.Heroes[0].UserID)
	assertUnmarshalling(t, "Game.Heroes[0].Elo", 1200, s.Game.Heroes[0].Elo)
	assertUnmarshalling(t, "Game.Heroes[0].Pos", Position{5, 6}, s.Game.Heroes[0].Pos)
	assertUnmarshalling(t, "Game.Heroes[0].LastDir", Direction(""), s.Game.Heroes[0].LastDir)
	assertUnmarshalling(t, "Game.Heroes[0].Life", 60, s.Game.Heroes[0].Life)
	assertUnmarshalling(t, "Game.Heroes[0].Gold", 0, s.Game.Heroes[0].Gold)
	assertUnmarshalling(t, "Game.Heroes[0].MineCount", 0, s.Game.Heroes[0].MineCount)
	assertUnmarshalling(t, "Game.Heroes[0].SpawnPos", Position{5, 6}, s.Game.Heroes[0].SpawnPos)
	assertUnmarshalling(t, "Game.Heroes[0].Crashed", true, s.Game.Heroes[0].Crashed)

	assertUnmarshalling(t, "Game.Heroes[1].ID", 2, s.Game.Heroes[1].ID)
	assertUnmarshalling(t, "Game.Heroes[1].Name", "vjousse", s.Game.Heroes[1].Name)
	assertUnmarshalling(t, "Game.Heroes[1].UserID", "j07ws669", s.Game.Heroes[1].UserID)
	assertUnmarshalling(t, "Game.Heroes[1].Elo", 1200, s.Game.Heroes[1].Elo)
	assertUnmarshalling(t, "Game.Heroes[1].Pos", Position{12, 6}, s.Game.Heroes[1].Pos)
	assertUnmarshalling(t, "Game.Heroes[1].LastDir", Direction(""), s.Game.Heroes[1].LastDir)
	assertUnmarshalling(t, "Game.Heroes[1].Life", 100, s.Game.Heroes[1].Life)
	assertUnmarshalling(t, "Game.Heroes[1].Gold", 0, s.Game.Heroes[1].Gold)
	assertUnmarshalling(t, "Game.Heroes[1].MineCount", 0, s.Game.Heroes[1].MineCount)
	assertUnmarshalling(t, "Game.Heroes[1].SpawnPos", Position{12, 6}, s.Game.Heroes[1].SpawnPos)
	assertUnmarshalling(t, "Game.Heroes[1].Crashed", true, s.Game.Heroes[1].Crashed)

	assertUnmarshalling(t, "Game.Heroes[2].ID", 3, s.Game.Heroes[2].ID)
	assertUnmarshalling(t, "Game.Heroes[2].Name", "vjousse", s.Game.Heroes[2].Name)
	assertUnmarshalling(t, "Game.Heroes[2].UserID", "j07ws669", s.Game.Heroes[2].UserID)
	assertUnmarshalling(t, "Game.Heroes[2].Elo", 1200, s.Game.Heroes[2].Elo)
	assertUnmarshalling(t, "Game.Heroes[2].Pos", Position{12, 11}, s.Game.Heroes[2].Pos)
	assertUnmarshalling(t, "Game.Heroes[2].LastDir", Direction(""), s.Game.Heroes[2].LastDir)
	assertUnmarshalling(t, "Game.Heroes[2].Life", 80, s.Game.Heroes[2].Life)
	assertUnmarshalling(t, "Game.Heroes[2].Gold", 0, s.Game.Heroes[2].Gold)
	assertUnmarshalling(t, "Game.Heroes[2].MineCount", 0, s.Game.Heroes[2].MineCount)
	assertUnmarshalling(t, "Game.Heroes[2].SpawnPos", Position{12, 11}, s.Game.Heroes[2].SpawnPos)
	assertUnmarshalling(t, "Game.Heroes[2].Crashed", true, s.Game.Heroes[2].Crashed)

	assertUnmarshalling(t, "Game.Heroes[3].ID", 4, s.Game.Heroes[3].ID)
	assertUnmarshalling(t, "Game.Heroes[3].Name", "vjousse", s.Game.Heroes[3].Name)
	assertUnmarshalling(t, "Game.Heroes[3].UserID", "j07ws669", s.Game.Heroes[3].UserID)
	assertUnmarshalling(t, "Game.Heroes[3].Elo", 1200, s.Game.Heroes[3].Elo)
	assertUnmarshalling(t, "Game.Heroes[3].Pos", Position{4, 8}, s.Game.Heroes[3].Pos)
	assertUnmarshalling(t, "Game.Heroes[3].LastDir", South, s.Game.Heroes[3].LastDir)
	assertUnmarshalling(t, "Game.Heroes[3].Life", 38, s.Game.Heroes[3].Life)
	assertUnmarshalling(t, "Game.Heroes[3].Gold", 1078, s.Game.Heroes[3].Gold)
	assertUnmarshalling(t, "Game.Heroes[3].MineCount", 6, s.Game.Heroes[3].MineCount)
	assertUnmarshalling(t, "Game.Heroes[3].SpawnPos", Position{5, 11}, s.Game.Heroes[3].SpawnPos)
	assertUnmarshalling(t, "Game.Heroes[3].Crashed", false, s.Game.Heroes[3].Crashed)

	assertUnmarshalling(t, "Game.Board.Size", 18, s.Game.Board.Size)

	// Do a couple of spot checks of the board
	if s.Game.Board.TileAt(Position{0, 0}) != WallTile {
		t.Errorf("Expected tile at (0,0) to be a WallTile, got %v.", s.Game.Board.TileAt(Position{0, 0}))
	}

	if s.Game.Board.TileAt(Position{0, 7}) != AirTile {
		t.Errorf("Expected tile at (0,7) to be an AirTile, got %v.", s.Game.Board.TileAt(Position{0, 7}))
	}

	if s.Game.Board.TileAt(Position{4, 8}) != HeroTile {
		t.Errorf("Expected tile at (4,8) to be a HeroTile, got %v.", s.Game.Board.TileAt(Position{4, 8}))
	}

	// The tile north-west of (4,8) should be a mine
	if s.Game.Board.TileAt(s.Game.Board.To(s.Game.Board.To(Position{4, 8}, North), West)) != MineTile {
		t.Errorf("Expected the tile north-west of (4,8) to be a MineTile, got %v.", s.Game.Board.TileAt(s.Game.Board.To(s.Game.Board.To(Position{4, 8}, North), West)))
	}

	assertUnmarshalling(t, "Game.Finished", true, s.Game.Finished)

	assertUnmarshalling(t, "Game.Hero.ID", 4, s.Hero.ID)
	assertUnmarshalling(t, "Game.Hero.Name", "vjousse", s.Hero.Name)
	assertUnmarshalling(t, "Game.Hero.UserID", "j07ws669", s.Hero.UserID)
	assertUnmarshalling(t, "Game.Hero.Elo", 1200, s.Hero.Elo)
	assertUnmarshalling(t, "Game.Hero.Pos", Position{4, 8}, s.Hero.Pos)
	assertUnmarshalling(t, "Game.Hero.LastDir", South, s.Hero.LastDir)
	assertUnmarshalling(t, "Game.Hero.Life", 38, s.Hero.Life)
	assertUnmarshalling(t, "Game.Hero.Gold", 1078, s.Hero.Gold)
	assertUnmarshalling(t, "Game.Hero.MineCount", 6, s.Hero.MineCount)
	assertUnmarshalling(t, "Game.Hero.SpawnPos", Position{5, 11}, s.Hero.SpawnPos)
	assertUnmarshalling(t, "Game.Hero.Crashed", false, s.Hero.Crashed)

	assertUnmarshalling(t, "Token", "lte0", s.Token)
	assertUnmarshalling(t, "ViewURL", "http://localhost:9000/s2xh3aig", s.ViewURL)
	assertUnmarshalling(t, "PlayURL", "http://localhost:9000/api/s2xh3aig/lte0/play", s.PlayURL)
}
