package FSMGo

type FSM struct {
	states  map[string]FSMEntry
	Bag     interface{}
	current FSMEntry
	finish  bool
}

type FSMEntry interface {
	GetName() string
	Enter(fsm *FSM)
	Execute(fsm *FSM) interface{}
	Exit(fsm *FSM)
}
