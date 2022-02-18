package Models

type PrizeDB struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (prize *PrizeDB) TableName() string {
	return "tbl_prize_db"
}

type AddPrizeDB struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (prize *AddPrizeDB) TableName() string {
	return "tbl_prize_db"
}

type UpdatePrizeDB struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
