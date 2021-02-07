package application

import vidomain "hexagonalGo/video/domain"

type videoCreator struct {
	repo vidomain.VideoRepo
}

//NewVideoCreator creates a instance of this app services
func NewVideoCreator(rep vidomain.VideoRepo) *videoCreator {
	return &videoCreator{
		repo: rep,
	}
}

func (vicre *videoCreator) SaveVideo(vid vidomain.Video) error {
	return vicre.repo.Save(vid)
}
