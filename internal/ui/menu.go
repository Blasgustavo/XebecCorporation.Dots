// Package: ui
// Menú interactivo principal usando bubbletea
// author: XebecCorporation
// version: 1.0.0

package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/XebecCorporation/XebecCorporation.Dots/internal/os"
)

// GradientText aplica un gradiente a un texto
func GradientText(text string, startColor, endColor string) string {
	start := colorToRGB(startColor)
	end := colorToRGB(endColor)

	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return text
	}

	var result []string
	for i, line := range lines {
		if line == "" {
			result = append(result, "")
			continue
		}

		t := float64(i) / float64(len(lines)-1)
		if len(lines) == 1 {
			t = 0.5
		}

		r := uint8(float64(start.r) + t*float64(end.r-start.r))
		g := uint8(float64(start.g) + t*float64(end.g-start.g))
		b := uint8(float64(start.b) + t*float64(end.b-start.b))

		color := fmt.Sprintf("#%02x%02x%02x", r, g, b)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		result = append(result, style.Render(line))
	}

	return strings.Join(result, "\n")
}

// RGB para gradiente
type rgb struct {
	r, g, b uint8
}

func colorToRGB(s string) rgb {
	if len(s) > 0 && s[0] == '#' {
		s = s[1:]
	}
	if len(s) != 6 {
		return rgb{99, 99, 99}
	}

	var r, g, b uint8
	fmt.Sscanf(s, "%02x%02x%02x", &r, &g, &b)
	return rgb{r, g, b}
}

// Opción del menú
type MenuOption struct {
	ID          string
	Icon        string
	Title       string
	Description string
	Action      func() error
}

// Modelo del menú
type MenuModel struct {
	Options  []MenuOption
	Selected int
	Quitting bool
	Version  string
	Platform string
	Width    int
	Height   int
}

// Inicializar el modelo
func NewMenuModel(version string) MenuModel {
	// Si no se proporciona versión, usar la del branding
	if version == "" {
		version = GetVersion()
	}

	return MenuModel{
		Options:  getMainMenuOptions(),
		Selected: 0,
		Version:  version,
		Platform: getPlatformInfo(),
	}
}

// Obtener opciones del menú principal desde branding
func getMainMenuOptions() []MenuOption {
	menuOpts := GetMenuOptions()
	options := make([]MenuOption, len(menuOpts))

	for i, opt := range menuOpts {
		options[i] = MenuOption{
			ID:          opt.ID,
			Icon:        opt.Icon,
			Title:       opt.Title,
			Description: opt.Description,
		}
	}

	return options
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
		fmt.Println()
		fmt.Println(RenderInfo(fmt.Sprintf("%s %s", BrandingConfig.Texts.Executing, option.Title)))
		fmt.Println(MutedTextStyle.Render(option.Description))
		fmt.Println()

		// Ejecutar acción específica según la opción
		executeMenuAction(option.ID)

		return tea.Msg("executed:" + option.ID)
	}
}

// Renderizar la vista
func (m MenuModel) View() string {
	width := m.Width
	if width == 0 {
		width = 80
	}

	// Ancho del contenido interno (para el marco)
	contentWidth := width - 4
	if contentWidth < 60 {
		contentWidth = 60
	}

	cliLabel := BrandingConfig.Texts.CLILabel
	platformLabel := BrandingConfig.Texts.PlatformLabel
	menuTitle := BrandingConfig.Texts.MenuTitle
	footerNav := BrandingConfig.Texts.FooterNav
	separator := GetSeparator()

	// Estilos
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Width(contentWidth).
		Align(lipgloss.Center)

	infoStyle := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center)

	optionSelectedStyle := lipgloss.NewStyle().
		Foreground(AccentPurple).
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
		Width(contentWidth).
		Align(lipgloss.Center)

	separatorStyle := lipgloss.NewStyle().
		Foreground(Gray).
		Width(contentWidth).
		Align(lipgloss.Center)

	// Marco/borde
	borderStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(CorporateBlue).
		BorderBackground(GrayDark).
		Padding(1, 2)

	// Construir contenido interno
	content := ""

	// Banner con gradiente
	content += titleStyle.Render(GradientText(BannerASCII, BrandingConfig.Colors.GradientStart, BrandingConfig.Colors.GradientEnd)) + "\n"
	content += "\n"

	// Info de versión y plataforma
	content += infoStyle.Render(fmt.Sprintf("%s v%s  |  %s: %s", cliLabel, m.Version, platformLabel, m.Platform)) + "\n"
	content += separatorStyle.Render(separator) + "\n"
	content += "\n"

	// Título del menú
	content += titleStyle.Foreground(AccentPurple).Render(menuTitle) + "\n"
	content += "\n"

	// Opciones
	for i, option := range m.Options {
		if i == m.Selected {
			content += optionSelectedStyle.Render(fmt.Sprintf("► %s %s", option.Icon, option.Title)) + "\n"
			content += optionDescStyle.Render(option.Description) + "\n"
		} else {
			content += optionUnselectedStyle.Render(fmt.Sprintf("  %s %s", option.Icon, option.Title)) + "\n"
		}
	}

	content += "\n"
	content += separatorStyle.Render(separator) + "\n"
	content += "\n"

	// Footer
	content += footerStyle.Render(footerNav)

	// Aplicar el marco al contenido
	s := borderStyle.Width(contentWidth).Render(content)

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

