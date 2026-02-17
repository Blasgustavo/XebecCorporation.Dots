// Package: actions
// Acciones de configuración de terminales
// author: XebecCorporation
// version: 1.0.0

package actions

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Opciones de configuración de Alacritty
type AlacrittyConfigOptions struct {
	Window bool // decorations, opacity, startup_mode, padding
	Colors bool // colors.primary, colors.normal, colors.bright
	Font   bool // font.normal, font.bold, font.italic, font.size
	Cursor bool // cursor.style, cursor.thickness
	Shell  bool // terminal.shell, terminal.osc52
}

// AlacrittyConfigOption representa una opción en el menú de checkboxes
type AlacrittyConfigOption struct {
	ID          string
	Title       string
	Description string
	Key         string // "window", "colors", "font", "cursor", "shell"
}

// GetAlacrittyConfigOptions retorna las opciones disponibles para configurar
func GetAlacrittyConfigOptions() []AlacrittyConfigOption {
	return []AlacrittyConfigOption{
		{
			ID:          "window",
			Title:       "Ventana",
			Description: "decorations, opacity, startup_mode, padding",
			Key:         "window",
		},
		{
			ID:          "colors",
			Title:       "Colores",
			Description: "Tema XEBEC - primary, normal, bright",
			Key:         "colors",
		},
		{
			ID:          "font",
			Title:       "Fuente",
			Description: "JetBrains Mono, tamaño",
			Key:         "font",
		},
		{
			ID:          "cursor",
			Title:       "Cursor",
			Description: "shape (Beam), blinking",
			Key:         "cursor",
		},
		{
			ID:          "shell",
			Title:       "Shell",
			Description: "Program y argumentos del shell",
			Key:         "shell",
		},
	}
}

// GetAlacrittyConfigPath retorna la ruta de configuración según el SO
func GetAlacrittyConfigPath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "alacritty", "alacritty.toml")
	case "linux", "darwin":
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			home := os.Getenv("HOME")
			xdgConfig = filepath.Join(home, ".config")
		}
		return filepath.Join(xdgConfig, "alacritty", "alacritty.toml")
	}
	return ""
}

// GetAlacrittyConfigDir retorna el directorio de configuración
func GetAlacrittyConfigDir() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "alacritty")
	case "linux", "darwin":
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			home := os.Getenv("HOME")
			xdgConfig = filepath.Join(home, ".config")
		}
		return filepath.Join(xdgConfig, "alacritty")
	}
	return ""
}

// EnsureAlacrittyDir crea el directorio de configuración si no existe
func EnsureAlacrittyDir() error {
	dir := GetAlacrittyConfigDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creando directorio %s: %w", dir, err)
		}
		fmt.Printf("✓ Directorio creado: %s\n", dir)
	}
	return nil
}

