package FSMGo

type FSM interface {
	GetCurrentState() State
	AddStates(entries ...State)
	GetAllStates() []string
	ContainsState(state string) bool
	MoveToState(state string) error
	Execute() interface{}
	ExecuteTillFinish()
	Finish()
	IsFinished() bool
	// get and set data
}

type fsm struct {
	states  map[string]State
	bag     interface{}
	current State
	finish  bool
}

func (sm *fsm) GetCurrentState() State {
	return sm.current
}

type State interface {
	GetName() string
	Enter(fsm FSM)
	Execute(fsm FSM) interface{}
	Exit(fsm FSM)
}