// Ejecutar acción del menú
func executeMenuAction(optionID string) {
	switch optionID {
	case "terminal":
		showTerminalSelection()
	case "shell":
		fmt.Println(RenderInfo("🚀 Configurando Shell..."))
	case "tools":
		fmt.Println(RenderInfo("🛠️ Instalando herramientas..."))
	case "status":
		fmt.Println(RenderInfo("📊 Mostrando estado..."))
	case "backup":
		fmt.Println(RenderInfo("💾 Creando backup..."))
	case "restore":
		fmt.Println(RenderInfo("♻️ Restaurando backup..."))
	case "exit":
		fmt.Println(SuccessStyle.Render(BrandingConfig.Texts.Goodbye))
	default:
		fmt.Println(RenderInfo("Opción no implementada"))
	}
}

// Mostrar selección de terminal
func showTerminalSelection() {
	fmt.Println()
	fmt.Println(TitleStyle.Render("🖥️ Detectar Terminales Instalados"))
	fmt.Println()

	terminals := os.DetectTerminals()

	if len(terminals) == 0 {
		fmt.Println(MutedTextStyle.Render("No se detectó ningún terminal compatible"))
		return
	}

	for i, t := range terminals {
		status := "❌ No instalado"
		if t.Installed {
			if t.Exists {
				status = "✅ Configurado"
			} else {
				status = "⚙️ Sin configurar"
			}
		}
		fmt.Printf("%d. %s %s\n", i+1, t.Icon, TitleStyle.Render(t.Name))
		fmt.Printf("   %s\n", MutedTextStyle.Render(status))
		if t.ConfigPath != "" {
			fmt.Printf("   📁 %s\n", MutedTextStyle.Render(t.ConfigPath))
		}
		fmt.Println()
	}

	fmt.Print(PromptStyle.Render("Selecciona un terminal para configurar: "))
}

// Ejecutar menú simple (sin bubbletea para modo no-TTY)
func RunSimpleMenu(version string) error {
	ShowBanner()
	fmt.Println()

	options := getMainMenuOptions()

	// Usar textos del branding
	menuTitle := BrandingConfig.Texts.MenuTitle
	promptSel := BrandingConfig.Texts.PromptSel
	optionInvalid := BrandingConfig.Texts.OptionInvalid
	goodbye := BrandingConfig.Texts.Goodbye
	executing := BrandingConfig.Texts.Executing

	for {
		fmt.Println("\n" + TitleStyle.Render(menuTitle) + "\n")

		for i, option := range options {
			fmt.Printf("%d. %s\n", i+1, option.Title)
			fmt.Printf("   %s\n", MutedTextStyle.Render(option.Description))
		}

		fmt.Println()
		fmt.Print(PromptStyle.Render(promptSel))

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			return err
		}

		if choice < 1 || choice > len(options) {
			fmt.Println(ErrorStyle.Render(optionInvalid))
			continue
		}

		selected := options[choice-1]

		if selected.ID == "exit" {
			fmt.Println(SuccessStyle.Render(goodbye))
			return nil
		}

		fmt.Println()
		fmt.Println(RenderInfo(fmt.Sprintf("%s %s", executing, selected.Title)))
		fmt.Println(MutedTextStyle.Render(selected.Description))
		fmt.Println()
	}
}
