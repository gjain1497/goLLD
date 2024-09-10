package cricbuzz

type Player struct{
	Name string
	Id string
	Fours int
	Sixes int
	BallsFaced int
	IsOut bool
}

func updateStats(player *Player, runs int, isFour bool, isSix bool){
	player.BallsFaced++
	if isFour {
		player.Fours++
	}
	if isSix {
		player.Sixes++
	}
}

