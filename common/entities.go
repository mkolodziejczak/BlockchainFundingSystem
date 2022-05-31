package common

import "time"

type User struct {
	Id            int
	Name          string
	Login         string
	Password      string
	FirstName     string
	LastName      string
	Email         string
	WalletAddress string
}

type Project struct {
	Id                   int
	OwnerId              int
	Title                string
	Description          string
	GoalBacking          float64
	SmartContractAddress string
	Milestone1Date       time.Time
	Milestone2Date       time.Time
	Milestone3Date       time.Time
	GoalDate             time.Time
}

type Backing struct {
	BackerId  int
	ProjectId int
	Amount    float64
}
