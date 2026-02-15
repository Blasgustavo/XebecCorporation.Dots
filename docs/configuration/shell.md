---
title: Configuración de Shell
description: Guía para configurar Nushell y Starship prompt
---

# Configuración de Shell

> Configura Nushell como shell principal y Starship como prompt

## Nushell

| Campo | Valor |
|-------|-------|
| **Nombre** | Nushell |
| **Versión** | 0.90+ |
| **Tipo** | Shell |
| **Repo** | [nushell/nushell](https://github.com/nushell/nushell) |
| **Web** | [nushell.rs](https://www.nushell.rs) |

### Por qué Nushell?

- **Datos estructurados**: Cada salida es una tabla estructurada
- **Sintaxis moderna**: Pipeline intuitivo similar a FP
- **Plugins**: Extensible con plugins Rust
- **Cross-platform**: Windows, Linux, macOS
- **Configuración en Nu**: Todo en `.nu`

## Starship

| Campo | Valor |
|-------|-------|
| **Nombre** | Starship |
| **Versión** | 1.16+ |
| **Tipo** | Prompt |
| **Repo** | [starship/starship](https://github.com/starship/starship) |
| **Web** | [starship.rs](https://starship.rs) |

### Por qué Starship?

- **Cross-shell**: Funciona en Bash, Zsh, Fish, PowerShell
- **Rápido**: Mínimo tiempo de carga
- **Personalizable**: Temas y módulos configurables
- **Información relevante**: Git, Node, Rust, Go, y más

## Instalación

### Nushell

#### Windows

```powershell
winget install Nushell.Nushell
```

#### Linux

```bash
# Ubuntu/Debian
sudo apt install nushell

# Arch
sudo pacman -S nushell

# Fedora
sudo dnf install nushell
```

### Starship

```bash
# Instalador automático
curl -sS https://starship.rs/install.sh | sh
```

O usando gestores de paquetes:

```bash
# Homebrew
brew install starship

# Pacman
sudo pacman -S starship
```

## Configuración de Nushell

### Ubicación de Configuración

| Sistema | Ruta |
|---------|------|
| Windows | `%APPDATA%\nushell\config.nu` |
| Linux | `~/.config/nushell/config.nu` |

### Configuración Base

El proyecto incluye una configuración base en `nushell/config.nu`:

```nu
# Configuración Nushell XEBEC

# Directorio de configuración
let config_dir = ($env.XDG_CONFIG_HOME | default { $env.HOME | path join '.config' } | path join 'nushell')

# Tema
$env.config = {
    show_banner: false
    
    ls: {
        use_ls_colors: true
        clickable_links: true
    }
    
    rm: {
        always_trash: false
    }
    
    table: {
        mode: rounded
        index_mode: always
        show_empty: true
    }
    
    explore: {
        exit_esc: true
        command_bar_text: '#ffffff'
    }
    
    history: {
        max_size: 10000
        sync_on_enter: true
        shell_integration: true
    }
    
    completions: {
        case_sensitive: false
        quick: true
        partial: true
        algorithm: fuzzy
    }
    
    filesize: {
        metric: false
        format: 'auto'
    }
    
    cursor_shape: {
        emacs: line
        vi_insert: block
        vi_normal: underscore
    }
    
    color_config: {
        separator: white
        leading_trailing_space_bg: { attr: n }
        header: green_bold
        empty: blue
        bool: light_cyan
        int: white
        filesize: cyan
        duration: white
        date: purple
        range: white
        float: white
        string: white
        nothing: white
        binary: white
        cell-path: white
        row_index: green_bold
        record: white
        list: white
        block: white
        hints: dark_gray
        search_result: { bg: red fg: white }
        shape_and: purple_bold
        shape_binary: purple_bold
        shape_block: blue_bold
        shape_bool: light_cyan
        shape_closure: green_bold
        shape_custom: green
        shape_datetime: cyan_bold
        shape_directory: cyan
        shape_external: cyan
        shape_externalarg: green_bold
        shape_filepath: cyan
        shape_flag: blue_bold
        shape_float: purple_bold
        shape_garbage: { fg: white bg: red attr: b }
        shape_globpattern: cyan_bold
        shape_int: purple_bold
        shape_internalcall: cyan_bold
        shape_list: cyan_bold
        shape_literal: blue
        shape_match_pattern: green
        shape_matching_brackets: { attr: u }
        shape_nothing: light_cyan
        shape_operator: yellow
        shape_or: purple_bold
        shape_pipe: purple_bold
        shape_range: yellow_bold
        shape_record: cyan_bold
        shape_redirection: purple_bold
        shape_signature: green_bold
        shape_string: green
        shape_string_interpolation: cyan_bold
        shape_table: blue_bold
        shape_variable: purple
        shape_vardecl: purple
    }
    
    use_grid_icons: true
    footer_mode: '25'
    float_precision: 2
    buffer_editor: { $env.EDITOR | default { 'vim' } }
    use_ansi_coloring: true
    bracketed_paste: true
    edit_mode: vi
    shell_integration: true
    render_right_prompt_on_last_line: false
    
    hooks: {
        pre_prompt: [{ null }]
        pre_execution: [{ null }]
        env_change: {
            PWD: [{|before, after| null }]
        }
        display_output: 'if (term size).columns >= 100 { table -e } else { table }'
        command_not_found: { null }
    }
    
    menus: [
        {
            name: completion_menu
            only_buffer_difference: false
            marker: |sp| if $sp == '' { '█' } else { $sp }
            type: {
                layout: columnar
                columns: 4
                col_width: 20
                col_padding: 2
            }
            style: {
                text: green
                selected_text: green_reverse
                description_text: yellow
            }
        }
    ]
    
    keybindings: [
        {
            name: completion_menu
            modifier: none
            keycode: tab
            mode: [emacs vi_normal vi_insert]
            event: {
                until: [
                    { send: menu name: completion_menu }
                    { send: menunext }
                ]
            }
        }
    ]
}

# Alias útiles
alias ll = ls -la
alias la = ls -a
alias l = ls -l
alias .. = cd ..
alias ... = cd ../..
alias .... = cd ../../..

# Funciones personalizadas
def mkcd [dir: string] {
    mkdir $dir
    cd $dir
}

# Configuración de Starship
mkdir ($config_dir | path join 'starship')
$env.STARSHIP_CONFIG = ($config_dir | path join 'starship' | path join 'starship.toml')
$env.STARSHIP_CACHE = ($config_dir | path join 'starship' | path join 'cache')
```

## Configuración de Starship

### Ubicación

| Sistema | Ruta |
|---------|------|
| Windows | `%APPDATA%\starship.toml` |
| Linux | `~/.config/starship.toml` |

### Configuración Base

```toml
# Configuración Starship XEBEC

format = """
[┌───────────────────────────────────────────────────────>](bold green)
[│](bold green)$directory$rust$python$golang$nodejs$git_branch$git_status
[└>](bold green) $character """

add_newline = true

[character]
success_symbol = "[❯](bold green)"
error_symbol = "[✗](bold red)"
vimcmd_symbol = "[❮](bold green)"

[directory]
truncation_length = 3
truncate_to_repo = true
format = "[$read_only]($read_only )[$symbol$path]($style)[](white)"

[git_branch]
symbol = ""
format = "on [$symbol$branch]($style) "

[git_status]
format = '([\[$all_status$ahead_behind\]]($style) )'
conflicted = "⚔️ "
ahead = "⇡${count}"
behind = "⇣${count}"
diverged = "⇕⇡${ahead_count}⇣${behind_count}"
untracked = "?${count}"
stashed = "📦"
modified = "!${count}"
staged = '+${count}'
renamed = "»${count}"
deleted = "✘${count}"

[golang]
symbol = "🐹 "
format = "via [$symbol($version )]($style)"

[python]
symbol = "🐍 "
format = "via [$symbol($virtualenv )($version )]($style)"
python_binary = ["python3", "python"]

[nodejs]
symbol = "⬢ "
format = "via [$symbol($version )]($style)"

[rust]
symbol = "🦀 "
format = "via [$symbol($version )]($style)"

[docker_context]
symbol = "🐳 "
format = "via [$symbol$context ]($style) "

[package]
symbol = "📦 "
format = "is [$symbol$version ]($style) "

[time]
disabled = true

[aws]
format = 'on [$symbol($profile )(\($region\) )]($style)'

[docker]
symbol = "🐳 "
```

## Aplicar Configuración con CLI

```bash
# Configurar Nushell
xebec config shell
```

Este comando:
1. Detecta el shell instalado
2. Copia `config.nu` a la ubicación correcta
3. Integra Starship
4. Configura el prompt

## Solución de Problemas

### Nushell no inicia

Verifica que tienes la versión correcta:

```bash
nu --version
```

### Starship no aparece

Añade esto al final de tu `config.nu`:

```nu
mkdir ($env.XDG_CONFIG_HOME | default { $env.HOME | path join '.config' } | path join 'starship')
$env.STARSHIP_CONFIG = ($env.XDG_CONFIG_HOME | default { $env.HOME | path join '.config' } | path join 'starship' | path join 'starship.toml')
starship init nu | save -f ($env.XDG_CONFIG_HOME | default { $env.HOME | path join '.config' } | path join 'starship' | path join 'init.nu')
```

---

*Consulta también: [Configuración de Terminal](terminal.md) y [Herramientas](tools.md)*
