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

func (h *Handlers) NewPlaylist(c *gin.Context) {
	var req structs.NewPlaylistReq
	var resp structs.NewPlaylistResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.NewPlaylist(req)
	if err != nil {
		h.logger.Error("error creating new playlist", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) DeletePlaylist(c *gin.Context) {
	var req structs.DeleteUserPlaylistReq
	var resp structs.DeleteUserPlaylistResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.DeleteUserPlaylist(req)
	if err != nil {
		h.logger.Error("error creating new playlist", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetUserPlaylists(c *gin.Context) {
	var req structs.GetUserAllPlaylistsReq
	var resp structs.GetUserAllPlaylistsResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.GetUserPlaylists(req)
	if err != nil {
		h.logger.Error("error getting user laylists", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetUserPlaylist(c *gin.Context) {
	var req structs.GetPlaylistReq
	var resp structs.GetPlaylistResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.GetUserPlaylist(req)
	if err != nil {
		h.logger.Error("error getting user playlist playlist", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) AddSongToUserPlaylist(c *gin.Context) {
	var req structs.AddSongToUserPlaylistReq
	var resp structs.AddSongToUserPlaylistResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.AddSongToUserPlaylist(req)
	if err != nil {
		h.logger.Error("error adding new song to playlist", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handlers) RemoveSongFromUserPlaylist(c *gin.Context) {
	var req structs.RemoveSongFromUserPlaylistReq
	var resp structs.RemoveSongFromUserPlaylistResp
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error binding req", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err := h.s.RemoveSongFromUserPlaylist(req)
	if err != nil {
		h.logger.Error("error removing song from user playlist", zap.Error(err), zap.Any("req", req))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
