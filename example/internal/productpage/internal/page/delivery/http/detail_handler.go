package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/loki-zhou/arthas/example/internal/productpage/internal/domain"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type DetailHandler struct {
	DetailUsecase domain.DetailUsecase
	logger     *zap.Logger
}

func NewDetailHandler(logger *zap.Logger,detailUsecase domain.DetailUsecase) *DetailHandler{
	handler := &DetailHandler{
		DetailUsecase: detailUsecase,
		logger: logger,
	}
	return handler
}

type DetailHTTPDeliveryFn func(e *gin.Engine)

func NewDetailHTTPDeliveryFn(handler *DetailHandler) DetailHTTPDeliveryFn {
	return func(e *gin.Engine) {
		e.GET("/api/v1/products/:product_id", handler.GetProductDetail)
	}
}

var ProviderSet = wire.NewSet(NewDetailHandler, NewDetailHTTPDeliveryFn )

func (d *DetailHandler)GetProductDetail(c *gin.Context) {
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