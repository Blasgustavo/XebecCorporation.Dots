---
name: skill-commits
description: Conventional Commits con Iconos para XebecCorporation.Dots
---

# Skill: skill-commits

# Conventional Commits con Iconos

Este skill proporciona la convención de commits con emojis para XebecCorporation.Dots.

## Tipos de Commit

| Tipo | Icono | Descripción |
|------|-------|-------------|
| `feat` | ✨ | Nueva funcionalidad |
| `fix` | 🐛 | Corrección de bug |
| `docs` | 📚 | Documentación |
| `chore` | 🔧 | Tareas de mantenimiento |
| `refactor` | ♻️ | Refactorización de código |
| `test` | ✅ | Tests unitarios o de integración |
| `style` | 💎 | Cambios de formato/código style |
| `perf` | ⚡ | Mejora de rendimiento |
| `ci` | 👷 | Cambios en CI/CD |
| `build` | 📦 | Cambios en build/dependencias |
| `revert` | ⏪ | Revertir commit anterior |

## Formato de Commit

```
<icono> <tipo>: <descripción>

[opcional: cuerpo del commit]

[opcional: footer]
```

### Reglas

1. **Tipo**: siempre en minúsculas
2. **Descripción**: máximo 50 caracteres, sin punto final
3. **Cuerpo**:separado por línea en blanco, máximo 72 caracteres por línea
4. **Footer**: para Breaking Changes o issues relacionados

## Ejemplos

### Feature
```
✨ feat: agregar comando install para winget
```

### Bug Fix
```
🐛 fix: corregir detección de SO en Windows 11
```

### Con cuerpo
```
📚 docs: actualizar README con nuevas dependencias

Se agregaron las secciones de:
- Requisitos del sistema
- Instalación en Linux
- Configuración de Alacritty
```

### Breaking Change
```
⚡ perf: mejorar velocidad de detección de paquetes

BREAKING CHANGE: El flag --quiet ahora es --silent
```

## Reglas de Uso

1. Usa un solo tipo por commit
2. Commits atómicos: una cambio = un commit
3. describe el "qué" y "por qué", no el "cómo"
4. Commits en inglés (para el mensaje), español (para contexto adicional)

## Integración con Git

### Commitizen (opcional)
```bash
npx cz commit
```

### Git hooks
El proyecto incluye pre-commit hooks que validan el formato.

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/skill-commits
