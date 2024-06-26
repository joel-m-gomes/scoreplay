package thirdparty

import (
	"scoreplay/internal/dto/thirdparty"
)

type TheSportsDBService interface {
	SearchTeam(teamName string) (*thirdparty.TheSportsDBSearchTeamDto, error)
	SearchPlayers(teamName string) ([]thirdparty.TheSportsDBSearchPlayerDto, error)
}

type ThirdPartyService struct {
	theSportsDBService TheSportsDBService
}
