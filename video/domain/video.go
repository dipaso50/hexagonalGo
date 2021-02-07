package domain

//Video video
type Video struct {
	ID   videoID
	Name videoName
}

//VideoRepo for database operations
type VideoRepo interface {
	Save(Video) error
}

//NewVideo creates a video object
func NewVideo(uuid, name string) (*Video, error) {

	id, err := newVideoID(uuid)

	if err != nil {
		return nil, err
	}

	nm, err := newVideoName(name)

	if err != nil {
		return nil, err
	}

	return &Video{
		ID:   *id,
		Name: *nm,
	}, nil
}
