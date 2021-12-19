package structs

import globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"

type AddSegmentsReq struct {
	UserID string `json:"user_id"`
	Ts       []globalStructs.SongData `json:"ts"`
	M3H8     globalStructs.SongData   `json:"m3h8"`
	SongData globalStructs.Song       `json:"song_data"`
}

type AddSegmentsResp struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

type GetAllSongsResp struct {
	Error string               `json:"error"`
	Songs []globalStructs.Song `json:"songs"`
}

type GetSegmentReq struct {
	ID string `json:"id"`
}

type GetSegmentResp struct {
	Segment globalStructs.SongData `json:"segment"`
	Error   string                 `json:"error"`
}

type NewUserResp struct {
	OK bool `json:"ok"`
	Error string `json:"error"`
}

type GetUserReq struct {
	ID string `json:"id"`
}

type GetUserResp struct {
	User globalStructs.User `json:"user"`
	Error string `json:"error"`
}

type NewPlaylistReq struct {
	UserID string `json:"user_id"`
	PlaylistName string `json:"playlist_name"`
	Description string `json:"description"`
	Shared bool `json:"shared"`
}

type NewPlaylistResp struct {
	Error string `json:"error"`
	OK bool `json:"ok"`
}

type DeleteUserPlaylistReq struct {
	UserID string `json:"user_id"`
	PlaylistID string `json:"playlist_id"`
}

type DeleteUserPlaylistResp struct {
	Error string `json:"error"`
	OK bool `json:"ok"`
}

type GetUserAllPlaylistsReq struct {
	UserID string `json:"user_id"`
}

type GetUserAllPlaylistsResp struct {
	Error string `json:"error"`
	Playlists []globalStructs.ShortPlaylist `json:"playlists"`
}

type GetPlaylistReq struct {
	UserID string `json:"user_id"`
	PlaylistID string `json:"playlist_id"`
}

type GetPlaylistResp struct {
	Error string `json:"error"`
	Playlist globalStructs.Playlist `json:"playlist"`
}

type AddSongToUserPlaylistReq struct {
	UserID string `json:"user_id"`
	PlaylistID string `json:"playlist_id"`
	SongID string `json:"song_id"`
}

type AddSongToUserPlaylistResp struct {
	Error string `json:"error"`
	OK bool `json:"ok"`
}

type RemoveSongFromUserPlaylistReq struct {
	UserID string `json:"user_id"`
	SongID string `json:"song_id"`
	PlaylistID string `json:"playlist_id"`
}

type RemoveSongFromUserPlaylistResp struct {
	Error string `json:"error"`
	OK bool `json:"ok"`
}

