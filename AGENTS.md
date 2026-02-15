# XebecCorporation.Dots - Project Context

## Project Overview

XebecCorporation.Dots es un CLI interactivo, moderno y multiplataforma (Windows + Linux) escrito en Go. Su propósito es instalar, configurar y mantener el ecosistema XEBEC CORPORATION en cualquier máquina mediante un binario único con menús TUI.

## Project Type

- **Language**: Go 1.21+
- **Build Tool**: Go modules
- **CLI Framework**: Cobra + Survey/Bubbletea
- **Target Platforms**: Windows 11, Linux (Arch, Debian/Ubuntu, Fedora)

## Project Structure

```
XebecCorporation.Dots/
├── cmd/xebec/           # Punto de entrada del CLI
├── internal/
│   ├── ui/              # Menús, animaciones, ASCII art
│   ├── os/              # Detección de SO, rutas, instaladores
│   └── actions/         # Flujos de instalación y configuración
├── configs/             # Archivos base (alacritty.toml, config.nu, starship.toml)
├── scripts/             # Scripts auxiliares
├── docs/                # Documentación
└── .opencode/           # Configuración de OpenCode (este archivo)
```

## Code Standards

- **Commits**: Conventional Commits (`feat:`, `fix:`, `docs:`, `chore:`)
- **Go**: gofmt, golangci-lint, unit tests con `go test ./...`
- **Error Handling**: Siempre retornar errores con contexto, usar `fmt.Errorf("descripción: %w", err)`
- **Naming**: camelCase para variables/funciones, PascalCase para exportados, snake_case para archivos de config

## Development Workflow

1. `go run ./cmd/xebec` - Ejecutar CLI en modo desarrollo
2. `go build ./cmd/xebec` - Compilar binario
3. `go test ./...` - Ejecutar tests
4. `golangci-lint run` - Linting (cuando esté configurado)

## Auto-Commit Rule

**IMPORTANTE**: Después de cada función implementada, DEBES ejecutar auto-commit:

```bash
python scripts/auto_commit.py
```

Esto commiteará y hará push automáticamente con mensaje conventional.

Para mensaje personalizado:
```bash
python scripts/auto_commit.py "feat: Nueva funcionalidad"
```

Para solo commit sin push:
```bash
python scripts/auto_commit.py --no-push
```

## Key Components

- **Orquestador**: Coordina la ejecución de agentes y subagentes para flujos completos de configuración.
- **Agentes**: Especializados en tareas específicas (configuración de terminal, shell, herramientas).
- **Subagentes**: Exploración de código, análisis de configuraciones, generación de archivos.

## Supported Tools

- Terminal: Alacritty (foco actual)
- Shell: Nushell + Starship
- Herramientas futuras: fzf, zoxide, bat, delta, eza
