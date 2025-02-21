package dto

type UserFilterDto struct {
	BranchId *string `json:"branch_id"`
	Start    *string `json:"start"`
	End      *string `json:"end"`
}