// BackupAlacrittyConfig hace un backup del archivo de configuración existente
func BackupAlacrittyConfig() (string, error) {
	configPath := GetAlacrittyConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "", nil // No hay archivo existente
	}

	// Crear directorio de backups
	backupDir := filepath.Join(GetAlacrittyConfigDir(), "backups")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("error creando directorio de backups: %w", err)
	}

	// Nombre del backup con timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	backupPath := filepath.Join(backupDir, fmt.Sprintf("alacritty_%s.toml", timestamp))

	// Copiar archivo
	src, err := os.Open(configPath)
	if err != nil {
		return "", fmt.Errorf("error abriendo archivo original: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(backupPath)
	if err != nil {
		return "", fmt.Errorf("error creando backup: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("error copiando backup: %w", err)
	}

	return backupPath, nil
}

// GetSourceConfigPath retorna la ruta del archivo de configuración base
func GetSourceConfigPath() string {
	// Obtener la ruta del proyecto
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
	return filepath.Join(projectRoot, "alacritty", "alacritty.toml")
}

// ConfigureAlacritty aplica la configuración de Alacritty según las opciones seleccionadas
func ConfigureAlacritty(opts AlacrittyConfigOptions) error {
	// Verificar que Alacritty esté instalado
	if !IsAlacrittyInstalled() {
		return fmt.Errorf("Alacritty no está instalado en el sistema")
	}

	// Crear directorio si no existe
	if err := EnsureAlacrittyDir(); err != nil {
		return fmt.Errorf("error asegurando directorio: %w", err)
	}

	// Hacer backup si existe configuración
	backupPath, err := BackupAlacrittyConfig()
	if err != nil {
		return fmt.Errorf("error en backup: %w", err)
	}
	if backupPath != "" {
		fmt.Printf("✓ Backup creado: %s\n", backupPath)
	}

	// Leer configuración base
	sourcePath := GetSourceConfigPath()
	sourceData, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error leyendo configuración base: %w", err)
	}

	// Filtrar según opciones seleccionadas
	configContent := filterConfigByOptions(string(sourceData), opts)

	// Escribir configuración
	destPath := GetAlacrittyConfigPath()
	if err := os.WriteFile(destPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("error escribiendo configuración: %w", err)
	}

	fmt.Printf("✓ Configuración aplicada: %s\n", destPath)
	return nil
}

// filterConfigByOptions filtra el contenido del config según las opciones seleccionadas
func filterConfigByOptions(content string, opts AlacrittyConfigOptions) string {
	lines := strings.Split(content, "\n")
	var result []string
	var includeSection bool

	// Mapeo de secciones a opciones
	sectionToOption := map[string]string{
		"window":   "Window",
		"font":     "Font",
		"cursor":   "Cursor",
		"colors":   "Colors",
		"terminal": "Shell",
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Detectar inicio de sección
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			section := strings.Trim(trimmed, "[]")

			// Determinar si es una sección principal
			mainSection := strings.Split(section, ".")[0]

			// Determinar si incluir esta sección
			if optKey, ok := sectionToOption[mainSection]; ok {
				switch optKey {
				case "Window":
					includeSection = opts.Window
				case "Colors":
					includeSection = opts.Colors
				case "Font":
					includeSection = opts.Font
				case "Cursor":
					includeSection = opts.Cursor
				case "Shell":
					includeSection = opts.Shell
				}
			} else if mainSection == "terminal" && !opts.Shell {
				// Excluir [terminal] completamente si no se seleccionó Shell
				includeSection = false
			} else if mainSection == "" {
				// Sección vacía o comentarios
				includeSection = true
			} else {
				// Para otras secciones, incluir por defecto si hay algo seleccionado
				includeSection = opts.Window || opts.Colors || opts.Font || opts.Cursor || opts.Shell
			}
		}

		// Incluir línea si estamos en una sección válida
		if includeSection || trimmed == "" || strings.HasPrefix(trimmed, "#") {
			result = append(result, line)
		}
	}

	// Si no hay opciones seleccionadas, no escribir nada
	if !opts.Window && !opts.Colors && !opts.Font && !opts.Cursor && !opts.Shell {
		return "# Configuración vacía - ninguna opción seleccionada\n"
	}

	return strings.Join(result, "\n")
}

// IsAlacrittyInstalled verifica si Alacritty está instalado
func IsAlacrittyInstalled() bool {
	// Buscar en PATH
	paths := []string{"alacritty"}
	if runtime.GOOS == "windows" {
		paths = []string{"alacritty.exe", "alacritty"}
	}

	for _, p := range paths {
		if _, err := exec.LookPath(p); err == nil {
			return true
		}
	}

	// En Windows también buscar en ubicaciones comunes
	if runtime.GOOS == "windows" {
		possiblePaths := []string{
			filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "Alacritty", "alacritty.exe"),
			filepath.Join(os.Getenv("ProgramFiles"), "Alacritty", "alacritty.exe"),
			filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Alacritty", "alacritty.exe"),
		}
		for _, p := range possiblePaths {
			if _, err := os.Stat(p); err == nil {
				return true
			}
		}
	}

	return false
}

// GetAlacrittyStatus retorna el estado actual de Alacritty
func GetAlacrittyStatus() (installed bool, configured bool, configPath string) {
	configPath = GetAlacrittyConfigPath()
	installed = IsAlacrittyInstalled()

	if _, err := os.Stat(configPath); err == nil {
		configured = true
	}

	return
}
