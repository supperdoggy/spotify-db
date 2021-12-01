package db

import (
	globalStructs "github.com/supperdoggy/spotify-web-project/spotify-globalStructs"
	"gopkg.in/mgo.v2"
)

type obj map[string]interface{}

type IDB interface {
	GetAllSongs() (result []globalStructs.Song, err error)
	GetSegment(id string) (result globalStructs.SongData, err error)
	InsertSegment(ts ...globalStructs.SongData) error
	InsertSong(s globalStructs.Song) error
	GetUserByID(id string) (resp globalStructs.User, err error)
	NewUser(u globalStructs.User) error
}

type DB struct {
	Session *mgo.Session

	SegmentsCollection *mgo.Collection
	SongsCollection    *mgo.Collection
	UsersCollection    *mgo.Collection
}

const GetAllSongsLimit = 1000

func NewDB(dbname string) (IDB, error) {
	session, err := mgo.Dial("")
	if err != nil {
		return nil, err
	}
	return &DB{
		Session:            session,
		SegmentsCollection: session.DB(dbname).C("segments"),
		SongsCollection:    session.DB(dbname).C("songs"),
		UsersCollection:    session.DB(dbname).C("users"),
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

func (d *DB) NewUser(u globalStructs.User) error {
	err := d.UsersCollection.Insert(u)
	return err
}

func (d *DB) GetUserByID(id string) (resp globalStructs.User, err error) {
	err = d.UsersCollection.Find(obj{"_id": id}).One(&resp)
	return
}
