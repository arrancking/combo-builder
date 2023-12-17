package api

type card struct {
	Archetype string `json:"archetype,omitempty"`
	Atk       int    `json:"atk"`
	Attribute string `json:"attribute"`
	Def       int    `json:"def"`
	Desc      string `json:"desc"`
	Level     int    `json:"level"`
	Name      string `json:"name"`
	Race      string `json:"race"`
	Type      string `json:"type"`
}

type getter interface {
	GetCard(string) (*card, error)
}

func Init() getter {
	return ygoProAPI{
		URL: "https://db.ygoprodeck.com/api/v7/cardinfo.php",
	}
}
