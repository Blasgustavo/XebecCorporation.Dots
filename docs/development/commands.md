---
title: Comandos del CLI
description: Referencia completa de todos los comandos de XebecCorporation.Dots
---

# Comandos del CLI

> Referencia completa de todos los comandos disponibles en XebecCorporation.Dots

## Vista General

```
XebecCorporation.Dots es un CLI interactivo para instalar, 
configurar y mantener el ecosistema XEBEC CORPORATION.

Usage:
  xebec [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Configura componentes del ecosistema XEBEC
  help        Help about any command
  install     Instala herramientas del ecosistema XEBEC
  version     Muestra la versión del CLI
```

## Comandos

---

### `xebec --help`

Muestra la ayuda general del CLI.

```bash
xebec --help
# o
xebec -h
```

**Descripción**: Muestra información sobre comandos disponibles y opciones globales.

**Opciones**

| Opción | Alias | Descripción |
|--------|-------|-------------|
| `--help` | `-h` | Muestra ayuda |

---

### `xebec --version`

Muestra la versión del CLI.

```bash
xebec --version
# o
xebec -v
```

**Salida típica**:

```
XebecCorporation.Dots v0.1.0
```

---

### `xebec config`

Configura componentes del ecosistema XEBEC.

```bash
xebec config [subcomando]
```

**Subcomandos**

| Subcomando | Descripción |
|------------|-------------|
| `terminal` | Configura Alacritty |
| `shell` | Configura Nushell + Starship |

**Ejemplos**

```bash
# Configurar terminal
xebec config terminal

# Configurar shell
xebec config shell

# Ver ayuda
xebec config --help
```

---

### `xebec config terminal`

Aplica la configuración de Alacritty.

```bash
xebec config terminal [opciones]
```

**Opciones**

| Opción | Alias | Descripción | Default |
|--------|-------|-------------|---------|
| `--force` | `-f` | Sobrescribir config existente | false |
| `--theme` | `-t` | Tema a aplicar | xebec |

**Ejemplos**

```bash
# Aplicar configuración por defecto
xebec config terminal

# Forzar sobrescritura
xebec config terminal --force

# Aplicar tema específico
xebec config terminal --theme dracula
```

**Qué hace**:

1. Detecta el sistema operativo
2. Localiza el directorio de configuración de Alacritty
3. Copia `alacritty.toml` a la ubicación correcta
4. Aplica el tema XEBEC
5. Muestra mensaje de éxito/error

---

### `xebec config shell`

Configura Nushell y Starship.

```bash
xebec config shell [opciones]
```

**Opciones**

| Opción | Alias | Descripción | Default |
|--------|-------|-------------|---------|
| `--force` | `-f` | Sobrescribir config existente | false |
| `--prompt` | `-p` | Prompt a usar (starship/default) | starship |

**Ejemplos**

```bash
# Configurar con Starship
xebec config shell

# Configurar con prompt por defecto
xebec config shell --prompt default

# Forzar sobrescritura
xebec config shell --force
```

**Qué hace**:

1. Detecta Nushell instalado
2. Copia `config.nu` a `~/.config/nushell/config.nu`
3. Configura Starship
4. Inicializa Starship en Nushell

---

### `xebec install`

Instala herramientas del ecosistema XEBEC.

```bash
xebec install [subcomando]
```

**Subcomandos**

| Subcomando | Descripción |
|------------|-------------|
| `tools` | Instala herramientas adicionales |
| `deps` | Instala dependencias del sistema |

---

### `xebec install tools`

Instala herramientas adicionales.

```bash
xebec install tools [opciones]
```

**Opciones**

| Opción | Alias | Descripción | Default |
|--------|-------|-------------|---------|
| `--all` | `-a` | Instalar todas las herramientas | false |
| `--tool` | `-t` | Herramienta específica a instalar | todas |

**Herramientas disponibles**:

| Herramienta | Descripción |
|-------------|-------------|
| `fzf` | Fuzzy finder |
| `zoxide` | Smarter cd |
| `bat` | Cat clone con highlighting |
| `delta` | Git pager |
| `eza` | Modern ls |

**Ejemplos**

```bash
# Instalar todas las herramientas
xebec install tools --all

# Instalar herramienta específica
xebec install tools --tool fzf
```

---

### `xebec completion`

Genera script de autocompletado.

```bash
xebec completion [shell]
```

**Shells soportados**:

- `bash`
- `zsh`
- `fish`
- `powershell`
- `elvish`

**Ejemplos**

```bash
# Bash
xebec completion bash > /etc/bash_completion.d/xebec

# Zsh
xebec completion zsh > ~/.zsh/completions/_xebec

# Fish
xebec completion fish > ~/.config/fish/completions/xebec.fish
```

---

## Variables de Entorno

| Variable | Descripción |
|----------|-------------|
| `XEBEC_CONFIG_DIR` | Directorio de configuración |
| `XEBEC_CACHE_DIR` | Directorio de caché |
| `XEBEC_LOG_LEVEL` | Nivel de logging (debug, info, warn, error) |

---

## Archivo de Configuración

Ubicación: `~/.config/xebec/config.toml` (Linux) / `%APPDATA%\xebec\config.toml` (Windows)

```toml
# Configuración de XebecCorporation.Dots

[general]
theme = "xebec"
auto_update = true
log_level = "info"

[paths]
config = "~/.config/xebec"
data = "~/.local/share/xebec"

[terminal]
default_shell = "nu"
default_editor = "vim"

[tools]
auto_install = false
```

---

## Códigos de Salida

| Código | Significado |
|--------|-------------|
| 0 | Éxito |
| 1 | Error general |
| 2 | Error de sintaxis en argumentos |
| 3 | Error de permisos |
| 4 | Herramienta no encontrada |

---

*Consulta también: [Arquitectura del Proyecto](architecture.md)*
