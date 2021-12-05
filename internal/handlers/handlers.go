package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/spotify-web-project/spotify-db/internal/service"
	"github.com/supperdoggy/spotify-web-project/spotify-db/shared/structs"
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type Handlers struct {
	s      service.IService
	logger *zap.Logger
}

func NewHandlers(s service.IService, l *zap.Logger) Handlers {
	return Handlers{
		s:      s,
		logger: l,
	}
}

func (h *Handlers) AddSegments(c *gin.Context) {
	var req structs.AddSegmentsReq
	var resp structs.AddSegmentsResp
	if err := c.Bind(&req); err != nil {
		data, _ := ioutil.ReadAll(c.Request.Body)
		h.logger.Error("error binding req", zap.Error(err), zap.String("data", string(data)))
		resp.Error = "error binding req"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.NewSegments(req)
	if err != nil {
		h.logger.Error("error NewSegments()", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetAllSongs(c *gin.Context) {
	resp, err := h.s.GetAllSongs()
	if err != nil {
		h.logger.Error("error getting all songs", zap.Error(err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetSegment(c *gin.Context) {
	var req structs.GetSegmentReq
	var resp structs.GetSegmentResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.GetSegment(req)
	if err != nil {
		h.logger.Error("error getting segment", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) NewUser(c *gin.Context) {
	var req globalStructs.User
	var resp structs.NewUserResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.NewUser(req)
	if err != nil {
		h.logger.Error("error getting segment", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetUser(c *gin.Context) {
	var req structs.GetUserReq
	var resp structs.GetUserResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.GetUser(req)
	if err != nil {
		h.logger.Error("error getting segment", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
