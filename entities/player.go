package entities

type Player struct {
	Id             int    `json:"id"`
	TeamId          int `json:"team_id"`
	Name         string    `json:"name"`
	Asists          int `json:"asists"`
	TwoPoints     int    `json:"two_points"`
	ThreePoints    int    `json:"three_points"`
	Score int `json:"score"`
}