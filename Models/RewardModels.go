package Models

type Reward struct {
	ID     string `json:"id"`
	Reward string `json:"reward"`
}

func (reward *Reward) TableName() string {
	return "tbl_reward"
}

type InputReward struct {
	Reward string `json:"reward"`
}
