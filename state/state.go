package state

import "github.com/tomasbasham/gonads"

// State is a monad that models computations that depend on some global state.
type State[S, A any] struct {
	g func(S) (A, S)
}

// Run runs a state and returns its value.
func (s *State[S, A]) Run(state S) (A, S) {
	return s.g(state)
}

// Pure returns a State[S, A] containing the given value.
func Pure[S, A any](a A) State[S, A] {
	return State[S, A]{func(s S) (A, S) { return a, s }}
}

// Map maps a function over a state.
func Map[S, A, B any](s State[S, A], f gonads.Func[A, B]) State[S, B] {
	return State[S, B]{func(state S) (B, S) {
		a, newState := s.g(state)
		return f(a), newState
	}}
}

// Apply applies a state function to a state value.
func Apply[S, A, B any](s State[S, A], f State[S, gonads.Func[A, B]]) State[S, B] {
	return State[S, B]{func(state S) (B, S) {
		a, newState := s.g(state)
		g, _ := f.g(newState)
		return g(a), newState
	}}
}

// FlatMap maps a function over a state.
func FlatMap[S, A, B any](s State[S, A], f gonads.Func[A, State[S, B]]) State[S, B] {
	return State[S, B]{func(state S) (B, S) {
		a, newState := s.g(state)
		return f(a).g(newState)
	}}
}

// Get gets the state as a value of itself.
func Get[S any]() State[S, S] {
	return State[S, S]{func(state S) (S, S) {
		return state, state
	}}
}

// Put sets the state to the given value.
func Put[S any](state S) State[S, gonads.Unit] {
	return State[S, gonads.Unit]{func(_ S) (gonads.Unit, S) { return gonads.UnitValue, state }}
}

// RunState runs a state and returns its value.
func RunState[S, A any](s State[S, A], state S) (A, S) {
	return s.Run(state)
}
