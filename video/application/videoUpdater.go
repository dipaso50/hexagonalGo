package application

import (
	"fmt"
	vidomain "hexagonalGo/video/domain"
)

type videoUpdater struct {
	repo vidomain.VideoRepo
}

//NewVideoUpdater creates a new video updater
func NewVideoUpdater(rep vidomain.VideoRepo) *videoUpdater {
	return &videoUpdater{
		repo: rep,
	}
}

func (viup *videoUpdater) UpdateVideo(vid vidomain.Video) error {
	vi, err := viup.repo.Select(vid.ID.Value)

	if err != nil {
		return fmt.Errorf("Error buscando el video a actualizar, %v", err)
	}

	if vi == nil {
		return fmt.Errorf("No hemos encontrado el video para hacer la actualización")
	}

	viup.repo.Update(vid)

	return nil
}
