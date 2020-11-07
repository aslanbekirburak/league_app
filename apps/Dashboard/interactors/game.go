package interactors

import (
	"league/common/gorm"
	"league/entities"
	"fmt"
	gormg "github.com/jinzhu/gorm"

)
var dbGorm *gormg.DB

type Game struct{}

func (g Game) UpdatePlayer(playerId int,player entities.Player) error {
	dbGorm := gorm.GetInstance()
	res := dbGorm.Model(&player).Where("id = ?", playerId).Updates(&player)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	return nil
}

func (g Game) UpdateTeam(teamId int,team entities.Team) error {
	dbGorm := gorm.GetInstance()
	res := dbGorm.Model(&team).Where("id = ?", teamId).Updates(&team)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	return nil
}

func (g Game) GetPlayer(playerId int) (playerRes *entities.Player,err error) {
	dbGorm := gorm.GetInstance()
	var player entities.Player
	res := dbGorm.First(&player, playerId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	return &player,nil
}

func (g Game) GetTeam(teamId int) (teamRes *entities.Team,err error) {
	dbGorm := gorm.GetInstance()
	var team entities.Team
	res := dbGorm.First(&team, teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	return &team,nil
}

func (g Game) GetPlayerOfTeam(teamId int) (teamRes []entities.Player,err error) {
	dbGorm := gorm.GetInstance()
	var players []entities.Player
	res := dbGorm.Where("team_id = ?", teamId).Find(&players)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	return players,nil
}

func (g Game) GetAllTeam() (teamRes []entities.Team,err error) {
	dbGorm := gorm.GetInstance()
	var teams []entities.Team
	res := dbGorm.Find(&teams)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	return teams,nil
}