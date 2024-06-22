package thirdparty

// TheSportsDBSearchTeamDto represents the team data from TheSportsDB API.
type TheSportsDBSearchTeamDto struct {
	TeamLogo string `json:"strTeamLogo"`
}

// TheSportsDBSearchTeamResponseDto represents the response structure for searching teams.
type TheSportsDBSearchTeamResponseDto struct {
	Teams []TheSportsDBSearchTeamDto `json:"teams"`
}
