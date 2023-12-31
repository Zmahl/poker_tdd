package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	i.store[name]++
	i.mu.Unlock()
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}
