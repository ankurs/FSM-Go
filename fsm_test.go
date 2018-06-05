package FSMGo

/*

const (
	state1          = "state1"
	state2          = "state2"
	state_not_added = "state_not_added"
)

type entry struct {
	name    string
	enter   bool
	exit    bool
	execute bool
}

func (e *entry) GetName() string {
	return e.name
}

func (e *entry) Enter(f *FSM) {
	e.enter = true
}

func (e *entry) Exit(f *FSM) {
	e.exit = true
}

func (e *entry) Execute(f *FSM) interface{} {
	e.execute = true
	return nil
}

func setupFSMForTest() *FSM {
	e := entry{}
	e.name = state1
	f := New()
	f.AddState(&e)
	return f
}

func TestNew(t *testing.T) {
	f := New()
	assert.NotNil(t, f.states, "States should not be nil")
	assert.Nil(t, f.current, "Current state should be nil")
}

func TestAddNewState(t *testing.T) {
	f := setupFSMForTest()
	assert.Len(t, f.states, 1, "Should have only 1 state")
	assert.NotNil(t, f.states[state1])
}

func TestGetAllStates(t *testing.T) {
	f := setupFSMForTest()
	e := entry{}
	e.name = state2
	f.AddState(&e)
	states := f.GetAllStates()
	assert.Len(t, states, 2, "Should have 2 states")
	assert.Contains(t, states, state1)
	assert.Contains(t, states, state2)
}

func TestContainsState(t *testing.T) {
	f := setupFSMForTest()
	assert.True(t, f.ContainsState(state1))
	assert.False(t, f.ContainsState(state2))
	e := entry{}
	e.name = state2
	f.AddState(&e)
	assert.True(t, f.ContainsState(state1))
	assert.True(t, f.ContainsState(state2))
}

func TestMoveToState(t *testing.T) {
	f := setupFSMForTest()
	ent := entry{}
	ent.name = state2
	f.AddState(&ent)

	// move to state1
	err := f.MoveToState(state1)
	assert.NoError(t, err)
	e := f.states[state1]
	en := e.(*entry)
	// check if entry was called
	assert.NotNil(t, en, "FSMEntry is not of added type entry")
	assert.True(t, en.enter, "Enter method not called")

	// move to state2
	err = f.MoveToState(state2)
	assert.NoError(t, err)
	e = f.states[state1]
	en = e.(*entry)
	// check if exit was called for state1
	assert.NotNil(t, en, "FSMEntry is not of added type entry")
	assert.True(t, en.exit, "Exit method not called")
	e = f.states[state2]
	en = e.(*entry)
	//check if entry was called state2
	assert.NotNil(t, en, "FSMEntry is not of added type entry")
	assert.True(t, en.enter, "Enter method not called")

	err = f.MoveToState(state_not_added)
	assert.Equal(t, err, StateNotFoundError, "non existent state should return error")
}

func TestExecute(t *testing.T) {
	f := setupFSMForTest()
	f.MoveToState(state1)
	f.Execute()
	e := f.states[state1]
	en := e.(*entry)
	assert.NotNil(t, en, "FSMEntry is not of added type entry")
	assert.True(t, en.execute, "Execute method not called")
}
*/
