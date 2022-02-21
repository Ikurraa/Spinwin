package Models

type Reward struct {
	ID     string `json:"id"`
	Reward string `json:"reward"`
	Value  string `json:"value"`
}

func (reward *Reward) TableName() string {
	return "tbl_reward"
}

type InputReward struct {
	Reward string `json:"reward"`
	Value  string `json:"value"`
}
