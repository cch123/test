package main

type state string

const (
	first  state = "first"
	second       = "second"
	third        = "third"
)

var allowedState map[state][]state

func init() {
	allowedState = make(map[state][]state)
	allowedState[first] = []state{second, third}
	allowedState[second] = []state{third}
	allowedState[third] = []state{first}
}

func (s state) canTransit(newState state) bool {
	for _, s := range allowedState[s] {
		if s == newState {
			return true
		}

	}
	return false
}

func (s *state) transit(newState state) {
	if s.canTransit(newState) {
		s = &newState
	}
}
