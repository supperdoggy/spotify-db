package db

import (
	"errors"
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"github.com/u2takey/go-utils/rand"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

type obj map[string]interface{}

type IDB interface {
	GetAllSongs() (result []globalStructs.Song, err error)
	GetSegment(id string) (result globalStructs.SongData, err error)
	InsertSegment(ts ...globalStructs.SongData) error
	InsertSong(s globalStructs.Song) error
	GetSongByID(id string) (s globalStructs.Song, err error)
	GetUserByID(id string) (resp globalStructs.User, err error)
	NewUser(u globalStructs.User) error
	NewPlaylist(p globalStructs.Playlist) error
	DeleteUserPlaylist(id, owner string) error
	DeletePlaylistByID(id string) error
	GetPlaylistByID(id string) (p globalStructs.Playlist, err error)
	AddSongsToUserPlaylist(id, owner string, song globalStructs.Song) error
	AddSongsToPlaylist(id string, song globalStructs.Song) error
	RemoveSongFromUserPlaylist(id, owner, songID string) error
	RemoveSongFromPlaylist(id, songID string) error
	GetAllUserPlaylists(owner string) (p []globalStructs.ShortPlaylist, err error)
}

type DB struct {
	Logger  *zap.Logger
	Session *mgo.Session

	SegmentsCollection *mgo.Collection
	SongsCollection    *mgo.Collection
	UsersCollection    *mgo.Collection
	PlaylistCollection *mgo.Collection
}

const GetAllSongsLimit = 1000

func NewDB(dbname string, logger *zap.Logger) (IDB, error) {
	session, err := mgo.Dial("")
	if err != nil {
		return nil, err
	}
	return &DB{
		Logger:             logger,
		Session:            session,
		SegmentsCollection: session.DB(dbname).C("segments"),
		SongsCollection:    session.DB(dbname).C("songs"),
		UsersCollection:    session.DB(dbname).C("users"),
		PlaylistCollection: session.DB(dbname).C("playlists"),
	}, nil
}

// GetAllSongs - limit for 1000
func (d *DB) GetAllSongs() (result []globalStructs.Song, err error) {
	err = d.SongsCollection.Find(obj{}).Limit(GetAllSongsLimit).All(&result)
	return
}

func (d *DB) GetSegment(id string) (result globalStructs.SongData, err error) {
	err = d.SegmentsCollection.Find(obj{"_id": id}).One(&result)
	return
}

func (d *DB) InsertSegment(ts ...globalStructs.SongData) error {
	for _, v := range ts {
		err := d.SegmentsCollection.Insert(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) InsertSong(s globalStructs.Song) error {
	err := d.SongsCollection.Insert(s)
	return err
}

func (d *DB) GetSongByID(id string) (s globalStructs.Song, err error) {
	err = d.SongsCollection.Find(obj{"_id": id}).One(&s)
	return
}

//func (d *DB) DeleteSongByID(id string) (s globalStructs.Song)

func (d *DB) NewUser(u globalStructs.User) error {
	err := d.UsersCollection.Insert(u)
	return err
}

func (d *DB) GetUserByID(id string) (resp globalStructs.User, err error) {
	err = d.UsersCollection.Find(obj{"_id": id}).One(&resp)
	return
}

func (d *DB) NewPlaylist(p globalStructs.Playlist) error {
	p.ID = rand.String(24)
	err := d.PlaylistCollection.Insert(p)
	if mgo.IsDup(err) {
		return d.NewPlaylist(p)
	}
	return err
}

func (d *DB) DeleteUserPlaylist(id, owner string) error {
	if id == "" || owner == "" {
		return errors.New("id and owner must not be empty")
	}
	err := d.PlaylistCollection.Remove(obj{"_id": id, "owner": owner})
	return err
}

func (d *DB) DeletePlaylistByID(id string) error {
	if id == "" {
		return errors.New("id must not be empty")
	}
	err := d.PlaylistCollection.Remove(obj{"_id": id})
	return err
}

func (d *DB) GetAllUserPlaylists(owner string) (p []globalStructs.ShortPlaylist, err error) {
	err = d.PlaylistCollection.Find(obj{"owner_id": owner}).All(&p)
	return
}

func (d *DB) GetPlaylistByID(id string) (p globalStructs.Playlist, err error) {
	if id == "" {
		return p, errors.New("id must not be empty")
	}
	err = d.PlaylistCollection.Find(obj{"_id": id}).One(&p)
	return
}

// AddSongsToUserPlaylist - adds songs to playlist linked with user
func (d *DB) AddSongsToUserPlaylist(id, owner string, song globalStructs.Song) error {
	if id == "" || owner == "" || song.ID == "" {
		return errors.New("id and owner must not be empty")
	}

	err := d.PlaylistCollection.Update(obj{
		"_id":      id,
		"owner_id": owner,
	}, obj{
		"$push": obj{"songs": song},
	})

	return err
}

// AddSongsToPlaylist - adds songs to playlist
func (d *DB) AddSongsToPlaylist(id string, song globalStructs.Song) error {
	if id == "" || song.ID == "" {
		return errors.New("id and owner must not be empty")
	}

	err := d.PlaylistCollection.Update(obj{
		"_id": id,
	}, obj{
		"$push": obj{"songs": song},
	})

	return err
}

// RemoveSongFromUserPlaylist - sends req to mongo to find and remove song by id from songs slice
func (d *DB) RemoveSongFromUserPlaylist(id, owner, songID string) error {
	if id == "" || owner == "" || songID == "" {
		return errors.New("id and owner must not be empty")
	}

	return d.PlaylistCollection.Update(obj{
		"_id":      id,
		"owner_id": owner,
		"songs": obj{
			"_id": songID,
		},
	}, obj{
		"$pull": obj{
			"songs": obj{"_id": songID},
		},
	})
}

func (d *DB) RemoveSongFromPlaylist(id, songID string) error {
	if id == "" || songID == "" {
		return errors.New("id and songID must not be empty")
	}

	return d.PlaylistCollection.Update(obj{
		"_id": id,
		"songs": obj{
			"_id": songID,
		},
	}, obj{
		"$pull": obj{
			"songs": obj{"_id": songID},
		},
	})
}
