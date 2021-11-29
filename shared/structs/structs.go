package structs

import globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"

type AddSegmentsReq struct {
	Ts []globalStructs.SongData `json:"ts"`
	M3H8 globalStructs.SongData `json:"m3h8"`
	SongData globalStructs.Song `json:"song_data"`
}

type AddSegmentsResp struct {
	OK bool `json:"ok"`
	Error string `json:"error"`
}

type GetAllSongsResp struct {
	Error string `json:"error"`
	Songs []globalStructs.Song `json:"songs"`
}

type GetSegmentReq struct {
	ID string `json:"id"`
}

type GetSegmentResp struct {
	Segment globalStructs.SongData `json:"segment"`
	Error string `json:"error"`
}
