package inmemorydb

import vidomain "hexagonalGo/video/domain"

type MemoryRepo struct {
	allVideos map[string]*vidomain.Video
}

//NewMemoryRepo creates a new repo
func NewMemoryRepo() *MemoryRepo {
	mr := MemoryRepo{}
	mr.allVideos = make(map[string]*vidomain.Video)
	return &mr
}

//Save save video in current implementation
func (merep *MemoryRepo) Save(video vidomain.Video) error {
	merep.allVideos[video.ID.Value] = &video
	return nil
}

//CountVideos returns total count of videos registered
func (merep *MemoryRepo) CountVideos() int {
	return len(merep.allVideos)
}

//Select get a video by its uuid
func (merep *MemoryRepo) Select(uuid string) (*vidomain.Video, error) {
	return merep.allVideos[uuid], nil
}

//Update updates video info
func (merep *MemoryRepo) Update(video vidomain.Video) error {
	merep.allVideos[video.ID.Value] = &video
}
