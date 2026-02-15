// Package: ui
// Menú interactivo principal usando bubbletea
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Opción del menú
type MenuOption struct {
	ID          string
	Title       string
	Description string
	Action      func() error
}

// Modelo del menú
type MenuModel struct {
	Options     []MenuOption
	Selected    int
	Quitting    bool
	Version     string
	Platform    string
	Width       int
	Height      int
}

// Inicializar el modelo
func NewMenuModel(version string) MenuModel {
	return MenuModel{
		Options:  getMainMenuOptions(),
		Selected: 0,
		Version:  version,
		Platform: getPlatformInfo(),
	}
}

// Obtener opciones del menú principal
func getMainMenuOptions() []MenuOption {
	return []MenuOption{
		{
			ID:          "terminal",
			Title:       "Configurar Terminal",
			Description: "Configura Alacritty con el tema XEBEC",
		},
		{
			ID:          "shell",
			Title:       "Configurar Shell",
			Description: "Configura Nushell + Starship",
		},
		{
			ID:          "tools",
			Title:       "Instalar Herramientas",
			Description: "Instala fzf, zoxide, bat, delta, eza",
		},
		{
			ID:          "status",
			Title:       "Ver Estado",
			Description: "Muestra el estado de las configuraciones",
		},
		{
			ID:          "backup",
			Title:       "Crear Backup",
			Description: "Crea un backup de las configuraciones",
		},
		{
			ID:          "restore",
			Title:       "Restaurar Backup",
			Description: "Restaura desde un backup anterior",
		},
		{
			ID:          "exit",
			Title:       "Salir",
			Description: "Salir del CLI",
		},
	}
}

// Inicializar el programa
func (m MenuModel) Init() tea.Cmd {
	return nil
}

// Actualizar el modelo según los mensajes
func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
			}
		case "down", "j":
			if m.Selected < len(m.Options)-1 {
				m.Selected++
			}
		case "enter", " ":
			return m, executeSelected(m.Options[m.Selected])
		case "q", "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return m, nil
}

// Ejecutar la opción seleccionada
func executeSelected(option MenuOption) tea.Cmd {
	return func() tea.Msg {
		// Aquí se ejecutarían las acciones
		// Por ahora solo mostraremos un mensaje
		fmt.Println()
		fmt.Println(RenderInfo(fmt.Sprintf("Ejecutando: %s", option.Title)))
		fmt.Println(MutedTextStyle.Render(option.Description))
		fmt.Println()
		
		// Simular ejecución
		return tea.Msg("executed:" + option.ID)
	}
}

// Renderizar la vista
func (m MenuModel) View() string {
	// Calcular dimensiones
	width := m.Width
	if width == 0 {
		width = 80
	}
	
	// Estilos
	titleStyle := lipgloss.NewStyle().
		Foreground(CorporateBlue).
		Bold(true).
		Width(width).
		Align(lipgloss.Center)

	optionSelectedStyle := lipgloss.NewStyle().
		Foreground(CorporateBlue).
		Bold(true).
		Padding(0, 2)

	optionDescStyle := lipgloss.NewStyle().
		Foreground(GrayLighter).
		Padding(0, 4)

	optionUnselectedStyle := lipgloss.NewStyle().
		Foreground(GrayLight).
		Padding(0, 2)

	footerStyle := lipgloss.NewStyle().
		Foreground(GrayLighter).
		Width(width).
		Align(lipgloss.Center)

	separatorStyle := lipgloss.NewStyle().
		Foreground(Gray).
		Width(width).
		Align(lipgloss.Center)

	// Construir vista
	s := ""
	
	// Banner
	s += titleStyle.Render(BannerASCII) + "\n"
	s += "\n"
	
	// Info de versión y plataforma
	s += titleStyle.Width(width - 20).Render(fmt.Sprintf("CLI v%s  |  Platform: %s", m.Version, m.Platform)) + "\n"
	s += separatorStyle.Render("═══════════════════════════════════════════════════") + "\n"
	s += "\n"

	// Título del menú
	s += titleStyle.Render("Menú Principal") + "\n"
	s += "\n"

	// Opciones
	for i, option := range m.Options {
		if i == m.Selected {
			s += optionSelectedStyle.Render(fmt.Sprintf("► %s", option.Title)) + "\n"
			s += optionDescStyle.Render(option.Description) + "\n"
		} else {
			s += optionUnselectedStyle.Render(fmt.Sprintf("  %s", option.Title)) + "\n"
		}
	}

	s += "\n"
	s += separatorStyle.Render("═══════════════════════════════════════════════════") + "\n"
	s += "\n"

	// Footer
	s += footerStyle.Render("Presiona ↑/↓ para navegar, Enter para seleccionar, q para salir")

	return s
}

// Ejecutar el menú interactivo
func RunMenu(version string) error {
	p := tea.NewProgram(
		NewMenuModel(version),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	
	_, err := p.Run()
	return err
}

// Ejecutar menú simple (sin bubbletea para modo no-TTY)
func RunSimpleMenu(version string) error {
	ShowBanner()
	fmt.Println()
	
	options := getMainMenuOptions()
	
	for {
		fmt.Println("\n" + TitleStyle.Render("Menú Principal") + "\n")
		
		for i, option := range options {
			fmt.Printf("%d. %s\n", i+1, option.Title)
			fmt.Printf("   %s\n", MutedTextStyle.Render(option.Description))
		}
		
		fmt.Println()
		fmt.Print(PromptStyle.Render("Selecciona una opción: "))
		
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			return err
		}
		
		if choice < 1 || choice > len(options) {
			fmt.Println(ErrorStyle.Render("Opción inválida"))
			continue
		}
		
		selected := options[choice-1]
		
		if selected.ID == "exit" {
			fmt.Println(SuccessStyle.Render("¡Hasta luego!"))
			return nil
		}
		
		fmt.Println()
		fmt.Println(RenderInfo(fmt.Sprintf("Ejecutando: %s", selected.Title)))
		fmt.Println(MutedTextStyle.Render(selected.Description))
		fmt.Println()
	}
}
