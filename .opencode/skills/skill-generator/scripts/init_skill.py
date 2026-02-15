#!/usr/bin/env python3
"""
Skill Generator - Init Script

Crea la estructura básica de un nuevo skill.
"""

import argparse
import os
import re
from pathlib import Path


def validate_skill_name(name: str) -> bool:
    """Valida que el nombre del skill sea válido."""
    if not name:
        return False
    if len(name) > 64:
        return False
    if '--' in name:
        return False
    if name.startswith('-') or name.endswith('-'):
        return False
    pattern = r'^[a-z0-9]+(-[a-z0-9]+)*$'
    return bool(re.match(pattern, name))


def create_skill_structure(skill_name: str, output_path: Path) -> None:
    """Crea la estructura de directorios y archivos para un nuevo skill."""
    
    skill_path = output_path / skill_name
    skill_path.mkdir(parents=True, exist_ok=True)
    
    # Crear directorios
    (skill_path / "scripts").mkdir(exist_ok=True)
    (skill_path / "references").mkdir(exist_ok=True)
    (skill_path / "assets").mkdir(exist_ok=True)
    
    # Crear SKILL.md base siguiendo estándar Anthropic
    skill_md = f"""---
name: {skill_name}
description: Descripción breve del skill (máx 200 caracteres). Explica qué hace y cuándo usarlo.
---

# {skill_name.replace('-', ' ').title()}

Este skill proporciona... [describe brevemente la funcionalidad]

## Cuándo Usar

Describe cuándo Claude debe activar este skill. Por ejemplo:
- "Usa este skill cuando el usuario necesite..."

## Instrucciones

1. **Paso 1**: [Descripción del primer paso]
2. **Paso 2**: [Descripción del segundo paso]
3. **Paso 3**: [Descripción del tercer paso]

## Ejemplos

### Ejemplo 1
Input: "..."
Output: "..."

### Ejemplo 2
Input: "..."
Output: "..."

## Recursos

- [Referencia](references/archivo.md)
- [Script](scripts/script.py)
"""
    
    (skill_path / "SKILL.md").write_text(skill_md, encoding='utf-8')
    
    # Crear scripts de ejemplo
    script_example = """#!/usr/bin/env python3
# Example script - delete or customize

def main():
    print("Hello from skill!")

if __name__ == "__main__":
    main()
"""
    (skill_path / "scripts" / "example.py").write_text(script_example, encoding='utf-8')
    
    # Crear reference de ejemplo
    reference_example = """# Referencia de Ejemplo

Agrega documentación de referencia aquí...

## Sección 1

Contenido...

## Sección 2

Contenido...
"""
    (skill_path / "references" / "example.md").write_text(reference_example, encoding='utf-8')
    
    # Crear .gitkeep en assets
    (skill_path / "assets" / ".gitkeep").touch()
    
    print(f"[OK] Skill '{skill_name}' creado en: {skill_path}")
    print("\nEstructura creada:")
    print(f"  {skill_name}/")
    print(f"  +-- SKILL.md")
    print(f"  +-- scripts/")
    print(f"  |   +-- example.py")
    print(f"  +-- references/")
    print(f"  |   +-- example.md")
    print(f"  +-- assets/")
    print("\nProximos pasos:")
    print(f"  1. Edita {skill_name}/SKILL.md con tu contenido")
    print(f"  2. Personaliza o elimina scripts/references/assets")
    print(f"  3. Ejecuta package_skill.py para validar")


def main():
    parser = argparse.ArgumentParser(
        description="Inicializa un nuevo skill con estructura básica"
    )
    parser.add_argument(
        "name",
        help="Nombre del skill (formato: lowercase con guiones)"
    )
    parser.add_argument(
        "--path",
        type=Path,
        default=Path("."),
        help="Directorio de salida (default: actual)"
    )
    
    args = parser.parse_args()
    
    if not validate_skill_name(args.name):
        print(f"Error: Nombre '{args.name}' inválido.")
        print("Requisitos:")
        print("  - Solo letras minúsculas, números y guiones")
        print("  - No puede tener '--'")
        print("  - No puede empezar o terminar con '-'")
        print("  - Máximo 64 caracteres")
        return 1
    
    create_skill_structure(args.name, args.path)
    return 0


if __name__ == "__main__":
    exit(main())
