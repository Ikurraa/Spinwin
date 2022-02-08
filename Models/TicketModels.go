package Models

type ViewTicket struct {
	Id            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playername"`
	Ticket_status int    `json:"ticketstatus"`
	Created_by    string `json:"createdby"`
	Reward        string `json:"reward"`
}

func (ticket *ViewTicket) TableName() string {
	return "v_ticket"
}

type AddTicket struct {
	Ticket_code   string `json:"ticketcode" binding:"required,min=6,max=6,alphanum"`
	Player_name   string `json:"playername" binding:"required,max=30"`
	Ticket_status int    `json:"ticketstatus"`
	Redeem_status int    `json:"redeemstatus"`
	Created_by    string `json:"createdby"`
}

func (ticket *AddTicket) TableName() string {
	return "tbl_ticket"
}

type InputTicket struct {
	Ticket_code string `json:"ticketcode" binding:"required,min=6,max=6,alphanum"`
	Player_name string `json:"playername" binding:"required,max=30,alphanum"`
}

type DeleteTicket struct {
	ID          int    `json:"id"`
	Ticket_code string `json:"ticketcode"`
	Status      int    `json:"status"`
	Update_by   string `json:"updateby"`
	Updated_at  string `json:"updateat"`
}

func (ticket *DeleteTicket) TableName() string {
	return "tbl_ticket"
}

type InputDeleteTicket struct {
	Status int `json:"status" binding:"required"`
}

type UpdateTicket struct {
	ID            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playername"`
	Status        int    `json:"status"`
	Ticket_status int    `json:"ticketstatus"`
	Update_by     string `json:"updateby"`
	Updated_at    string `json:"updatedat"`
}

func (ticket *UpdateTicket) TableName() string {
	return "tbl_ticket"
}

type InputUpdateTicket struct {
	Ticket_code string `json:"ticketcode"`
	Player_name string `json:"playername"`
}

type CheckTicket struct {
	Id            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playername"`
	Ticket_status int    `json:"ticketstatus"`
	Status        int    `json:"status"`
}

func (ticket *CheckTicket) TableName() string {
	return "tbl_ticket"
}

type InputCheckTicket struct {
	Ticket_Code string `json:"ticketcode"`
	Player_name string `json:"playername"`
}

type ClaimTicket struct {
	Id            uint   `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playerticket"`
	Redeem_at     string `json:"redeemat"`
	Ticket_status int    `json:"ticketstatus"`
	Status        int    `json:"status"`
}

func (ticket *ClaimTicket) TableName() string {
	return "tbl_ticket"
}

type RedeemTicket struct {
	Id            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Redeem_status int    `json:"redeemstatus"`
	Ticket_status int    `json:"ticketstatus"`
	Status        int    `json:"status"`
	Updated_at    string `json:"updatedat"`
	Update_by     string `json:"updateby"`
}

func (ticket *RedeemTicket) TableName() string {
	return "tbl_ticket"
}

type InputRedeemTicket struct {
	Ticket_Code string `json:"ticketcode"`
}

type ViewUsedTicket struct {
	Id            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playername"`
	Ticket_status int    `json:"ticketstatus"`
	Created_by    string `json:"createdby"`
	Reward        string `json:"reward"`
}

func (ticket *ViewUsedTicket) TableName() string {
	return "v_ticket"
}

type ViewUnusedTicket struct {
	Id            int    `json:"id"`
	Ticket_code   string `json:"ticketcode"`
	Player_name   string `json:"playername"`
	Ticket_status int    `json:"ticketstatus"`
	Created_by    string `json:"createdby"`
	Reward        string `json:"reward"`
}

func (ticket *ViewUnusedTicket) TableName() string {
	return "v_ticket"
}
