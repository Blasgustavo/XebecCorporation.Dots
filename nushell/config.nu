$env.NU_DISABLE_BANNER = "true"
print "🚀 Nushell listo para despegar"


# ============================================
#  NUSHELL CONFIG – LIMPIO Y MINIMAL
# ============================================

# Directorio inicial
cd 'C:\Users\qty94'

# Indicadores de prompt simples
$env.PROMPT_INDICATOR = {|| "> " }
$env.PROMPT_INDICATOR_VI_INSERT = {|| ": " }
$env.PROMPT_INDICATOR_VI_NORMAL = {|| "> " }
$env.PROMPT_MULTILINE_INDICATOR = {|| "::: " }

# Conversión de variables de entorno
$env.ENV_CONVERSIONS = {
  "PATH": {
    from_string: { |s| $s | split row (char esep) | path expand --no-symlink }
    to_string: { |v| $v | path expand --no-symlink | str join (char esep) }
  }
  "Path": {
    from_string: { |s| $s | split row (char esep) | path expand --no-symlink }
    to_string: { |v| $v | path expand --no-symlink | str join (char esep) }
  }
}

# Directorios de scripts y plugins
$env.NU_LIB_DIRS = [
  ($nu.default-config-dir | path join 'scripts')
  ($nu.data-dir | path join 'completions')
]

$env.NU_PLUGIN_DIRS = [
  ($nu.default-config-dir | path join 'plugins')
]

# Editor
$env.EDITOR = "nvim"
$env.VISUAL = "nvim"

# PATH extendido
$env.PATH = (
  $env.PATH
  | split row (char esep)
  | prepend ($nu.home-path | path join ".local/bin")
  | append ($nu.home-path | path join ".cargo/bin")
)

# Prompt minimal sin módulos ni colores
$env.PROMPT_COMMAND = {|| "> " }
$env.PROMPT_COMMAND_RIGHT = {|| "" }

source ~/.cache/starship/init.nu
