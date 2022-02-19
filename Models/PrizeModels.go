package Models

type Prize struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Value      int    `json:"value"`
	Percentage int    `json:"percentage"`
}

func (prize *Prize) TableName() string {
	return "v_prize"
}

type AddPrize struct {
	Name       string `json:"name"`
	Value      int    `json:"value"`
	Percentage int    `json:"percentage"`
}

func (prize *AddPrize) TableName() string {
	return "tbl_prize"
}

type InputAddPrize struct {
	Name       string `json:"name" binding:"required"`
	Value      int    `json:"value" binding:"required"`
	Percentage int    `json:"percentage" binding:"required"`
}

type UpdatePrize struct {
	Name       string `json:"name"`
	Value      int    `json:"value"`
	Percentage int    `json:"percentage"`
}

func (prize *UpdatePrize) TableName() string {
	return "tbl_prize"
}

type InputUpdatePrize struct {
	Name       string `json:"name" binding:"required"`
	Value      int    `json:"value" binding:"required"`
	Percentage int    `json:"percentage" binding:"required"`
}

type GetPrizeNum struct {
	Number int `json:"number"`
}
