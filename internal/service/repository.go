package service

import (
	"fmt"

	"github.com/PrakharSrivastav/track-service-grpc/internal/model"
	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) get(id string) (*model.Track, error) {
	var track model.Track
	if err := r.db.Get(&track, "Select * from tracks where id = $1", id); err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *repository) setupDatabase() error {

	fmt.Println("Starting Schema")
	r.db.MustExec(schema)
	fmt.Println("After Schema")
	var tracks []model.Track
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_1", AlbumID: "album_id_1", ArtistID: "artist_id_1", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_2", AlbumID: "album_id_1", ArtistID: "artist_id_1", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_3", AlbumID: "album_id_1", ArtistID: "artist_id_1", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_4", AlbumID: "album_id_2", ArtistID: "artist_id_2", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_5", AlbumID: "album_id_2", ArtistID: "artist_id_2", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_6", AlbumID: "album_id_2", ArtistID: "artist_id_1", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_7", AlbumID: "album_id_2", ArtistID: "artist_id_1", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_8", AlbumID: "album_id_2", ArtistID: "artist_id_2", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_9", AlbumID: "album_id_3", ArtistID: "artist_id_2", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_10", AlbumID: "album_id_3", ArtistID: "artist_id_2", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_11", AlbumID: "album_id_3", ArtistID: "artist_id_3", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_12", AlbumID: "album_id_3", ArtistID: "artist_id_3", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_13", AlbumID: "album_id_4", ArtistID: "artist_id_3", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_14", AlbumID: "album_id_4", ArtistID: "artist_id_3", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_15", AlbumID: "album_id_5", ArtistID: "artist_id_4", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_16", AlbumID: "album_id_6", ArtistID: "artist_id_4", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_17", AlbumID: "album_id_6", ArtistID: "artist_id_5", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_18", AlbumID: "album_id_6", ArtistID: "artist_id_5", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_19", AlbumID: "album_id_7", ArtistID: "artist_id_5", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_20", AlbumID: "album_id_8", ArtistID: "artist_id_6", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_21", AlbumID: "album_id_9", ArtistID: "artist_id_6", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_22", AlbumID: "album_id_10", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_23", AlbumID: "album_id_10", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_24", AlbumID: "album_id_10", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_25", AlbumID: "album_id_11", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_26", AlbumID: "album_id_11", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_27", AlbumID: "album_id_11", ArtistID: "artist_id_7", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_28", AlbumID: "album_id_11", ArtistID: "artist_id_8", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_29", AlbumID: "album_id_11", ArtistID: "artist_id_8", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_30", AlbumID: "album_id_11", ArtistID: "artist_id_8", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_31", AlbumID: "album_id_12", ArtistID: "artist_id_8", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_32", AlbumID: "album_id_12", ArtistID: "artist_id_8", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_33", AlbumID: "album_id_11", ArtistID: "artist_id_9", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_34", AlbumID: "album_id_11", ArtistID: "artist_id_9", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_35", AlbumID: "album_id_13", ArtistID: "artist_id_9", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_36", AlbumID: "album_id_12", ArtistID: "artist_id_10", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tracks = append(tracks, model.Track{Description: gofakeit.Sentence(5), ID: "track_id_37", AlbumID: "album_id_14", ArtistID: "artist_id_10", Name: gofakeit.Name(), Duration: gofakeit.Float64()})
	tx := r.db.MustBegin()
	for _, u := range tracks {
		fmt.Printf("%#v\n", &u)
		_, err := tx.NamedExec("INSERT into tracks (id,name,description,artistId,albumId,duration) values(:id, :name, :description, :artistId, :albumId, :duration)", u)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	tx.Commit()
	return nil
}

var schema = `Drop Table if exists tracks;
CREATE TABLE tracks (
    id text,
    name text,
	description text,
	albumId text,
	artistId text,
	duration float
);`
