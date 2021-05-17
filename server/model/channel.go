package model

type Channel struct {
	Model
	Name      string `json:"name"`
	Url       string `json:"url"`
	IconUrl   string `json:"icon_url"`
	SortOrder uint   `json:"sort_order"`
}

func GetChannels() (channels []Channel, err error) {
	channels = make([]Channel, 0)
	err = db.Model(Channel{}).Find(&channels).Error
	return
}
