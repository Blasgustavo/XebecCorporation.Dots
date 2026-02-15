---
name: auto-commit
description: Auto-commit y push automático después de implementar funciones en XebecCorporation.Dots.
---

# Skill: auto-commit

# Auto-Commit y Push Automático

Este skill automatiza el proceso de commit y push después de implementar funciones en XebecCorporation.Dots.

## Descripción

El skill detecta cambios en el repositorio, genera automáticamente un mensaje de commit basado en los archivos modificados y ejecuta commit + push.

## Activación

Este skill se auto-invoca automáticamente después de cada función implementada. No requiere activación manual.

## Tipos de Cambio Detectados

| Tipo | Icono | Archivos |
|------|-------|----------|
| `feat` | ✨ | `cmd/`, `internal/` |
| `fix` | 🐛 | Archivos con "fix", "bug", "error" |
| `docs` | 📚 | `docs/`, `README`, `*.md` |
| `chore` | 🔧 | `scripts/`, `.gitignore`, configuración |
| `build` | 📦 | `go.mod`, `go.sum`, dependencias |
| `test` | ✅ | `*_test.go`, archivos de test |
| `refactor` | ♻️ | Cambios en código existente |

## Flujo de Ejecución

1. **Detectar cambios**: `git status --porcelain`
2. **Analizar archivos**: Determinar tipo de cambio
3. **Generar mensaje**: Crear mensaje conventional commit
4. **Ejecutar commit**: `git add -A && git commit -m "mensaje"`
5. **Push**: `git push`

## Comandos Disponibles

### Auto-commit completo (commit + push)
```
auto-commit
```

### Solo commit (sin push)
```
auto-commit --no-push
```

### Commit con mensaje personalizado
```
auto-commit "feat: Nueva funcionalidad"
```

## Ejemplo de Uso

```bash
# Después de implementar una función:
# 1. El agente implementa la función
# 2. El skill auto-commit se activa
# 3. Se hace commit automático con mensaje apropiado
# 4. Se hace push al remoto
```

## Script de Ejecución

El skill usa el script `scripts/auto_commit.py`:

```python
#!/usr/bin/env python3
import subprocess
import sys

def main():
    message = None
    push = True
    
    # Parsear argumentos
    for arg in sys.argv[1:]:
        if arg == "--no-push":
            push = False
        elif not arg.startswith("-"):
            message = arg
    
    # Obtener cambios
    result = subprocess.run(["git", "status", "--porcelain"], capture_output=True, text=True)
    changes = result.stdout.strip()
    
    if not changes:
        print("[INFO] No hay cambios para commitear")
        return
    
    # Detectar tipo y generar mensaje
    if not message:
        message = detect_and_generate_message(changes)
    
    # Commit
    subprocess.run(["git", "add", "-A"], check=True)
    subprocess.run(["git", "commit", "-m", message], check=True)
    print(f"[OK] Commit: {message}")
    
    # Push si se requiere
    if push:
        subprocess.run(["git", "push"], check=True)
        print("[OK] Push completado")

if __name__ == "__main__":
    main()
```

## Reglas

1. **Siempre commit**: Después de cada función implementada
2. **Siempre push**: Subir cambios al remoto inmediatamente
3. **Mensaje claro**: Descripción corta del cambio
4. **Conventional commits**: Formato `<tipo>: <descripción>`

## Configuración

El script está en: `scripts/auto_commit.py`

Para ejecutar manualmente:
```bash
python scripts/auto_commit.py
```

Para ejecutar con mensaje personalizado:
```bash
python scripts/auto_commit.py "feat: Nueva funcionalidad"
```

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/auto-commit
