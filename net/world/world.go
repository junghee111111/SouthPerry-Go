/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package world

type World struct {
	Name         string
	Flag         uint32
	EventMessage string
}

func NewWorld(name string, flag uint32, eventMessage string) *World {
	return &World{
		Name:         name,
		Flag:         flag,
		EventMessage: eventMessage,
	}
}
