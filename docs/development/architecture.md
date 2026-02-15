---
title: Arquitectura del Proyecto
description: Arquitectura y estructura del código de XebecCorporation.Dots
---

# Arquitectura del Proyecto

> Entiende la estructura y diseño del código de XebecCorporation.Dots

## Visión General

XebecCorporation.Dots es un CLI (Command Line Interface) escrito en **Go** que proporciona una forma unificada de instalar, configurar y mantener el ecosistema XEBEC CORPORATION.

### Tecnologías

| Tecnología | Propósito | Versión |
|------------|-----------|---------|
| **Go** | Lenguaje de programación | 1.21+ |
| **Cobra** | Framework CLI | v1.10.2 |
| **TOML** | Formato de configuración | - |

## Estructura del Proyecto

```
XebecCorporation.Dots/
├── cmd/                    # Punto de entrada del CLI
│   └── xebec/
│       ├── main.go        # Entry point
│       └── commands/      # Comandos Cobra
│           └── root.go
├── internal/              # Paquetes internos
│   ├── os/               # Detección de SO y rutas
│   ├── actions/          # Acciones de instalación/config
│   └── ui/               # Componentes de UI
├── configs/              # Archivos de configuración base
│   ├── alacritty.toml
│   ├── nushell/
│   └── starship.toml
├── docs/                # Documentación
├── scripts/             # Scripts auxiliares
└── .opencode/          # Configuración de OpenCode
```

## Arquitectura de Capas

```
┌─────────────────────────────────────┐
│           CLI (Cobra)               │  cmd/xebec/
├─────────────────────────────────────┤
│          Comandos                  │  commands/
├─────────────────────────────────────┤
│          Acciones                   │  internal/actions/
├─────────────────────────────────────┤
│    Detección de Sistema             │  internal/os/
├─────────────────────────────────────┤
│       Utilidades/UI                 │  internal/ui/
└─────────────────────────────────────┘
```

## Componentes Principales

### 1. Punto de Entrada (`cmd/xebec/main.go`)

```go
package main

import (
    "os"
    "github.com/XebecCorporation/XebecCorporation.Dots/cmd/xebec/commands"
)

func main() {
    if err := commands.Execute(); err != nil {
        os.Exit(1)
    }
}
```

**Responsabilidad**: Inicializar y ejecutar el CLI.

---

### 2. Comandos (`cmd/xebec/commands/`)

#### `root.go`

```go
var rootCmd = &cobra.Command{
    Use:   "xebec",
    Short: "XEBEC CORPORATION CLI",
    // ...
}

func Execute() error {
    return rootCmd.Execute()
}
```

**Responsabilidad**: Definir estructura de comandos y subcomandos.

---

### 3. Detección de Sistema (`internal/os/`)

#### Módulos principales:

| Módulo | Función |
|--------|---------|
| `detect.go` | Detectar SO, arquitectura, gestor de paquetes |
| `paths.go` | Rutas de configuración según plataforma |

```go
// Ejemplo: Detectar SO
type OSInfo struct {
    Name      string // windows, linux, darwin
    Arch      string // amd64, arm64
    PackageMgr string // winget, apt, pacman, dnf
}

func DetectOS() (*OSInfo, error) { ... }
```

---

### 4. Acciones (`internal/actions/`)

| Módulo | Función |
|--------|---------|
| `terminal.go` | Configurar Alacritty |
| `shell.go` | Configurar Nushell/Starship |
| `tools.go` | Instalar herramientas |

```go
// Ejemplo: Configurar Terminal
func ConfigureTerminal(theme string, force bool) error {
    // 1. Detectar SO
    // 2. Localizar directorio de config
    // 3. Copiar archivo de configuración
    // 4. Aplicar tema
}
```

---

### 5. UI (`internal/ui/`)

Componentes para interfaces TUI interactivas.

| Componente | Función |
|------------|---------|
| `menu.go` | Menús interactivos |
| `progress.go` | Barras de progreso |
| `spinner.go` | Indicadores de carga |

---

## Flujo de Ejecución

```
Usuario ejecuta comando
        │
        ▼
main.go
        │
        ▼
commands.Execute()
        │
        ▼
Cobra parsea argumentos
        │
        ▼
Se ejecuta función del comando
        │
        ├──► xebec config terminal
        │         │
        │         ▼
        │    actions.ConfigureTerminal()
        │         │
        │         ▼
        │    os.DetectOS() → os.GetConfigPath()
        │         │
        │         ▼
        │    files.Copy() → files.ApplyTheme()
        │
        └──► xebec install tools
                  │
                  ▼
             actions.InstallTools()
                  │
                  ▼
             os.DetectPackageManager()
                  │
                  ▼
             installer.Install()
```

## Patrones de Diseño

### 1. Error Handling

```go
// Siempre usar fmt.Errorf con wrapping
func ConfigureTerminal() error {
    osInfo, err := os.DetectOS()
    if err != nil {
        return fmt.Errorf("failed to detect OS: %w", err)
    }
    // ...
}
```

### 2. Detección de SO

```go
func GetConfigPath() (string, error) {
    switch runtime.GOOS {
    case "windows":
        return os.Getenv("APPDATA"), nil
    case "linux":
        return os.UserConfigDir()
    default:
        return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
    }
}
```

### 3. Configuración Flexible

```go
type Config struct {
    Theme      string `toml:"theme"`
    LogLevel   string `toml:"log_level"`
    AutoUpdate bool   `toml:"auto_update"`
}

func LoadConfig() (*Config, error) {
    // Cargar desde archivo o usar defaults
}
```

## Convenciones de Código

### Naming

| Tipo | Convención | Ejemplo |
|------|------------|---------|
| Variables | camelCase | `configPath` |
| Funciones exportadas | PascalCase | `ConfigureTerminal` |
| Funciones privadas | camelCase | `detectOS` |
| Constantes | PascalCase | `DefaultTheme` |
| Paquetes | snake_case | `internal/os` |

### Estructura de Archivos

```go
// Paquete: nombre_del_paquete
// Descripción breve de qué hace
// author: XebecCorporation
// version: 1.0.0

package packagename

import (
    // stdlib primero
    "fmt"
    "os"
    
    // luego paquetes externos
    "github.com/spf13/cobra"
)
```

### Testing

```go
// Archivo: package_test.go
package package_test

import (
    "testing"
)

func TestFunction(t *testing.T) {
    // Test código
}
```

---

## Extensibilidad

### Agregar Nuevo Comando

1. Crear archivo en `cmd/xebec/commands/`
2. Definir comando con Cobra
3. Agregar al root command

```go
// cmd/xebec/commands/nuevo.go
var nuevoCmd = &cobra.Command{
    Use:   "nuevo",
    Short: "Nuevo comando",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Lógica del comando
        return nil
    },
}

func init() {
    rootCmd.AddCommand(nuevoCmd)
}
```

### Agregar Nueva Acción

1. Crear archivo en `internal/actions/`
2. Definir función con lógica
3. Llamar desde comando

---

## Referencias

- [Documentación de Go](https://go.dev/doc/)
- [Cobra User Guide](https://github.com/spf13/cobra/blob/main/user_guide.md)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

---

*Consulta también: [Comandos del CLI](commands.md)*
