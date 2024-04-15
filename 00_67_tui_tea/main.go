package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type CommonID struct {
	id int
}

type ControlInput struct {
	CommonID
	textInput textinput.Model
}

type ControlTextArea struct {
	CommonID
	textArea textarea.Model
}

type model struct {
	focusIndex   int
	ControlInput []ControlInput
	ControlArea  []ControlTextArea
	kvInput      map[int]any

	err error
}

type errMsg error

func initModel() model {
	textAreaModel := textarea.New()
	textAreaModel.Placeholder = "Write something..."
	textAreaModel.ShowLineNumbers = true

	inputs := []ControlInput{
		{
			CommonID: CommonID{
				id: 0,
			},
			textInput: textinput.New(),
		},
	}

	textInputModel := textinput.New()
	textInputModel.Placeholder = "Current time is..."
	textInputModel.Focus()

	textAreas := []ControlTextArea{
		{
			CommonID: CommonID{
				id: 1,
			},
			textArea: textAreaModel,
		},
	}

	return model{
		focusIndex:   0,
		ControlInput: inputs,
		ControlArea:  textAreas,
		kvInput: map[int]any{
			0: textInputModel,
			1: textAreaModel,
		},
		err: nil,
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	// m.focusedControl = &m.ControlInput[0]
	return m.ControlInput[0].textInput.Focus()
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			switch m.kvInput[m.focusIndex].(type) {
			case textinput.Model:
				v := m.kvInput[m.focusIndex].(textinput.Model)
				v.Blur()
			case textarea.Model:
				v := m.kvInput[m.focusIndex].(textarea.Model)
				v.Blur()
			}

		case "tab":
			switch m.kvInput[m.focusIndex].(type) {
			case textinput.Model:
				v := m.kvInput[m.focusIndex].(textinput.Model)
				v.Blur()
				m.focusIndex += 1
				if m.focusIndex > len(m.kvInput) {
					m.focusIndex = 0
				}
				fv := m.kvInput[m.focusIndex].(textarea.Model)
				cmd = fv.Focus()
				cmds = append(cmds, cmd)
			case textarea.Model:
				v := m.kvInput[m.focusIndex].(textarea.Model)
				v.Blur()
				m.focusIndex += 1
				if m.focusIndex > len(m.kvInput) {
					m.focusIndex = 0
				}
				fv := m.kvInput[m.focusIndex].(textinput.Model)
				cmd = fv.Focus()
				cmds = append(cmds, cmd)
			}

		default:

		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	// m.textArea, cmd = m.textArea.Update(msg)
	switch m.kvInput[m.focusIndex].(type) {
	case textinput.Model:
		v := m.kvInput[m.focusIndex].(textinput.Model)
		v, cmd = v.Update(msg)
		m.kvInput[m.focusIndex] = v
	case textarea.Model:
		v := m.kvInput[m.focusIndex].(textarea.Model)
		v, cmd = v.Update(msg)
		m.kvInput[m.focusIndex] = v
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	s := fmt.Sprintf("%s \n %s",
		m.kvInput[0].(textinput.Model).View(),
		m.kvInput[1].(textarea.Model).View())
	return s
}
