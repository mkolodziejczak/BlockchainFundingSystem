package api

import (
	"fmt"
	"fundingSystem/common"
	"fundingSystem/db"
	"fundingSystem/ethereum"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strings"
	"time"
)

func RegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json common.UserDto

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		session := sessions.Default(c)
		id, _ := db.CreateUser(json)
		session.Set(gin.AuthUserKey, id)
	}
}

func CreateProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json common.ProjectDto

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		session := sessions.Default(c)
		userId := session.Get(gin.AuthUserKey).(int)
		creatorAddress, _ := db.GetUserWalletId(userId)

		goalDate := time.Now().AddDate(0, 0, 30)
		floatGoal, _ := parseBigFloat(json.GoalBacking)
		goalAmount := etherToWei(floatGoal)
		m1Date := dateToBigInt(json.Milestone1Date.Unix())
		m2Date := dateToBigInt(json.Milestone1Date.Unix())
		m3Date := dateToBigInt(json.Milestone1Date.Unix())

		address := ethereum.DeployContract(creatorAddress, goalAmount, m1Date, m2Date, m3Date,
			dateToBigInt(goalDate.Unix()))
		id, _ := db.CreateProject(json, userId, address.Hex())

		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/project?id=%d", id))
	}
}

func dateToBigInt(date int64) *big.Int {
	return new(big.Int).SetInt64(date)
}

func parseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func etherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}
