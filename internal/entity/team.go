package entity

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	Players []int  `json:"players"`
}
