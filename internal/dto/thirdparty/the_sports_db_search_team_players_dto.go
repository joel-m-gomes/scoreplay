package thirdparty

// TheSportsDBSearchPlayerDto represents the player data from TheSportsDB API.
type TheSportsDBSearchPlayerDto struct {
	PlayerName string `json:"strPlayer"`
	Thumbnail  string `json:"strThumb"`
}

// TheSportsDBSearchPlayerResponseDto represents the response structure for searching players.
type TheSportsDBSearchPlayerResponseDto struct {
	Players []TheSportsDBSearchPlayerDto `json:"player"`
}
