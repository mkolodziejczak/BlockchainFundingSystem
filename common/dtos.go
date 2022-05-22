package common

import "time"

type LoginDto struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDto struct {
	Login         string `json:"login" binding:"required"`
	Password      string `json:"password" binding:"required"`
	FirstName     string `json:"firstName" binding:"required"`
	LastName      string `json:"lastName" binding:"required"`
	Email         string `json:"email" binding:"required"`
	WalletAddress string `json:"walletId" binding:"required"`
}

type ProjectDto struct {
	Title                string    `json:"title" binding:"required"`
	Description          string    `json:"description" binding:"required"`
	GoalBacking          float64   `json:"goal_backing" binding:"required"`
	Milestone1Date       time.Time `json:"milestone_1_date" binding:"required"`
	Milestone2Date       time.Time `json:"milestone_2_date" binding:"required"`
	Milestone3Date       time.Time `json:"milestone_3_date" binding:"required"`
	SmartContractAddress string    `json:"smart_contract_address" binding:"required"`
}
