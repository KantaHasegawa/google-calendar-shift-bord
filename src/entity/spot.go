package entity

type TSpot struct {
    User      string    `json:"User"`
    StartWork   string `json:"StartWork"`
		SpotId			string `json:"SpotId"`
}

type SpotInteractorInteface interface {
    DetailSpot(string, string, string, string) (TSpot, error)
}
