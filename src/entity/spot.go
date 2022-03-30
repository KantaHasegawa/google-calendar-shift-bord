package entity

type TSpot struct {
	User      string `json:"User"`
	StartWork string `json:"StartWork"`
	SpotId    string `json:"SpotId"`
	SpotData  TSpotData
}

type TSpotData struct {
	Name      string `json:"Name"`
	Salary    int    `json:"Salaly"`
	CutOffDay string `json:"CutOffDay"`
	PayDay    string `json:"PayDay"`
}

type SpotInteractorInteface interface {
	ListSpot(table string, user string) ([]TSpot, error)
	DetailSpot(string, string, string, string) (TSpot, error)
	CreateSpot(string, string, string, int, string, string) (string, error)
}
