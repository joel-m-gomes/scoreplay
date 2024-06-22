package thirdparty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scoreplay/internal/dto/thirdparty"
	"scoreplay/internal/exception"
)

type DefaultTheSportsBDService struct {
	apiURL     string
	apiVersion string
	apiKey     string
}

func NewDefaultTheSportsBDService() *DefaultTheSportsBDService {
	return &DefaultTheSportsBDService{
		apiURL:     os.Getenv("THESPORTSDB_API_URL"),
		apiVersion: os.Getenv("THESPORTSDB_API_VERSION"),
		apiKey:     os.Getenv("THESPORTSDB_API_KEY"),
	}
}

func (s *DefaultTheSportsBDService) SearchTeam(teamName string) (*thirdparty.TheSportsDBSearchTeamDto, error) {
	url := fmt.Sprintf("%s/%s/json/%s/searchteams.php?t=%s", s.apiURL, s.apiVersion, s.apiKey, teamName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, exception.NewThirdPartyException("thesportsdb", http.StatusInternalServerError, "failed to call API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, exception.NewThirdPartyException("thesportsdb", resp.StatusCode, "failed to search team")
	}

	var result thirdparty.TheSportsDBSearchTeamResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, exception.NewThirdPartyException("thesportsdb", http.StatusInternalServerError, "failed to parse response")
	}

	if len(result.Teams) > 0 {
		return &result.Teams[0], nil
	}

	return nil, exception.NewThirdPartyException("thesportsdb", http.StatusNotFound, "team not found")
}

func (s *DefaultTheSportsBDService) SearchPlayers(teamName string) ([]thirdparty.TheSportsDBSearchPlayerDto, error) {
	url := fmt.Sprintf("%s/%s/json/%s/searchplayers.php?t=%s", s.apiURL, s.apiVersion, s.apiKey, teamName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, exception.NewThirdPartyException("thesportsdb", http.StatusInternalServerError, "failed to call API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, exception.NewThirdPartyException("thesportsdb", resp.StatusCode, "failed to search team "+teamName+" players")
	}

	var result thirdparty.TheSportsDBSearchPlayerResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, exception.NewThirdPartyException("thesportsdb", http.StatusInternalServerError, "failed to parse response")
	}

	return result.Players, nil
}
