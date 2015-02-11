package FSMGo

import (
	"errors"
)

var (
	StateNotFoundError = errors.New("State Not Found")
)

func New() *FSM {
	fsm := new(FSM)
	fsm.states = make(map[string]FSMEntry)
	fsm.finish = false
	fsm.current = nil
	return fsm
}

func (fsm *FSM) AddState(entry FSMEntry) {
	if entry != nil {
		fsm.states[entry.GetName()] = entry
	}
}

func (fsm *FSM) AddStates(entries ...FSMEntry) {
	for pos := range entries {
		fsm.AddState(entries[pos])
	}
}

func (fsm *FSM) GetAllStates() []string {
	states := make([]string, 0)
	for state := range fsm.states {
		states = append(states, state)
	}
	return states
}

func (fsm *FSM) ContainsState(state string) bool {
	_, ok := fsm.states[state]
	return ok
}

func (fsm *FSM) MoveToState(state string) error {
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

func (fsm *FSM) Execute() interface{} {
	if fsm.current == nil {
		return nil
	}
	return fsm.current.Execute(fsm)
}

func (fsm *FSM) ExecuteTillFinish() {
	if fsm.current != nil {
		for !fsm.IsFinished() {
			fsm.Execute()
		}
	}
}

func (fsm *FSM) Finish() {
	fsm.finish = true
}

func (fsm *FSM) IsFinished() bool {
	return fsm.finish
}
