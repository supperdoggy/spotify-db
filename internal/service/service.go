package service

import (
	"errors"
	"github.com/supperdoggy/spotify-web-project/spotify-db/internal/db"
	"github.com/supperdoggy/spotify-web-project/spotify-db/shared/structs"
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"go.uber.org/zap"
	"time"
)

type IService interface {
	NewSegments(req structs.AddSegmentsReq) (resp structs.AddSegmentsResp, err error)
	GetAllSongs() (resp structs.GetAllSongsResp, err error)
	GetSegment(req structs.GetSegmentReq) (resp structs.GetSegmentResp, err error)
	GetUser(req structs.GetUserReq) (resp structs.GetUserResp, err error)
	NewUser(req globalStructs.User) (resp structs.NewUserResp, err error)
	NewPlaylist(req structs.NewPlaylistReq) (resp structs.NewPlaylistResp, err error)
	DeletePlaylist(req structs.DeleteUserPlaylistReq) (resp structs.DeleteUserPlaylistResp, err error)
	DeleteUserPlaylist(req structs.DeleteUserPlaylistReq) (resp structs.DeleteUserPlaylistResp, err error)
	AddSongToUserPlaylist(req structs.AddSongToUserPlaylistReq) (resp structs.AddSongToUserPlaylistResp, err error)
	RemoveSongFromUserPlaylist(req structs.RemoveSongFromUserPlaylistReq) (resp structs.RemoveSongFromUserPlaylistResp, err error)
	GetUserPlaylists(req structs.GetUserAllPlaylistsReq) (resp structs.GetUserAllPlaylistsResp, err error)
	GetUserPlaylist(req structs.GetPlaylistReq) (resp structs.GetPlaylistResp, err error)
}

type Service struct {
	d      db.IDB
	logger *zap.Logger
}

func NewService(d db.IDB, l *zap.Logger) IService {
	return &Service{d: d, logger: l}
}

