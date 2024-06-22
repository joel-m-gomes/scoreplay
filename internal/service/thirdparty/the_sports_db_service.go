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
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, exception.NewThirdPartyException("thesportsdb", resp.StatusCode, "failed to search team")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if teams, ok := result["teams"].([]interface{}); ok && len(teams) > 0 {
		team := teams[0].(map[string]interface{})
		teamDTO := &thirdparty.TheSportsDBSearchTeamDto{
			TeamLogo: team["strLogo"].(string),
		}
		return teamDTO, nil
	}

	return nil, exception.NewThirdPartyException("thesportsdb", resp.StatusCode, "team not found")
}

func (s *DefaultTheSportsBDService) SearchPlayers(teamName string) ([]thirdparty.TheSportsDBSearchTeamPlayersDto, error) {
	url := fmt.Sprintf("%s/%s/json/%s/searchplayers.php?t=%s", s.apiURL, s.apiVersion, s.apiKey, teamName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, exception.NewThirdPartyException("thesportsdb", resp.StatusCode, "failed to search team "+teamName+" players")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var playersDTO []thirdparty.TheSportsDBSearchTeamPlayersDto
	if players, ok := result["player"].([]interface{}); ok {
		for _, player := range players {
			playerMap := player.(map[string]interface{})
			playerDTO := thirdparty.TheSportsDBSearchTeamPlayersDto{
				PlayerName: playerMap["strPlayer"].(string),
				Thumbnail:  playerMap["strThumb"].(string),
			}
			playersDTO = append(playersDTO, playerDTO)
		}
	}

	return playersDTO, nil
}
