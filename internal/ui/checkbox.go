// Package: ui
// Checkbox model para selección de opciones de configuración
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/XebecCorporation/XebecCorporation.Dots/internal/actions"
)

// ============================================
// Mensajes para Checkbox
// ============================================

// CheckboxOptionSelectedMsg cuando una opción es seleccionada
type CheckboxOptionSelectedMsg struct {
	OptionID string
	Selected bool
}

// CheckboxConfirmMsg cuando el usuario confirma la selección
type CheckboxConfirmMsg struct{}

// CheckboxCancelMsg cuando el usuario cancela
type CheckboxCancelMsg struct{}

// ============================================
// Modelo de Checkbox
// ============================================

// CheckboxOption representa una opción con checkbox
type CheckboxOption struct {
	ID          string
	Title       string
	Description string
	Checked     bool
}

// CheckboxModel modelo para manejar checkboxes
type CheckboxModel struct {
	Title       string
	Options     []CheckboxOption
	Selected    int
	ConfirmText string
	CancelText  string
	Result      chan []CheckboxOption // Canal para retornar resultado
	Confirmed   bool
	Cancelled   bool
}

// NewCheckboxModel crea un nuevo modelo de checkbox
func NewCheckboxModel(title string, options []actions.AlacrittyConfigOption) *CheckboxModel {
	checkboxOpts := make([]CheckboxOption, len(options))
	for i, opt := range options {
		checkboxOpts[i] = CheckboxOption{
			ID:          opt.ID,
			Title:       opt.Title,
			Description: opt.Description,
			Checked:     false, // Por defecto desmarcado
		}
	}

	return &CheckboxModel{
		Title:       title,
		Options:     checkboxOpts,
		Selected:    0,
		ConfirmText: "Configurar",
		CancelText:  "Cancelar",
		Result:      make(chan []CheckboxOption, 1),
		Confirmed:   false,
		Cancelled:   false,
	}
}

// Init implementación de tea.Model
func (m CheckboxModel) Init() tea.Cmd {
	return nil
}

// Update implementación de tea.Model
func (m CheckboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
			}
		case "down", "j":
			if m.Selected < len(m.Options) {
				m.Selected++
			}
		case " ":
			// Toggle checkbox (solo para opciones, no para confirmar/cancelar)
			if m.Selected < len(m.Options) {
				m.Options[m.Selected].Checked = !m.Options[m.Selected].Checked
			}
		case "enter":
			// Si estamos en la última opción, es confirmar
			if m.Selected == len(m.Options) {
				// Confirmar
				m.Confirmed = true
				m.Result <- m.Options
				return m, tea.Quit
			} else if m.Selected == len(m.Options)+1 {
				// Cancelar
				m.Cancelled = true
				m.Result <- nil
				return m, tea.Quit
			} else {
				// Seleccionar opción
				m.Options[m.Selected].Checked = !m.Options[m.Selected].Checked
			}
		case "q", "esc", "ctrl+c":
			m.Cancelled = true
			m.Result <- nil
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		// Ajustar tamaño si es necesario
	}

	return m, nil
}

// View implementación de tea.Model
func (m CheckboxModel) View() string {
	var b strings.Builder

	// Título
	b.WriteString(TitleStyle.Render(m.Title))
	b.WriteString("\n\n")

	// Opciones
	checkboxStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(BrandingConfig.Colors.White))

	checkedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(BrandingConfig.Colors.AccentGreen))

	uncheckedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(BrandingConfig.Colors.GrayLight))

	selectedStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(BrandingConfig.Colors.Primary)).
		Foreground(lipgloss.Color(BrandingConfig.Colors.White))

	for i, opt := range m.Options {
		checkbox := "[ ]"
		style := checkboxStyle

		if opt.Checked {
			checkbox = "[✓]"
			style = checkedStyle
		} else {
			checkbox = "[ ]"
			style = uncheckedStyle
		}

		// Si está seleccionado, usar estilo destacado
		if i == m.Selected {
			style = selectedStyle
			b.WriteString("► ")
		} else {
			b.WriteString("  ")
		}

		b.WriteString(style.Render(fmt.Sprintf("%s %s", checkbox, opt.Title)))
		b.WriteString("\n")
		b.WriteString(MutedTextStyle.Render(fmt.Sprintf("     %s", opt.Description)))
		b.WriteString("\n\n")
	}

	// Botones de confirmar y cancelar
	b.WriteString("\n")
	b.WriteString(MutedTextStyle.Render("─────────────────────────────"))
	b.WriteString("\n\n")

	// Opción confirmar
	if m.Selected == len(m.Options) {
		b.WriteString(selectedStyle.Render("► " + m.ConfirmText))
	} else {
		b.WriteString("  " + m.ConfirmText)
	}
	b.WriteString("  ")

	// Opción cancelar
	if m.Selected == len(m.Options)+1 {
		b.WriteString(selectedStyle.Render("► " + m.CancelText))
	} else {
		b.WriteString(MutedTextStyle.Render(m.CancelText))
	}
	b.WriteString("\n\n")

	// Footer
	b.WriteString(MutedTextStyle.Render("Espacio/Enter: marcar │ ↑↓: navegar │ q: salir"))

	return b.String()
}

// GetSelectedOptions retorna las opciones seleccionadas
func (m CheckboxModel) GetSelectedOptions() []CheckboxOption {
	var selected []CheckboxOption
	for _, opt := range m.Options {
		if opt.Checked {
			selected = append(selected, opt)
		}
	}
	return selected
}

// RunCheckboxModel ejecuta el modelo de checkbox y retorna las opciones
func RunCheckboxModel(title string, options []actions.AlacrittyConfigOption) []CheckboxOption {
	model := NewCheckboxModel(title, options)

	p := tea.NewProgram(model)
	result, err := p.Run()

	if err != nil {
		fmt.Println("Error ejecutando checkbox:", err)
		return nil
	}

	finalModel := result.(CheckboxModel)
	return <-finalModel.Result
}
