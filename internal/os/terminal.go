// Package: os
// Detección de terminales instalados en el sistema
// author: XebecCorporation
// version: 1.0.0

package os

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Terminal detectado en el sistema
type Terminal struct {
	ID          string   // "alacritty", "wezterm", etc.
	Name        string   // Nombre para mostrar
	Icon        string   // Icono emoji
	Version     string   // Versión del terminal
	Installed   bool     // Si está instalado
	ConfigPath  string   // Ruta de configuración principal
	ConfigPaths []string // Rutas alternativas de configuración
	Exists      bool     // Si existe archivo de config
}

// Lista completa de terminales a detectar
var terminalList = []struct {
	id         string
	name       string
	icon       string
	commands   []string        // Comandos ejecutables a buscar
	paths      []string        // Rutas de ejecutables alternativas
	configs    []func() string // Funciones que retornan rutas de config
	versionCmd string          // Comando para obtener versión
}{
	// Terminales modernos populares
	{
		id:         "alacritty",
		name:       "Alacritty",
		icon:       "🖥️",
		commands:   []string{"alacritty"},
		configs:    []func() string{alacrittyConfig},
		versionCmd: "alacritty --version",
	},
	{
		id:         "wezterm",
		name:       "WezTerm",
		icon:       "🔥",
		commands:   []string{"wezterm"},
		configs:    []func() string{weztermConfig},
		versionCmd: "wezterm --version",
	},
	{
		id:         "kitty",
		name:       "Kitty",
		icon:       "🐱",
		commands:   []string{"kitty"},
		configs:    []func() string{kittyConfig},
		versionCmd: "kitty --version",
	},
	{
		id:         "ghostty",
		name:       "Ghostty",
		icon:       "👻",
		commands:   []string{"ghostty"},
		configs:    []func() string{ghosttyConfig},
		versionCmd: "ghostty --version",
	},
	{
		id:         "windows-terminal",
		name:       "Windows Terminal",
		icon:       "🪟",
		commands:   []string{"wt", "wt.exe"},
		configs:    []func() string{windowsTerminalConfig},
		versionCmd: "wt --version",
	},
	{
		id:         "hyper",
		name:       "Hyper",
		icon:       "⚡",
		commands:   []string{"hyper"},
		configs:    []func() string{hyperConfig},
		versionCmd: "hyper --version",
	},
	{
		id:         "tabby",
		name:       "Tabby",
		icon:       "📋",
		commands:   []string{"tabby"},
		configs:    []func() string{tabbyConfig},
		versionCmd: "tabby --version",
	},
	{
		id:         "windterm",
		name:       "WindTerm",
		icon:       "💨",
		commands:   []string{"WindTerm"},
		configs:    []func() string{windtermConfig},
		versionCmd: "windterm --version",
	},
	{
		id:         "electerm",
		name:       "Electerm",
		icon:       "🔌",
		commands:   []string{"electerm"},
		configs:    []func() string{electermConfig},
		versionCmd: "electerm --version",
	},

	// Terminales Linux populares
	{
		id:         "gnome-terminal",
		name:       "GNOME Terminal",
		icon:       "🐧",
		commands:   []string{"gnome-terminal", "gnome-terminal-server"},
		configs:    []func() string{gnomeTerminalConfig},
		versionCmd: "gnome-terminal --version",
	},
	{
		id:         "konsole",
		name:       "Konsole",
		icon:       "🎮",
		commands:   []string{"konsole"},
		configs:    []func() string{konsoleConfig},
		versionCmd: "konsole --version",
	},
	{
		id:         "terminator",
		name:       "Terminator",
		icon:       "🔱",
		commands:   []string{"terminator"},
		configs:    []func() string{terminatorConfig},
		versionCmd: "terminator --version",
	},
	{
		id:         "tilix",
		name:       "Tilix",
		icon:       "📦",
		commands:   []string{"tilix", "tilix.dcc"},
		configs:    []func() string{tilixConfig},
		versionCmd: "tilix --version",
	},
	{
		id:         "guake",
		name:       "Guake",
		icon:       "⬇️",
		commands:   []string{"guake"},
		configs:    []func() string{guakeConfig},
		versionCmd: "guake --version",
	},
	{
		id:         "yakuake",
		name:       "Yakuake",
		icon:       "⬆️",
		commands:   []string{"yakuake"},
		configs:    []func() string{yakuakeConfig},
		versionCmd: "yakuake --version",
	},
	{
		id:         "xfce4-terminal",
		name:       "XFCE Terminal",
		icon:       "🐆",
		commands:   []string{"xfce4-terminal"},
		configs:    []func() string{xfce4TerminalConfig},
		versionCmd: "xfce4-terminal --version",
	},
	{
		id:         "lxterminal",
		name:       "LXTerminal",
		icon:       "🪶",
		commands:   []string{"lxterminal"},
		configs:    []func() string{lxterminalConfig},
		versionCmd: "lxterminal --version",
	},
	{
		id:         "qterminal",
		name:       "QTerminal",
		icon:       "🟢",
		commands:   []string{"qterminal"},
		configs:    []func() string{qterminalConfig},
		versionCmd: "qterminal --version",
	},
	{
		id:         "lilyterm",
		name:       "LilyTerm",
		icon:       "🌸",
		commands:   []string{"lilyterm"},
		configs:    []func() string{lilytermConfig},
		versionCmd: "lilyterm --version",
	},
	{
		id:         "sakura",
		name:       "Sakura",
		icon:       "🌸",
		commands:   []string{"sakura"},
		configs:    []func() string{sakuraConfig},
		versionCmd: "sakura --version",
	},
	{
		id:         "st",
		name:       "st (Simple Terminal)",
		icon:       "📟",
		commands:   []string{"st"},
		configs:    []func() string{stConfig},
		versionCmd: "st --version",
	},
	{
		id:         "foot",
		name:       "foot",
		icon:       "🦶",
		commands:   []string{"foot"},
		configs:    []func() string{footConfig},
		versionCmd: "foot --version",
	},
	{
		id:         "rio",
		name:       "Rio Terminal",
		icon:       "🌊",
		commands:   []string{"rio"},
		configs:    []func() string{rioConfig},
		versionCmd: "rio --version",
	},

	// Terminales clásicos
	{
		id:         "xterm",
		name:       "XTerm",
		icon:       "❎",
		commands:   []string{"xterm"},
		configs:    []func() string{xtermConfig},
		versionCmd: "xterm -version",
	},
	{
		id:         "urxvt",
		name:       "URxvt / Rxvt-unicode",
		icon:       "📻",
		commands:   []string{"urxvt", "rxvt-unicode", "urxvt256c-ml", "urxvtc"},
		configs:    []func() string{urxvtConfig},
		versionCmd: "urxvt --version",
	},
	{
		id:         "eterm",
		name:       "Eterm",
		icon:       "🟣",
		commands:   []string{"Eterm"},
		configs:    []func() string{etermConfig},
		versionCmd: "Eterm --version",
	},
	{
		id:         "mlterm",
		name:       "MLTerm",
		icon:       "📺",
		commands:   []string{"mlterm"},
		configs:    []func() string{mltermConfig},
		versionCmd: "mlterm --version",
	},

	// Terminales macOS
	{
		id:         "iterm2",
		name:       "iTerm2",
		icon:       "💻",
		commands:   []string{"iTerm2"},
		configs:    []func() string{iterm2Config},
		versionCmd: "iTerm2 --version",
	},
	{
		id:         "terminal",
		name:       "Terminal.app",
		icon:       "🖥️",
		commands:   []string{"Terminal"},
		configs:    []func() string{terminalAppConfig},
		versionCmd: "osascript -e 'version of app \"Terminal\"'",
	},
	{
		id:         "alacritty",
		name:       "Alacritty (macOS)",
		icon:       "🖥️",
		commands:   []string{"alacritty"},
		configs:    []func() string{alacrittyConfig},
		versionCmd: "alacritty --version",
	},
	{
		id:         "hyper",
		name:       "Hyper (macOS)",
		icon:       "⚡",
		commands:   []string{"hyper"},
		configs:    []func() string{hyperConfig},
		versionCmd: "hyper --version",
	},

	// Terminales adicionales Windows
	{
		id:         "terminus",
		name:       "Terminus",
		icon:       "🏁",
		commands:   []string{"terminus"},
		configs:    []func() string{terminusConfig},
		versionCmd: "terminus --version",
	},
	{
		id:         "conemu",
		name:       "ConEmu",
		icon:       "⬛",
		commands:   []string{"ConEmu", "ConEmu64"},
		configs:    []func() string{conemuConfig},
		versionCmd: "ConEmu -Version",
	},
	{
		id:         "cmder",
		name:       "Cmder",
		icon:       "📦",
		commands:   []string{"cmder"},
		configs:    []func() string{cmderConfig},
		versionCmd: "cmder --version",
	},
	{
		id:         "fterminal",
		name:       "Fluent Terminal",
		icon:       "🌊",
		commands:   []string{"FluentTerminal"},
		configs:    []func() string{fluentTerminalConfig},
		versionCmd: "FluentTerminal --version",
	},
	{
		id:         "terminal-buddy",
		name:       "Terminal Buddy",
		icon:       "👥",
		commands:   []string{"TerminalBuddy"},
		configs:    []func() string{terminalBuddyConfig},
		versionCmd: "TerminalBuddy --version",
	},
}

