package entities

type Team struct {
	Id              int    `json:"id"`
	Name           string `json:"name"`
	Score       int `json:"score"`
	AttackCount        int `json:"attack_count"`
	Point           int `json:"point"`
}
