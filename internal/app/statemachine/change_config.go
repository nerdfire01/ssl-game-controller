package statemachine

import (
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	"log"
)

func (s *StateMachine) processChangeUpdateConfig(newState *state.State, change *UpdateConfig) (changes []*Change) {
	if change.Division != nil {
		log.Printf("Change division to %v", *change.Division)
		newState.Division = change.Division
		s.updateMaxBots(newState)
		s.Geometry = s.gameConfig.DefaultGeometry[change.Division.Div()]
		log.Printf("Updated geometry to %+v", s.Geometry)
	}
	if change.FirstKickoffTeam != nil {
		log.Printf("Change first kickoff team to %v", *change.FirstKickoffTeam)
		newState.FirstKickoffTeam = change.FirstKickoffTeam
		if *newState.Stage == state.Referee_NORMAL_FIRST_HALF_PRE {
			newState.NextCommand = state.NewCommand(state.Command_KICKOFF, *newState.FirstKickoffTeam)
		}
	}
	if change.AutoContinue != nil {
		log.Printf("Change auto continue to %v", *change.AutoContinue)
		newState.AutoContinue = change.AutoContinue
	}
	if change.MatchType != nil {
		log.Printf("Change match type to %s", change.MatchType.String())
		newState.MatchType = change.MatchType
	}
	return
}
