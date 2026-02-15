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
	IsBack      bool
	IsExit      bool
	ParentID    string
}

// MenuLevel representa un nivel en el historial de navegación
type MenuLevel struct {
	ID      string
	Title   string
	Options []MenuOption
}

// Modelo del menú con soporte para submenús
type MenuModel struct {
	CurrentLevel int
	History      []MenuLevel
	Selected     int
	Quitting     bool
	Version      string
	Platform     string
	Width        int
	Height       int
}

// Inicializar el modelo
func NewMenuModel(version string) MenuModel {
	if version == "" {
		version = GetVersion()
	}

	m := MenuModel{
		CurrentLevel: 0,
		History:      []MenuLevel{},
		Selected:     0,
		Version:      version,
		Platform:     getPlatformInfo(),
	}

	// Cargar menú principal
	m.loadMainMenu()

	return m
}

// Cargar menú principal
func (m *MenuModel) loadMainMenu() {
	options := getMenuOptionsFromBranding()
	m.History = []MenuLevel{
		{
			ID:      "main",
			Title:   BrandingConfig.Texts.MenuTitle,
			Options: options,
		},
	}
	m.CurrentLevel = 0
	m.Selected = 0
}

// Obtener opciones desde branding
func getMenuOptionsFromBranding() []MenuOption {
	menuOpts := GetMenuOptions()
	options := make([]MenuOption, len(menuOpts))

	for i, opt := range menuOpts {
		options[i] = MenuOption{
			ID:          opt.ID,
			Icon:        opt.Icon,
			Title:       opt.Title,
			Description: opt.Description,
			IsBack:      opt.ID == "back",
			IsExit:      opt.ID == "exit",
		}
	}

	return options
}

// Obtener opciones de submenú desde branding
func getSubmenuOptions(parentID string) []MenuOption {
	submenuOpts := GetSubmenu(parentID)
	if submenuOpts == nil {
		return nil
	}

	options := make([]MenuOption, len(submenuOpts))
	for i, opt := range submenuOpts {
		options[i] = MenuOption{
			ID:          opt.ID,
			Icon:        opt.Icon,
			Title:       opt.Title,
			Description: opt.Description,
			IsBack:      opt.ID == "back",
			ParentID:    parentID,
		}
	}

	return options
}

