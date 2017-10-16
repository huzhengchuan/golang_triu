package main

import (
	"fmt"
)

var m *Manager

func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}

type Manager struct{}

func (p Manager) Manager() {
	fmt.Println("manage....")
}

func main() {
	var m *Manager
	m = GetInstance()
	m.Manager()
}
