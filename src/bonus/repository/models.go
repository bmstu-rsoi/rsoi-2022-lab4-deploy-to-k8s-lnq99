// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"time"

	"github.com/google/uuid"
)

type Privilege struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Balance  int32  `json:"balance"`
}

type PrivilegeHistory struct {
	ID            int32     `json:"id"`
	PrivilegeID   int32     `json:"privilegeID"`
	TicketUid     uuid.UUID `json:"ticketUid"`
	Datetime      time.Time `json:"datetime"`
	BalanceDiff   int32     `json:"balanceDiff"`
	OperationType string    `json:"operationType"`
}