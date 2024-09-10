+---------------------+
|       Player        |
+---------------------+
| - Name: string      |
| - Runs: int         |
| - Fours: int        |
| - Sixes: int        |
| - BallsFaced: int   |
| - IsOut: bool       |
+---------------------+
| + UpdateStats()     |
+---------------------+

         1
         |
         | has many
         |
+---------------------+
|       Team          |
+---------------------+
| - Name: string      |
| - Players: []Player |
| - TotalRuns: int    |
| - Wickets: int      |
| - Extras: int       |
+---------------------+
| + AddRun()          |
| + AddWicket()       |
+---------------------+

         1
         |
         | has
         |
+---------------------+
|        Ball         |
+---------------------+
| - Runs: int         |
| - IsWide: bool      |
| - IsNoBall: bool    |
| - IsWicket: bool    |
+---------------------+

        *        
         | is a part of
         |
+---------------------+
|        Over         |
+---------------------+
| - Balls: []Ball     |
+---------------------+

         1
         |
         | contains two
         |
+---------------------+
|       Match         |
+---------------------+
| - Team1: Team       |
| - Team2: Team       |
| - CurrentOver: int  |
| - Overs: []Over     |
+---------------------+
| + PrintScorecard()  |
| + PlayBall()        |
| + EndOver()         |
+---------------------+
