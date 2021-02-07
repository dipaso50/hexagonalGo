package domain

type videoName struct {
	Value string
}

func newVideoName(uuid string) (*videoName, error) {

	vid := &videoName{
		Value: uuid,
	}

	return vid, nil
}
