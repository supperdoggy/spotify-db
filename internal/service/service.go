package service

import (
	"errors"
	"github.com/supperdoggy/spotify-web-project/spotify-db/internal/db"
	"github.com/supperdoggy/spotify-web-project/spotify-db/shared/structs"
	"go.uber.org/zap"
)

type IService interface {
	NewSegments(req structs.AddSegmentsReq) (resp structs.AddSegmentsResp, err error)
	GetAllSongs() (resp structs.GetAllSongsResp, err error)
	GetSegment(req structs.GetSegmentReq) (resp structs.GetSegmentResp, err error)
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
