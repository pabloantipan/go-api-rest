package domain

type Profile struct {
	Stats        interface{} `json:"stats"`
	Achievements interface{} `json:"achievements"`
	Team         interface{} `json:"team"`
}
