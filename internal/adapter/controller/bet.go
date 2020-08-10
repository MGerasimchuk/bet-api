package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/MGerasimchuk/bet-api/internal/usecase"
)

type BetController struct {
	betUsecase *usecase.BetUsecase
	logger     *logrus.Logger
}

func NewBetController(betUsecase *usecase.BetUsecase, logger *logrus.Logger) *BetController {
	return &BetController{betUsecase: betUsecase, logger: logger}
}

func (c *BetController) BetPost(ctx *gin.Context) {
	accountId := ctx.Param("accountId")
	source := ctx.GetHeader("Source-type")

	var b bet
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amountNumber, err := b.GetAmountAsNumber()
	if err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "value of \"amount\" field should be a number",
		})
		return
	}

	err = c.betUsecase.BetResultProcess(accountId, amountNumber, b.TransactionID, b.State, source)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

type bet struct {
	State         string `json:"state" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	TransactionID string `json:"transactionId" binding:"required"`
}

func (b *bet) GetAmountAsNumber() (float64, error) {
	amount, err := strconv.ParseFloat(b.Amount, 64)

	return amount, err
}