// Detectar todos los terminales instalados
func DetectTerminals() []Terminal {
	var terminals []Terminal

	// Usar la lista completa de terminales
	for _, t := range terminalList {
		terminal := detectTerminalByCommands(t.id, t.name, t.icon, t.commands, t.configs, t.versionCmd)
		if terminal.Installed {
			terminals = append(terminals, terminal)
		}
	}

	// Si no hay ninguno instalado, agregar mensaje
	if len(terminals) == 0 {
		terminals = append(terminals, Terminal{
			ID:        "none",
			Name:      "No se detectó ningún terminal compatible",
			Icon:      "❌",
			Installed: false,
		})
	}

	return terminals
}

// Función genérica para detectar terminales por comandos
func detectTerminalByCommands(id, name, icon string, commands []string, configFuncs []func() string, versionCmd string) Terminal {
	t := Terminal{
		ID:   id,
		Name: name,
		Icon: icon,
	}

	// Buscar ejecutable en PATH
	for _, cmd := range commands {
		_, err := exec.LookPath(cmd)
		if err == nil {
			t.Installed = true
			break
		}
	}

	// Si no está instalado, verificar rutas alternativas
	if !t.Installed {
		t.Installed = searchInCommonPaths(commands)
	}

	// Si está instalado, detectar versión
	if t.Installed && versionCmd != "" {
		t.Version = getTerminalVersion(versionCmd)
	}

	// Buscar configuración
	for _, configFunc := range configFuncs {
		configPath := configFunc()
		if configPath != "" {
			if t.ConfigPath == "" {
				t.ConfigPath = configPath
			}
			t.ConfigPaths = append(t.ConfigPaths, configPath)
			if _, err := os.Stat(configPath); err == nil {
				t.Exists = true
			}
		}
	}

	return t
}

