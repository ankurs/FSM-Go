package FSMGo

import (
	"errors"
)

var (
	StateNotFoundError = errors.New("State Not Found")
)

func New() FSM {
	fsm := new(fsm)
	fsm.states = make(map[string]State)
	fsm.finish = false
	fsm.current = nil
	return fsm
}

func (fsm *fsm) AddState(entry State) {
	if entry != nil {
		fsm.states[entry.GetName()] = entry
	}
}

func (fsm *fsm) AddStates(entries ...State) {
	for pos := range entries {
		fsm.AddState(entries[pos])
	}
}

func (fsm *fsm) GetAllStates() []string {
	states := make([]string, 0)
	for state := range fsm.states {
		states = append(states, state)
	}
	return states
}

func (fsm *fsm) ContainsState(state string) bool {
	_, ok := fsm.states[state]
	return ok
}

func (fsm *fsm) MoveToState(state string) error {
	if entry, ok := fsm.states[state]; ok {
		if fsm.current != nil {
			fsm.current.Exit(fsm)
		}
		fsm.current = entry
		fsm.current.Enter(fsm)
		return nil
	} else {
		return StateNotFoundError
	}
}

func (fsm *fsm) Execute() interface{} {
	if fsm.current == nil {
		return nil
	}
	return fsm.current.Execute(fsm)
}

func (fsm *fsm) ExecuteTillFinish() {
	if fsm.current != nil {
		for !fsm.IsFinished() {
			fsm.Execute()
		}
	}
}

func (fsm *fsm) Finish() {
	fsm.finish = true
}

func (fsm *fsm) IsFinished() bool {
	return fsm.finish
}
