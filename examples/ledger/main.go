package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tomasbasham/gonads"
	"github.com/tomasbasham/gonads/state"
)

type (
	ledgerState = state.State[ledger, int]

	ledger map[string]entry
	entry  struct {
		amount       int
		balanceAfter int
	}
)

func main() {
	s := state.Pure[ledger](0)
	s, e1 := deposit(s, 75)
	s, e2 := withdraw(s, 50)

	l := ledger{} // Initial state.

	b, l1 := balance(s, l)
	fmt.Printf("balance: %d\n", b)

	b1, _ := balanceAt(l1, e1)
	fmt.Printf("balance at %s: %+v\n", e1, b1)

	b2, _ := balanceAt(l1, e2)
	fmt.Printf("balance at %s: %+v\n", e2, b2)
}

func deposit(s ledgerState, amount int) (ledgerState, string) {
	return update(s, amount)
}

func withdraw(s ledgerState, amount int) (ledgerState, string) {
	return update(s, -amount)
}

func update(s ledgerState, amount int) (ledgerState, string) {
	entryID := generateEntryID()

	return state.FlatMap(s, func(balance int) ledgerState {
		return state.FlatMap(state.Get[ledger](), func(l ledger) ledgerState {
			l[entryID] = entry{amount, balance + amount}
			return state.FlatMap(state.Put(l), func(_ gonads.Unit) ledgerState {
				return state.Pure[ledger](balance + amount)
			})
		})
	}), entryID
}

func balance(s ledgerState, l ledger) (int, ledger) {
	return s.Run(l)
}

func balanceAt(l ledger, entryID string) (int, error) {
	if _, ok := l[entryID]; !ok {
		return 0, fmt.Errorf("entry %s not found", entryID)
	}
	return l[entryID].balanceAfter, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

// generateEntryID is a helper function provided to generate a random ID
// to use for an entry. You can assume that these are completely unique
func generateEntryID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("err generating entry id: %v", err)
	}
	return fmt.Sprintf("entry-%x", b[0:8])
}