// getTerminalVersion ejecuta el comando de versión y parsea el resultado
func getTerminalVersion(versionCmd string) string {
	// Usar shell según la plataforma
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", versionCmd)
	} else {
		cmd = exec.Command("sh", "-c", versionCmd)
	}

	output, err := cmd.Output()
	if err != nil {
		return "N/A"
	}

	// Limpiar el output - quitar espacios y líneas extra
	version := strings.TrimSpace(string(output))

	// Si hay múltiples líneas, tomar solo la primera
	if strings.Contains(version, "\n") {
		version = strings.Split(version, "\n")[0]
	}

	// Truncar si es muy largo
	if len(version) > 20 {
		version = version[:20]
	}

	return version
}

// Buscar en rutas comunes de instalación
func searchInCommonPaths(commands []string) bool {
	home := os.Getenv("HOME")
	userProfile := os.Getenv("USERPROFILE")
	programFiles := os.Getenv("PROGRAMFILES")
	programFilesX86 := os.Getenv("PROGRAMFILES(X86)")

	// Rutas comunes en Windows
	winPaths := []string{
		filepath.Join(programFiles, "Alacritty"),
		filepath.Join(programFiles, "WezTerm"),
		filepath.Join(programFiles, "Kitty"),
		filepath.Join(programFiles, "Hyper"),
		filepath.Join(programFilesX86, "Alacritty"),
		filepath.Join(programFilesX86, "WezTerm"),
		filepath.Join(userProfile, "AppData", "Local", "Programs", "alacritty"),
		filepath.Join(userProfile, "AppData", "Local", "Programs", "wezterm"),
		filepath.Join(userProfile, "AppData", "Local", "Microsoft", "WindowsApps"),
	}

	// Rutas comunes en Linux/macOS
	nixPaths := []string{
		filepath.Join(home, ".local", "bin"),
		filepath.Join(home, ".local", "share", "applications"),
		"/usr/local/bin",
		"/usr/bin",
		"/opt",
	}

	var searchPaths []string
	if runtime.GOOS == "windows" {
		searchPaths = winPaths
	} else {
		searchPaths = nixPaths
	}

	for _, searchPath := range searchPaths {
		for _, cmd := range commands {
			// Buscar el ejecutable
			execPath := filepath.Join(searchPath, cmd)
			if _, err := os.Stat(execPath); err == nil {
				return true
			}
			// También buscar con .exe en Windows
			if runtime.GOOS == "windows" {
				execPath = filepath.Join(searchPath, cmd+".exe")
				if _, err := os.Stat(execPath); err == nil {
					return true
				}
			}
		}
	}

	return false
}

