package test

import (
	"fmt"
	"go_generics/helper"
	"testing"
)

type Employee interface {
	GetMessage() string
}

func PrintMessage[T Employee](params T) {
	fmt.Println(params.GetMessage())
}

type Manager struct {
	Name  string
	Money int
}

func NewManager(name string, money int) Employee {
	return &Manager{Name: name, Money: money}
}

func (m *Manager) GetMessage() string {
	return fmt.Sprintf("Hello My Name Is %s ,%s", m.Name, helper.If(m.IsRich(), "IM RICH", "IM NOT RICH"))
}

func (m *Manager) IsRich() bool {
	return m.Money > 1000
}

func TestInheritance(t *testing.T) {
	manager := NewManager("Ten Hag", 99)
	PrintMessage[Employee](manager)
}