func (s *Service) NewSegments(req structs.AddSegmentsReq) (resp structs.AddSegmentsResp, err error) {
	err = s.d.InsertSegment(req.M3H8)
	if err != nil {
		s.logger.Error("error inserting m3h8", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	err = s.d.InsertSegment(req.Ts...)
	if err != nil {
		s.logger.Error("error inserting ts", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	err = s.d.InsertSong(req.SongData)
	if err != nil {
		s.logger.Error("error inserting song data", zap.Error(err), zap.Any("song_data", req.SongData))
		resp.Error = err.Error()
		return resp, err
	}
	resp.OK = true
	return resp, nil
}

func (s *Service) GetAllSongs() (resp structs.GetAllSongsResp, err error) {
	songs, err := s.d.GetAllSongs()
	if err != nil {
		s.logger.Error("error getting songs", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}
	resp.Songs = songs
	return
}

func (s *Service) GetSegment(req structs.GetSegmentReq) (resp structs.GetSegmentResp, err error) {
	if req.ID == "" {
		resp.Error = "id cannot be empty"
		return resp, errors.New(resp.Error)
	}

	segment, err := s.d.GetSegment(req.ID)
	if err != nil {
		s.logger.Error("error getting segment", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		return resp, err
	}

	resp.Segment = segment
	return resp, err
}

func (s *Service) NewUser(req globalStructs.User) (resp structs.NewUserResp, err error) {
	if req.ID == "" {
		resp.Error = "id cannot be empty"
		return resp, errors.New(resp.Error)
	}

	err = s.d.NewUser(req)
	if err != nil {
		s.logger.Error("error creating new user", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		return resp, err
	}

	resp.OK = true
	return resp, nil
}

func (s *Service) GetUser(req structs.GetUserReq) (resp structs.GetUserResp, err error) {
	if req.ID == "" {
		resp.Error = "id cannot be empty"
		return resp, errors.New(resp.Error)
	}

	u, err := s.d.GetUserByID(req.ID)
	if err != nil {
		s.logger.Error("error getting user by id", zap.Error(err), zap.Any("id", req.ID))
		resp.Error = err.Error()
		return resp, err
	}

	resp.User = u
	return resp, nil
}

func (s *Service) NewPlaylist(req structs.NewPlaylistReq) (resp structs.NewPlaylistResp, err error) {
	if req.UserID == "" || req.PlaylistName == "" {
		resp.Error = "you need to fill playlist name and user_id"
		return resp, errors.New(resp.Error)
	}

	p := globalStructs.Playlist{
		Name:        req.PlaylistName,
		Description: req.Description,
		OwnerID:     req.UserID,
		Songs:       []globalStructs.Song{},
		Created:     time.Now(),
		Shared:      req.Shared,
	}

	err = s.d.NewPlaylist(p)
	if err != nil {
		s.logger.Error("error creating new playlist", zap.Error(err), zap.Any("playlist", p))
		resp.Error = err.Error()
		return
	}

	resp.OK = true
	return
}

func (s *Service) DeletePlaylist(req structs.DeleteUserPlaylistReq) (resp structs.DeleteUserPlaylistResp, err error) {
	if req.PlaylistID == "" || req.UserID == "" {
		resp.Error = "you need to fill ids"
		return resp, errors.New(resp.Error)
	}

	err = s.d.DeleteUserPlaylist(req.PlaylistID, req.UserID)
	if err != nil {
		s.logger.Error("error deleting user playlist", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		return
	}

	resp.OK = true
	return

}

func (s *Service) GetUserPlaylist(req structs.GetPlaylistReq) (resp structs.GetPlaylistResp, err error) {
	if req.PlaylistID == "" {
		resp.Error = "ids must not be empty"
		return resp, errors.New(resp.Error)
	}

	resp.Playlist, err = s.d.GetPlaylistByID(req.PlaylistID)
	if err != nil {
		s.logger.Error("error getting user playlist by id", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
	}

	return
}

func (s *Service) GetUserPlaylists(req structs.GetUserAllPlaylistsReq) (resp structs.GetUserAllPlaylistsResp, err error) {
	if req.UserID == "" {
		resp.Error = "ids must not be empty"
		return resp, errors.New(resp.Error)
	}

	resp.Playlists, err = s.d.GetAllUserPlaylists(req.UserID)
	if err != nil {
		s.logger.Error("error getting all user playlists")
		resp.Error = err.Error()
		return resp, err
	}

	return
}

func (s *Service) DeleteUserPlaylist(req structs.DeleteUserPlaylistReq) (resp structs.DeleteUserPlaylistResp, err error) {
	if req.PlaylistID == "" || req.UserID == "" {
		resp.Error = "ids must not be empty"
		return resp, errors.New(resp.Error)
	}

	err = s.d.DeleteUserPlaylist(req.PlaylistID, req.UserID)
	if err != nil {
		s.logger.Error("error deleting user playlist", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		return
	}

	resp.OK = true
	return
}

func (s *Service) AddSongToUserPlaylist(req structs.AddSongToUserPlaylistReq) (resp structs.AddSongToUserPlaylistResp, err error) {
	if req.PlaylistID == "" || req.SongID == "" || req.UserID == "" {
		return
	}

	song, err := s.d.GetSongByID(req.SongID)
	if err != nil {
		s.logger.Error("error getting song by id", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	err = s.d.AddSongsToUserPlaylist(req.PlaylistID, req.UserID, song)
	if err != nil {
		s.logger.Error("error adding song to playlist", zap.Error(err))
		resp.Error = err.Error()
		return resp, err
	}

	resp.OK = true
	return
}

func (s *Service) RemoveSongFromUserPlaylist(req structs.RemoveSongFromUserPlaylistReq) (resp structs.RemoveSongFromUserPlaylistResp, err error) {
	if req.PlaylistID == "" || req.SongID == "" || req.UserID == "" {
		resp.Error = "ids must not be empty"
		return
	}

	err = s.d.RemoveSongFromUserPlaylist(req.PlaylistID, req.UserID, req.SongID)
	if err != nil {
		s.logger.Error("error removing song from playlist", zap.Error(err))
		resp.Error = err.Error()
		return
	}

	resp.OK = true
	return
}
