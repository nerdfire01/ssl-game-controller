package engine

import (
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/config"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/statemachine"
	"github.com/go-test/deep"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Engine(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("Could not create a temporary directory: ", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatalf("Could not cleanup tmpDir %v: %v", tmpDir, err)
		}
	}()

	gameConfig := config.DefaultControllerConfig().Game
	gameConfig.StateStoreFile = tmpDir + "/store.json.stream"
	engine := NewEngine(gameConfig, config.Engine{ConfigFilename: tmpDir + "/engine.yaml"})
	hook := make(chan HookOut)
	engine.RegisterHook("engineTest", hook)
	if err := engine.Start(); err != nil {
		t.Fatal("Could not start engine")
	}
	wantNewState := state.NewState()
	proto.Merge(wantNewState, engine.currentState)
	engine.Enqueue(&statemachine.Change{
		Change: &statemachine.Change_NewCommand{
			NewCommand: &statemachine.NewCommand{
				Command: state.NewCommandNeutral(state.Command_HALT),
			},
		},
	})
	// wait for the changes to be processed
	<-hook
	engine.UnregisterHook("engineTest")
	engine.Stop()

	wantNewState.Command = state.NewCommandNeutral(state.Command_HALT)

	gotNewState := engine.currentState
	diffs := deep.Equal(gotNewState, wantNewState)
	if len(diffs) > 0 {
		t.Error("Entries differ: ", diffs)
	}
}
