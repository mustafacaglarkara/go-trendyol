package Dtos

type QuestionFilter struct {
	Barcode          string           `url:"barcode,omitempty"`
	Page             int              `url:"page,omitempty"`
	Size             int              `url:"size,omitempty"`
	EndDate          int              `url:"endDate,omitempty"`
	StartDate        int              `url:"startDate,omitempty"`
	Status           Status           `url:"status,omitempty"`
	OrderByField     OrderByField     `url:"orderByField,omitempty"`
	OrderByDirection OrderByDirection `url:"orderByDirection,omitempty"`
}
type Status string
type OrderByField string
type OrderByDirection string

const (
	ASC  OrderByDirection = "ASC"
	DESC OrderByDirection = "DESC"
)
const (
	LastModifiedDate OrderByField = "LastModifiedDate"
	CreatedDate      OrderByField = "CreatedDate"
)
const (
	WaitingForAnswer  Status = "WAITING_FOR_ANSWER"
	WaitingForApprove Status = "WAITING_FOR_APPROVE"
	Answered          Status = "ANSWERED"
	Reported          Status = "REPORTED"
	Rejected          Status = "REJECTED"
)
