package dto

import (
	"time"
)

type UserFilterDto struct {
	BranchId  *string    `json:"branch_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