// ========== Funciones de configuración por terminal ==========

func alacrittyConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "alacritty", "alacritty.toml")
	case "linux":
		return filepath.Join(os.Getenv("HOME"), ".config", "alacritty", "alacritty.toml")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "alacritty", "alacritty.toml")
	}
	return ""
}

func weztermConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "wezterm", "wezterm.lua")
	case "linux", "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "wezterm", "wezterm.lua")
	}
	return ""
}

func kittyConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("USERPROFILE"), ".config", "kitty", "kitty.conf")
	case "linux", "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "kitty", "kitty.conf")
	}
	return ""
}

func ghosttyConfig() string {
	if runtime.GOOS != "windows" {
		return filepath.Join(os.Getenv("HOME"), ".config", "ghostty", "config")
	}
	return ""
}

func windowsTerminalConfig() string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), "Packages", "Microsoft.WindowsTerminal_8wekyb3d8bbwe", "LocalState", "settings.json")
}

func hyperConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "hyper", "config.json")
	case "linux", "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "hyper", "config.json")
	}
	return ""
}

func tabbyConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "tabby", "config.yaml")
	case "linux", "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "tabby", "config.yaml")
	}
	return ""
}

func windtermConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "WindTerm", "config")
	}
	return ""
}

func electermConfig() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "electerm", "config.json")
	case "linux", "darwin":
		return filepath.Join(os.Getenv("HOME"), ".config", "electerm", "config.json")
	}
	return ""
}

func gnomeTerminalConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "dconf", "user")
}

func konsoleConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "konsole", "konsole.rc")
}

func terminatorConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "terminator", "config")
}

func tilixConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "dconf", "user")
}

func guakeConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "guake", "guake.cfg")
}

func yakuakeConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "yakuake", "yakuakerc")
}

func xfce4TerminalConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "xfce4", "terminal", "accels.scm")
}

func lxterminalConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "lxterminal", "lxterminal.conf")
}

func qterminalConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "qterminal.org", "terminal.conf")
}

func lilytermConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "lilyterm", "default.cfg")
}

func sakuraConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "sakura", "sakura.conf")
}

func stConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "st", "config.h")
}

func footConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "foot", "foot.ini")
}

func rioConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "rio", "config.toml")
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "rio", "config.toml")
}

func xtermConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".Xresources")
}

func urxvtConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".Xresources")
}

func etermConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".Eterm", "themes")
}

func mltermConfig() string {
	return filepath.Join(os.Getenv("HOME"), ".mlterm", "main")
}

func iterm2Config() string {
	return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "iTerm2", "DynamicProfiles")
}

func terminalAppConfig() string {
	return filepath.Join(os.Getenv("HOME"), "Library", "Preferences", "com.apple.Terminal.plist")
}

func terminusConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "terminus", "config.json")
	}
	return ""
}

func conemuConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "ConEmu.xml")
	}
	return ""
}

func cmderConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("USERPROFILE"), "Cmder", "config")
	}
	return ""
}

func fluentTerminalConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "FluentTerminal", "config.json")
	}
	return ""
}

func terminalBuddyConfig() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "TerminalBuddy", "config.json")
	}
	return ""
}

// GetTerminalsForSelection devuelve opciones formateadas para el menú
func GetTerminalsForSelection() []Terminal {
	return DetectTerminals()
}

// String representation
func (t Terminal) String() string {
	status := "❌ No instalado"
	if t.Installed {
		if t.Exists {
			status = "✅ Configurado"
		} else {
			status = "⚙️ Sin configurar"
		}
	}
	return fmt.Sprintf("%s %s - %s", t.Icon, t.Name, status)
}

// GetTerminalByID busca un terminal por su ID
func GetTerminalByID(id string) *Terminal {
	terminals := DetectTerminals()
	for i, t := range terminals {
		if t.ID == id {
			return &terminals[i]
		}
	}
	return nil
}

// GetSupportedTerminals retorna lista de todos los terminales soportados
func GetSupportedTerminals() []string {
	ids := make([]string, len(terminalList))
	for i, t := range terminalList {
		ids[i] = t.id
	}
	return ids
}

// IsTerminalSupported verifica si un terminal específico es soportado
func IsTerminalSupported(id string) bool {
	for _, t := range terminalList {
		if t.id == id {
			return true
		}
	}
	return false
}

// GetTerminalInfo retorna información detallada de un terminal
func GetTerminalInfo(id string) (name, icon string, supported bool) {
	for _, t := range terminalList {
		if t.id == id {
			return t.name, t.icon, true
		}
	}
	return "", "", false
}

// GetTerminalsByPlatform filtra terminales por plataforma
func GetTerminalsByPlatform() []Terminal {
	var result []Terminal
	terminals := DetectTerminals()

	for _, t := range terminals {
		// Incluir todos los instalados, el código ya filtra por plataforma en las funciones de config
		if t.Installed {
			result = append(result, t)
		}
	}

	return result
}

// FormatTerminalList formatea la lista de terminales para mostrar
func FormatTerminalList(terminals []Terminal) string {
	var lines []string
	for i, t := range terminals {
		status := "❌"
		if t.Installed {
			if t.Exists {
				status = "✅"
			} else {
				status = "⚙️"
			}
		}
		lines = append(lines, fmt.Sprintf("%d. %s %s %s", i+1, t.Icon, t.Name, status))
	}
	return strings.Join(lines, "\n")
}
