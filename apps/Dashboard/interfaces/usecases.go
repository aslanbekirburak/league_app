package interfaces

import (
	. "league_app/entities"
)

type GameUseCases interface {
	GetPlayer(playerId int)(player *Player, err error)
	UpdatePlayer(playerId int,player Player) error
	GetTeam(teamId int)(team *Team, err error)
	UpdateTeam(teamId int,team Team) error
	GetPlayerOfTeam(teamId int) (teamRes []Player,err error)
	GetAllTeam() (teamRes []Team,err error)
}