// Obtener opciones actuales
func (m *MenuModel) getCurrentOptions() []MenuOption {
	if m.CurrentLevel < len(m.History) {
		return m.History[m.CurrentLevel].Options
	}
	return []MenuOption{}
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
			if m.Selected < len(m.getCurrentOptions())-1 {
				m.Selected++
			}
		case "enter", " ":
			return m, m.executeSelected()
		case "left", "backspace", "h":
			// Volver al menú anterior
			if m.CurrentLevel > 0 {
				m.goBack()
			}
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
func (m *MenuModel) executeSelected() tea.Cmd {
	options := m.getCurrentOptions()
	if m.Selected >= len(options) {
		return nil
	}

	option := options[m.Selected]

	return func() tea.Msg {
		if option.IsExit {
			fmt.Println()
			fmt.Println(SuccessStyle.Render(BrandingConfig.Texts.Goodbye))
			return tea.Quit()
		}

		if option.IsBack {
			m.goBack()
			return nil
		}

		// Verificar si es un submenú
		submenuOpts := getSubmenuOptions(option.ID)
		if submenuOpts != nil {
			// Entrar al submenú
			m.History = append(m.History, MenuLevel{
				ID:      option.ID,
				Title:   option.Title,
				Options: submenuOpts,
			})
			m.CurrentLevel = len(m.History) - 1
			m.Selected = 0
			return nil
		}

		// Ejecutar acción de la opción
		executeMenuAction(option.ID)
		return nil
	}
}

// Volver al menú anterior
func (m *MenuModel) goBack() {
	if m.CurrentLevel > 0 {
		m.History = m.History[:m.CurrentLevel]
		m.CurrentLevel--
		m.Selected = 0
	}
}

// Renderizar la vista
func (m MenuModel) View() string {
	width := m.Width
	if width == 0 {
		width = 80
	}

	contentWidth := width - 4
	if contentWidth < 60 {
		contentWidth = 60
	}

	options := m.getCurrentOptions()
	currentLevel := m.History[m.CurrentLevel]

	cliLabel := BrandingConfig.Texts.CLILabel
	platformLabel := BrandingConfig.Texts.PlatformLabel
	separator := GetSeparator()
	isSubmenu := m.CurrentLevel > 0
	footerNav := GetFooterText(isSubmenu)

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

	backStyle := lipgloss.NewStyle().
		Foreground(AccentCyan).
		Padding(0, 2)

	footerStyle := lipgloss.NewStyle().
		Foreground(GrayLighter).
		Width(contentWidth).
		Align(lipgloss.Center)

	separatorStyle := lipgloss.NewStyle().
		Foreground(Gray).
		Width(contentWidth).
		Align(lipgloss.Center)

	borderStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(CorporateBlue).
		BorderBackground(GrayDark).
		Padding(1, 2)

	// Construir contenido
	content := ""

	// Banner con gradiente (solo en menú principal)
	if !isSubmenu {
		content += titleStyle.Render(GradientText(BannerASCII, BrandingConfig.Colors.GradientStart, BrandingConfig.Colors.GradientEnd)) + "\n"
		content += "\n"
		content += infoStyle.Render(fmt.Sprintf("%s v%s  |  %s: %s", cliLabel, m.Version, platformLabel, m.Platform)) + "\n"
		content += separatorStyle.Render(separator) + "\n"
		content += "\n"
	}

	// Título del menú actual
	menuTitle := currentLevel.Title
	content += titleStyle.Foreground(AccentPurple).Render(menuTitle) + "\n"
	content += "\n"

	// Opciones
	for i, option := range options {
		if i == m.Selected {
			if option.IsBack {
				content += backStyle.Render(fmt.Sprintf("► %s %s", option.Icon, option.Title)) + "\n"
			} else {
				content += optionSelectedStyle.Render(fmt.Sprintf("► %s %s", option.Icon, option.Title)) + "\n"
			}
			if option.Description != "" && !option.IsBack {
				content += optionDescStyle.Render(option.Description) + "\n"
			}
		} else {
			if option.IsBack {
				content += backStyle.Render(fmt.Sprintf("  %s %s", option.Icon, option.Title)) + "\n"
			} else {
				content += optionUnselectedStyle.Render(fmt.Sprintf("  %s %s", option.Icon, option.Title)) + "\n"
			}
		}
	}

	content += "\n"
	content += separatorStyle.Render(separator) + "\n"
	content += "\n"

	// Footer
	content += footerStyle.Render(footerNav)

	// Aplicar el marco
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
	fmt.Println()
	fmt.Println(RenderInfo(fmt.Sprintf("%s %s", BrandingConfig.Texts.Executing, getMenuActionTitle(optionID))))
	fmt.Println(MutedTextStyle.Render(getMenuActionDescription(optionID)))
	fmt.Println()

	switch optionID {
	case "terminal_alacritty":
		fmt.Println(SuccessStyle.Render("⚡ Configurando Alacritty..."))
		configureAlacritty()
	case "terminal_wezterm":
		fmt.Println(SuccessStyle.Render("🔥 Configurando WezTerm..."))
	case "terminal_kitty":
		fmt.Println(SuccessStyle.Render("🐱 Configurando Kitty..."))
	case "terminal_windows":
		fmt.Println(SuccessStyle.Render("🪟 Configurando Windows Terminal..."))
	case "shell_nushell":
		fmt.Println(SuccessStyle.Render("🐚 Configurando Nushell..."))
	case "shell_starship":
		fmt.Println(SuccessStyle.Render("⭐ Configurando Starship..."))
	case "shell_zsh":
		fmt.Println(SuccessStyle.Render("🦪 Configurando Zsh..."))
	case "shell_powershell":
		fmt.Println(SuccessStyle.Render("💜 Configurando PowerShell..."))
	case "tools_fzf":
		fmt.Println(SuccessStyle.Render("🔍 Instalando fzf..."))
	case "tools_zoxide":
		fmt.Println(SuccessStyle.Render("📍 Instalando zoxide..."))
	case "tools_bat":
		fmt.Println(SuccessStyle.Render("🦇 Instalando bat..."))
	case "tools_delta":
		fmt.Println(SuccessStyle.Render("📐 Instalando delta..."))
	case "tools_eza":
		fmt.Println(SuccessStyle.Render("📁 Instalando eza..."))
	case "tools_all":
		fmt.Println(SuccessStyle.Render("✨ Instalando todas las herramientas..."))
	case "status":
		showStatus()
	case "backup":
		fmt.Println(SuccessStyle.Render("💾 Creando backup..."))
	case "restore":
		fmt.Println(SuccessStyle.Render("♻️ Restaurando backup..."))
	default:
		fmt.Println(RenderInfo("Opción no implementada"))
	}
}

func getMenuActionTitle(id string) string {
	titles := map[string]string{
		"terminal_alacritty": "Alacritty",
		"terminal_wezterm":   "WezTerm",
		"terminal_kitty":     "Kitty",
		"terminal_windows":   "Windows Terminal",
		"shell_nushell":      "Nushell",
		"shell_starship":     "Starship",
		"shell_zsh":          "Zsh",
		"shell_powershell":   "PowerShell",
		"tools_fzf":          "fzf",
		"tools_zoxide":       "zoxide",
		"tools_bat":          "bat",
		"tools_delta":        "delta",
		"tools_eza":          "eza",
		"tools_all":          "Todas las herramientas",
		"status":             "Estado del sistema",
		"backup":             "Backup",
		"restore":            "Restaurar",
	}
	if title, ok := titles[id]; ok {
		return title
	}
	return id
}

func getMenuActionDescription(id string) string {
	descriptions := map[string]string{
		"terminal_alacritty": "Aplicando configuración de Alacritty",
		"terminal_wezterm":   "Aplicando configuración de WezTerm",
		"terminal_kitty":     "Aplicando configuración de Kitty",
		"terminal_windows":   "Aplicando configuración de Windows Terminal",
		"shell_nushell":      "Aplicando configuración de Nushell",
		"shell_starship":     "Aplicando configuración de Starship",
		"shell_zsh":          "Aplicando configuración de Zsh",
		"shell_powershell":   "Aplicando configuración de PowerShell",
		"tools_fzf":          "Instalando fzf - Buscador fuzzy",
		"tools_zoxide":       "Instalando zoxide - Navegador de directorios",
		"tools_bat":          "Instalando bat - Reemplazo de cat",
		"tools_delta":        "Instalando delta - Pager para git",
		"tools_eza":          "Instalando eza - Reemplazo de ls",
		"tools_all":          "Instalando todas las herramientas del ecosistema",
		"status":             "Mostrando estado de configuraciones",
		"backup":             "Creando copia de seguridad",
		"restore":            "Restaurando desde backup",
	}
	if desc, ok := descriptions[id]; ok {
		return desc
	}
	return ""
}

// Mostrar estado del sistema
func showStatus() {
	fmt.Println(TitleStyle.Render("📊 Estado del Sistema"))
	fmt.Println()

	sysInfo := DetectSystem()
	fmt.Printf("Sistema: %s\n", sysInfo.Platform)
	fmt.Printf("Arquitectura: %s\n", sysInfo.Architecture)
	fmt.Printf("Gestor de paquetes: %s\n", sysInfo.PackageMgr)
	fmt.Println()

	fmt.Println(TitleStyle.Render("🖥️ Terminales Detectados"))
	terminals := os.DetectTerminals()
	for _, t := range terminals {
		status := "❌ No instalado"
		if t.Installed {
			if t.Exists {
				status = "✅ Configurado"
			} else {
				status = "⚙️ Sin configurar"
			}
		}
		fmt.Printf("  %s %s - %s\n", t.Icon, t.Name, status)
	}
}

// Configurar Alacritty
func configureAlacritty() {
	fmt.Println(MutedTextStyle.Render("Detectando Alacritty..."))
	terminals := os.DetectTerminals()
	for _, t := range terminals {
		if t.ID == "alacritty" && t.Installed {
			fmt.Printf("Alacritty encontrado en: %s\n", t.ConfigPath)
			if t.Exists {
				fmt.Println(SuccessStyle.Render("✓ Alacritty ya está configurado"))
			} else {
				fmt.Println(WarningStyle.Render("⚠ Alacritty instalado pero sin configuración"))
				fmt.Println(MutedTextStyle.Render("Copiaremos la configuración base..."))
			}
			return
		}
	}
	fmt.Println(ErrorStyle.Render("✗ Alacritty no está instalado"))
	fmt.Println(MutedTextStyle.Render("Usa 'xebec install tools' para instalar herramientas"))
}

// Mostrar selección de terminal (legacy)
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

// Ejecutar menú simple (legacy)
func RunSimpleMenu(version string) error {
	ShowBanner()
	fmt.Println()

	options := getMenuOptionsFromBranding()

	menuTitle := BrandingConfig.Texts.MenuTitle
	promptSel := BrandingConfig.Texts.PromptSel
	optionInvalid := BrandingConfig.Texts.OptionInvalid
	goodbye := BrandingConfig.Texts.Goodbye
	executing := BrandingConfig.Texts.Executing

	for {
		fmt.Println("\n" + TitleStyle.Render(menuTitle) + "\n")

		for i, option := range options {
			fmt.Printf("%d. %s %s\n", i+1, option.Icon, option.Title)
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

		if selected.IsExit {
			fmt.Println(SuccessStyle.Render(goodbye))
			return nil
		}

		fmt.Println()
		fmt.Println(RenderInfo(fmt.Sprintf("%s %s", executing, selected.Title)))
		fmt.Println(MutedTextStyle.Render(selected.Description))
		fmt.Println()
	}
}
