package entity



type TSpot struct {
    User      string    `json:"User"`
    StartWork   string `json:"StartWork"`
		SpotId			string `json:"SpotId"`
		SpotData struct {
			Name string `json:"Name"`
			Salary int `json:"Salaly"`
			CutOffDay string `json:"CutOffDay"`
			PayDay string `json:"PayDay"`
	}
}

type SpotInteractorInteface interface {
    DetailSpot(string, string, string, string) (TSpot, error)
}
