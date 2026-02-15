# Xebec Corporation Dots

[![Status](https://img.shields.io/badge/status-alpha-orange)](#) [![Platforms](https://img.shields.io/badge/platform-Windows%20%7C%20Linux-blue)](#plataformas-soportadas) [![Built%20with-Go](https://img.shields.io/badge/built%20with-Go%201.21+-00AEEF)](#stack-tecnológico) [![License: MIT](https://img.shields.io/badge/license-MIT-darkgray.svg)](#licencia)

> CLI interactivo, moderno y multiplataforma para instalar, configurar y mantener el ecosistema **XEBEC CORPORATION** en Windows y Linux con un único binario.

## TL;DR

- Binario en Go (`xebec`) con menús guiados, detección automática de SO y acciones idempotentes.
- Foco inicial: llevar el tema XEBEC a Alacritty, Nushell y Starship (Windows + Linux) con instaladores propios (`winget`, `scoop`, `apt`, `pacman`).
- Hoja de ruta: extender a herramientas complementarias (fzf, zoxide, bat, delta, eza), aplicar temas corporativos y mantener un historial de instalaciones.

## Objetivo general

Crear un CLI inspirado en Gentleman.Dots pero con identidad XEBEC. Debe:

- Distribuirse como un binario único (Go) fácil de publicar en Scoop/AUR/Goreleaser.
- Incluir una interfaz limpia, profesional, con colores corporativos (#00AEEF, #0A0A0A), animaciones sutiles y ASCII art propio.
- Guiar al usuario en la configuración de terminales, shells y herramientas, aplicando temas corporativos y manteniendo logs.

## Tabla de Contenidos

- [Objetivo general](#objetivo-general)
- [Stack tecnológico](#stack-tecnológico)
- [Funcionalidades del CLI](#funcionalidades-del-cli)
- [Inicio rápido](#inicio-rápido)
- [Plataformas soportadas](#plataformas-soportadas)
- [🎮 Entrenador de Maestría en Vim](#-entrenador-de-maestría-en-vim)
- [Documentación](#documentación)
- [Resumen de herramientas](#resumen-de-herramientas)
- [Bleeding Edge](#bleeding-edge)
- [Arquitectura del proyecto](#arquitectura-del-proyecto)
- [Contribuir](#contribuir)
- [Soporte](#soporte)
- [Licencia](#licencia)

## Stack tecnológico

- **Lenguaje:** Go 1.21+ (cross-compile sencillo hacia Windows/Linux, bins estáticos).
- **CLI/TUI:** [`cobra`](https://github.com/spf13/cobra) para comandos + [`survey`](https://github.com/go-survey/survey) o [`bubbletea`](https://github.com/charmbracelet/bubbletea) para menús interactivos con numeración.
- **Render/UI:** temas propios con #00AEEF, #0A0A0A, soporte para ASCII art y animaciones (spinners `lipgloss/bubbles`).
- **Distribución:** GoReleaser para binarios + buckets Scoop/AUR; pipelines GitHub Actions.

## Funcionalidades del CLI

### 1. Menú principal interactivo

- Configurar terminal (Alacritty, foco actual).
- Configurar shell (Nushell + Starship).
- Instalar herramientas (fzf, zoxide, bat, delta, eza).
- Aplicar tema XEBEC CORPORATION.
- Utilidades adicionales (backups, restauración, limpiar cache).
- Salir.

### 2. Detección automática del sistema operativo

- Windows 11+ (PowerShell + Scoop/winget).
- Linux (Debian/Ubuntu, Arch, Fedora). El CLI determina gestor de paquetes y rutas de config.

### 3. Acciones automatizadas

- Crear carpetas de configuración según SO (`AppData\Roaming\Alacritty`, `~/.config/alacritty`).
- Copiar o templar archivos (`configs/alacritty.toml`, `configs/config.nu`, `configs/starship.toml`, `configs/zellij.kdl`).
- Ejecutar instaladores necesarios (fzf, zoxide, etc.) usando `winget`, `scoop`, `apt`, `pacman` o `dnf`.
- Aplicar temas XEBEC (colores, ASCII art, transparencias) y mantener logs en `~/.local/share/xebec/logs` o `%ProgramData%`.

### 4. Estilo e interacción

- Interfaz limpia, profesional, menús numerados y copy corporativo.
- Animaciones simples (fade-in, loaders) mediante la librería TUI seleccionada.
- ASCII art XEBEC en la pantalla inicial + banner de éxito.

> **Foco inmediato:** implementar el flujo “Configurar terminal (Alacritty)” de punta a punta sobre Windows y Linux.

## Tabla de Contenidos

- [Inicio rápido](#inicio-rápido)
- [Plataformas soportadas](#plataformas-soportadas)
- [🎮 Entrenador de Maestría en Vim](#-entrenador-de-maestría-en-vim)
- [Documentación](#documentación)
- [Resumen de herramientas](#resumen-de-herramientas)
- [Bleeding Edge](#bleeding-edge)
- [Estructura del proyecto](#estructura-del-proyecto)
- [Contribuir](#contribuir)
- [Soporte](#soporte)
- [Licencia](#licencia)

## Inicio rápido

### Requisitos previos

- Windows 11 o alguna distro Linux soportada.
- Git 2.40+ y Go 1.21+ (compilar el binario `xebec`).
- Terminal con soporte TrueColor (Alacritty recomendado, el objetivo actual del CLI).

### Compilación local (Go)

```bash
git clone https://github.com/XebecCorporation/XebecCorporation.Dots.git
cd XebecCorporation.Dots
go run ./cmd/xebec        # Ejecuta el menú interactivo en modo desarrollo
go build ./cmd/xebec      # Genera el binario local
```

### Flujo previsto

1. `xebec` muestra ASCII art y detecta el SO.
2. El usuario selecciona “Configurar terminal (Alacritty)”.
3. El CLI copia `configs/alacritty.toml` a la ruta correcta, crea backup y aplica tema.
4. (Opcional) Instala dependencias (`winget install Alacritty`, `sudo pacman -S alacritty`).
5. Registra log e invita a configurar Nushell/Starship.

> Distribuciones via Scoop/AUR y binarios firmados llegarán una vez que el CLI básico esté operativo.

## Plataformas soportadas

| Plataforma | Estado | Requisitos mínimos |
| --- | --- | --- |
| Arch Linux / Endeavour | 🧪 beta | `go`, `git`, `alacritty`, `sudo` |
| Debian/Ubuntu 22.04+ | 🧪 beta | `build-essential`, `curl`, `alacritty` |
| Fedora 39+ | 🧪 beta | `dnf`, `go`, `alacritty` |
| Windows 11 | 🧪 beta | Windows Terminal o Alacritty, Scoop/Chocolatey, Nushell instalado |

> macOS queda fuera del alcance inmediato para concentrar esfuerzos en Windows + Linux.

## 🎮 Entrenador de Maestría en Vim

Entrenador interactivo estilo RPG integrado en el instalador. Cada módulo incluye 15 lecciones, modo práctica con selección automática, bosses y seguimiento de XP.

| Módulo | Teclas cubiertas |
| --- | --- |
| 🔤 Movimiento horizontal | `w`, `e`, `b`, `f`, `t`, `0`, `$`, `^` |
| ↕️ Movimiento vertical | `j`, `k`, `G`, `gg`, `{`, `}` |
| 📦 Objetos de texto | `iw`, `aw`, `i"`, `a(`, `it`, `at` |
| ✂️ Cambiar y repetir | `d`, `c`, `dd`, `cc`, `D`, `C`, `x` |
| 🔄 Sustitución | `r`, `R`, `s`, `S`, `~`, `gu`, `gU`, `J` |
| 🎬 Macros y registros | `qa`, `@a`, `@@`, `"ay`, `"+p` |
| 🔍 Regex / Búsqueda | `/`, `?`, `n`, `N`, `*`, `#`, `\v` |

> Próximo paso: publicar `docs/vim-trainer.md` con la especificación técnica y capturas animadas.

## Documentación

| Documento | Descripción | Estado |
| --- | --- | --- |
| `docs/installer.md` | Guía completa del instalador TUI, flags y flujos de backup/restore. | 🛠️ en redacción |
| `docs/manual-install.md` | Pasos manuales por plataforma, rutas de cada dotfile y scripts auxiliares. | 🛠️ en redacción |
| `docs/keymaps.md` | Referencia de atajos Neovim + integración IA. | 🛠️ en redacción |
| `docs/ai-config.md` | Configuración de Claude Code, Copilot y proveedores API. | 🛠️ en redacción |
| `docs/testing-e2e.md` | Tests end-to-end con Docker/Podman para validar instalaciones. | 🛠️ en redacción |
| `docs/contributing.md` | Setup local, lineamientos de PR y skills de agentes. | 🛠️ en redacción |

## Resumen de herramientas

### Emuladores de terminal

| Herramienta | Ruta en repo | Detalles |
| --- | --- | --- |
| Alacritty | `alacritty/alacritty.toml` | Tema oscuro con transparencia 0.85, JetBrains Mono y Nushell como shell por defecto. |
| Ghostty | _pendiente_ | Perfil GPU-first con bindings para el entrenador Vim. |
| Kitty | _pendiente_ | Enfoque en ligaduras y layout gráfico para sesiones múltiples. |
| WezTerm | `.wezterm.lua` (planificado) | Config con Lua para sincronizar perfiles remotos. |

### Shells

| Shell | Ruta | Detalles |
| --- | --- | --- |
| Nushell | `nushell/config.nu` | Prompt minimal, conversión PATH automática, integración con Starship. |
| Fish | _pendiente_ | Perfis y funciones para atajos TUI. |
| Zsh | _pendiente_ | Stack orientado a plugins asíncronos. |

### Multiplexores

| Herramienta | Estado | Notas |
| --- | --- | --- |
| Tmux | planificado | Config con sesiones predefinidas para proyectos IA. |
| Zellij | planificado | Plugins WebAssembly para métricas. |

### Editor

| Editor | Stack | Notas |
| --- | --- | --- |
| Neovim | LazyVim + LSP + IA | Configuración se publicará en `XebecCorporationNvim/` con perfiles por lenguaje. |

### Prompt

| Prompt | Ruta | Notas |
| --- | --- | --- |
| Starship | `starship.toml` (planificado) | Tema multi-shell con indicadores Git y modo focus. |

## Bleeding Edge

- Integración nativa con modelos IA (Claude Code, Gemini) para sugerencias in-editor.
- Exportación de perfiles hacia WSL/containers mediante `xebec export --target=docker`.
- Skills compartidas para agentes (carpeta `skills/`) con versionado independiente.

## Arquitectura del proyecto

```
XebecCorporation.Dots/
├── cmd/
│   └── xebec/            # main.go del CLI, comandos cobra
├── internal/
│   ├── ui/               # menús, animaciones, ASCII art
│   ├── os/               # detección SO, rutas, instaladores
│   └── actions/          # flujos de instalación/configuración
├── configs/              # alacritty.toml, config.nu, starship.toml, zellij.kdl
├── scripts/              # scripts auxiliares opcionales (post-install, tests)
├── docs/                 # documentación corporativa
└── README.md
```

| Componente | Rol | Notas |
| --- | --- | --- |
| `cmd/xebec` | Punto de entrada del CLI, definición de comandos y flags globales. | Usa Cobra y carga el menú principal. |
| `internal/ui` | Layout, colores, animaciones, ASCII art. | Implementa banner, loaders y estilo corporativo. |
| `internal/os` | Detecta SO, rutas y resolvers de gestores (`winget`, `apt`, `pacman`, `dnf`). | Expone structs `Platform` y helpers. |
| `internal/actions` | Funciones idempotentes para configurar terminal, shell, herramientas. | Cada acción escribe logs y realiza backups. |
| `configs/` | Archivos base XEBEC para copiar/templar. | Se versionan y se pueden parametrizar. |
| `scripts/` | Helpers para packaging, tests o validaciones (opcional). | Solo cuando sea imprescindible. |

### Configuraciones base (fase inicial)

- `configs/alacritty.toml` → tema oscuro, JetBrains Mono, transparencia 0.85.
- `configs/config.nu` → Nushell sin banner, PATH expandido, Starship integrado.
- `configs/starship.toml` → prompt corporativo con indicadores Git.
- `configs/zellij.kdl` → layout predefinido (cuando se añada Zellij).

## Contribuir

1. Haz fork y clona el repositorio.
2. Crea una rama descriptiva (`feat/ui-menu`, `feat/os-detector`).
3. Ejecuta `go test ./...` y `golangci-lint run` (cuando esté configurado) antes del PR.
4. Sigue el estilo Conventional Commits (`feat:`, `fix:`, `docs:`) y anexa capturas o GIFs para demos del CLI/TUI.
5. Indica plataforma probada (Windows/Linux) y comandos utilizados.

## Soporte

- Abre un issue en GitHub para bugs y requests.
- Canal de chat comunitario (Discord) en preparación; comparte feedback vía issues mientras tanto.

## Licencia

Proyecto liberado bajo licencia MIT. Contribuciones aceptan el mismo licenciamiento.

¡Feliz coding! 🎩
