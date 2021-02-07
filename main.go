package main

import (
	"fmt"
	viapp "hexagonalGo/video/application"
	vidomain "hexagonalGo/video/domain"
	memoryRep "hexagonalGo/video/infraestructure/inmemorydb"
)

var video vidomain.Video

func main() {
	vid, err := vidomain.NewVideo("ddd", "Primer video")

	if err != nil {
		fmt.Printf("Error creando instancia del video, %v \n", err)
		return
	}

	mmRepo := memoryRep.NewMemoryRepo()

	videoCreator := viapp.NewVideoCreator(mmRepo)

	if err = videoCreator.SaveVideo(*vid); err != nil {
		fmt.Printf("Error guardando el video, %v \n", err)
		return
	}

	videoCreator.SaveVideo(*vid)

	fmt.Printf("%d videos guardados \n", mmRepo.CountVideos())
}
