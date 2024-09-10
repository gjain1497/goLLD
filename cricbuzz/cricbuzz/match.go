package cricbuzz

type Match struct {
	Team1 Team
	Team2 Team
	CurrentOver int
	Overs []Over
}

func PrintScoreCard(match Match) {
	//complete this
}

func PlayBall(match Match, run int) {
	//complete this
}

func EndOver(match Match) {
	//complete this
}