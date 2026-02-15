#!/usr/bin/env python3
"""Crea estructura básica de un nuevo skill."""

import os
import sys
from pathlib import Path

def create_skill(name: str, description: str = ""):
    base = Path(".opencode/skills")
    skill_dir = base / name
    skill_dir.mkdir(parents=True, exist_ok=True)
    
    content = f"""# Skill: {name}

# {name.title()}

{description or "Skill generado automáticamente."}

## Contenido

- [Descripción](#descripción)
- [Uso](#uso)
- [Ejemplos](#ejemplos)

## Uso

```bash
# Ejemplo de uso
```

## Ejemplos

### Ejemplo 1

Descripción del ejemplo.

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/{name}
"""
    
    (skill_dir / "SKILL.md").write_text(content, encoding="utf-8")
    print(f"✅ Skill '{name}' creado en {skill_dir}")
    return skill_dir

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Uso: create_skill.py <nombre> [descripción]")
        sys.exit(1)
    
    name = sys.argv[1].lower().replace(" ", "-")
    desc = sys.argv[2] if len(sys.argv) > 2 else ""
    create_skill(name, desc)
