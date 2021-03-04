package http

import (
	"github.com/gin-gonic/gin"
	"github.com/loki-zhou/arthas/example/internal/productpage/internal/domain"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type DetailHander struct {
	DetailUsecase domain.DetailUsecase
	logger     *zap.Logger
}

func NewDetailHandler(e *gin.Engine, logger *zap.Logger,detailUsecase domain.DetailUsecase) {
	handler := &DetailHander{
		DetailUsecase: detailUsecase,
		logger: logger,
	}
	e.GET("/api/v1/products/:product_id", handler.GetProductDetail)
}

func (d *DetailHander)GetProductDetail(c *gin.Context) {
	product_id, err := strconv.ParseInt(c.Param("product_id"), 10, 32)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	detail, err := d.DetailUsecase.GetByID(c.Request.Context(), int32(product_id))
	if err != nil {
		d.logger.Error("get detail id error", zap.Error(err))
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, detail)
}