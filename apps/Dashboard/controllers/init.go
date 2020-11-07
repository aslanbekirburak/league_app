package controllers

import
	(
		"fmt"
		"league/apps/Dashboard/interactors"
		"league/apps/Dashboard/interfaces"
		"league/entities"
		"time"
		"math/rand"
		"net/http"
		"encoding/json"
	)

var Game interfaces.GameUseCases
var GameTime, TeamCount int

type GameRules struct {}

func(gr *GameRules) StartApplication(res http.ResponseWriter,req *http.Request) {
	Game = &interactors.Game{}
	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")
	res.Header().Set("Access-Control-Allow-Origin", "*") 
	
	// gr.CheckTime(res)
}

func Randomize ()(randomNumber int){
	return rand.Intn(30)+2
}

func(gr *GameRules) CreateRivals()(rivals map[int]int){
	TeamCount = 15
	m := make(map[int]bool)
	matches := make(map[int]int)
    var rival int	
	for i := 2; i < TeamCount ; i++ {
		rival = Randomize()
		if m[rival] != true  {
			matches[i] = rival
		}
		m[rival] = true
		m[i] = true
	}
	return matches
}

func(gr *GameRules) CheckTime(res http.ResponseWriter) error {
	GameTime = 24
	flusher,_:=res.(http.Flusher)
	rivals := gr.CreateRivals()
	for i := 0; i < GameTime; i++ {

		for k, v := range rivals { 
			if i%2 == 0{
				gr.AddScoreToTeam(k)
			}else{
				gr.AddScoreToTeam(v)
			}
		}
		teams,err := Game.GetAllTeam()
		if err != nil {
			fmt.Println("TeamPlayer",err)
		}
		data, err := json.Marshal(teams)
		if err != nil {
			fmt.Println("TeamPlayer",err)
		}
		res.Write(data)
		fmt.Fprintf(res, "%s\n", "SONUCA ulaşıldı")
		flusher.Flush()

		var d = 5000 * time.Microsecond
		time.Sleep(d)
	}
	return nil
}

func(gr *GameRules) AddScoreToTeam(teamId int)(){
	team,err := Game.GetTeam(teamId)
	if err != nil {
		fmt.Println("TeamPlayer",err)
	}
	scoreToAdd := rand.Intn(2)+2
	teamToUpdate := entities.Team{
		Score: team.Score + scoreToAdd,
	}
	err = Game.UpdateTeam(teamId,teamToUpdate)
	if err != nil {
		fmt.Println("UpdateTeam",err)
	}
	players, err := Game.GetPlayerOfTeam(teamId)
	if err != nil {
		fmt.Println("UpdatePlayer",err)
	}
	selectPlayer := rand.Intn(len(players)-1)+1
	gr.AddScoreToPlayer(players[selectPlayer], scoreToAdd)
}

func(gr *GameRules) AddScoreToPlayer(player entities.Player,scoreToAdd int)(){

	scoreTwoOfPlayer := player.TwoPoints
	scoreThreeOfPlayer := player.ThreePoints
	if scoreToAdd == 2 {
		scoreTwoOfPlayer = scoreTwoOfPlayer + 1
	}else if scoreToAdd == 3 {
		scoreThreeOfPlayer = scoreThreeOfPlayer + 1
	}
	playerToUpdate := entities.Player{
		Score: player.Score + scoreToAdd,
		TwoPoints: scoreTwoOfPlayer,
		ThreePoints: scoreThreeOfPlayer,
	}
	
	err := Game.UpdatePlayer(player.Id,playerToUpdate)
	fmt.Println(err)	
}

