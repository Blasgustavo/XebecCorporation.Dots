#!/usr/bin/env python3
"""
Skill Generator - Package Script

Valida y empaqueta un skill en un archivo .skill (zip)
"""

import argparse
import os
import re
import sys
import zipfile
from pathlib import Path
from datetime import datetime


def validate_frontmatter(content: str) -> tuple[bool, str]:
    """Valida el frontmatter YAML del SKILL.md."""
    if not content.startswith('---'):
        return False, "SKILL.md debe comenzar con ---"
    
    lines = content.split('\n')
    end_index = -1
    for i, line in enumerate(lines[1:], 1):
        if line.strip() == '---':
            end_index = i
            break
    
    if end_index == -1:
        return False, "No se encontró cierre de frontmatter ---"
    
    frontmatter = '\n'.join(lines[1:end_index])
    
    # Validar name
    if 'name:' not in frontmatter:
        return False, "Frontmatter debe contener 'name'"
    
    # Validar description
    if 'description:' not in frontmatter:
        return False, "Frontmatter debe contener 'description'"
    
    return True, "OK"


def validate_skill_name(name: str) -> tuple[bool, str]:
    """Valida el nombre del skill."""
    if not name:
        return False, "Nombre vacío"
    
    if len(name) > 64:
        return False, "Nombre demasiado largo (máx 64 caracteres)"
    
    if '--' in name:
        return False, "Nombre no puede contener '--'"
    
    if name.startswith('-') or name.endswith('-'):
        return False, "Nombre no puede empezar o terminar con '-'"
    
    pattern = r'^[a-z0-9]+(-[a-z0-9]+)*$'
    if not re.match(pattern, name):
        return False, "Solo lowercase, números y guiones simples"
    
    return True, "OK"


def validate_directory_structure(skill_path: Path) -> tuple[bool, list[str]]:
    """Valida la estructura de directorios del skill."""
    errors = []
    
    # SKILL.md requerido
    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        errors.append("Falta SKILL.md requerido")
    
    # Frontmatter de SKILL.md
    if skill_md.exists():
        content = skill_md.read_text(encoding='utf-8')
        valid, msg = validate_frontmatter(content)
        if not valid:
            errors.append(f"Error en SKILL.md frontmatter: {msg}")
        
        # Extraer y validar nombre
        for line in content.split('\n'):
            if line.startswith('name:'):
                name = line.split(':', 1)[1].strip()
                valid, msg = validate_skill_name(name)
                if not valid:
                    errors.append(f"Nombre inválido: {msg}")
                break
    
    # Validar subdirectorios opcionales
    optional_dirs = ['scripts', 'references', 'assets']
    for dir_name in optional_dirs:
        dir_path = skill_path / dir_name
        if dir_path.exists() and not any(dir_path.iterdir()):
            # Directorio vacío es OK, warn si es el único contenido
            pass
    
    return len(errors) == 0, errors


def validate_skill(skill_path: Path) -> tuple[bool, list[str]]:
    """Valida un skill completo."""
    errors = []
    
    if not skill_path.exists():
        return False, [f"Skill path no existe: {skill_path}"]
    
    if not skill_path.is_dir():
        return False, [f"No es un directorio: {skill_path}"]
    
    valid, dir_errors = validate_directory_structure(skill_path)
    errors.extend(dir_errors)
    
    # Validar longitud del description
    skill_md = skill_path / "SKILL.md"
    if skill_md.exists():
        content = skill_md.read_text(encoding='utf-8')
        
        # Extraer description
        desc_found = False
        desc_length = 0
        for line in content.split('\n'):
            if line.startswith('description:'):
                desc_found = True
                desc_length = len(line.split(':', 1)[1].strip())
            elif desc_found and line.strip().startswith('-'):
                # Multiline description
                desc_length += len(line.strip())
            elif desc_found and line.strip() and not line.startswith(' '):
                break
        
        if desc_length > 1024:
            errors.append(f"Description muy largo ({desc_length} chars, máx 1024)")
    
    return len(errors) == 0, errors


def package_skill(skill_path: Path, output_dir: Path = None) -> bool:
    """Empaqueta el skill en un archivo .skill (zip)."""
    
    skill_name = skill_path.name
    
    if output_dir is None:
        output_dir = skill_path.parent
    
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)
    
    output_file = output_dir / f"{skill_name}.skill"
    
    # Si ya existe, agregar timestamp
    if output_file.exists():
        timestamp = datetime.now().strftime("%Y%m%d%H%M%S")
        output_file = output_dir / f"{skill_name}_{timestamp}.skill"
    
    with zipfile.ZipFile(output_file, 'w', zipfile.ZIP_DEFLATED) as zf:
        for file_path in skill_path.rglob('*'):
            if file_path.is_file() and file_path.name != '.gitkeep':
                arcname = file_path.relative_to(skill_path)
                zf.write(file_path, arcname)
    
    print(f"✓ Skill empaquetado: {output_file}")
    return True


def main():
    parser = argparse.ArgumentParser(
        description="Valida y empaqueta un skill"
    )
    parser.add_argument(
        "path",
        type=Path,
        help="Ruta al directorio del skill"
    )
    parser.add_argument(
        "output",
        type=Path,
        nargs='?',
        default=None,
        help="Directorio de salida (opcional)"
    )
    
    args = parser.parse_args()
    
    skill_path = args.path.resolve()
    
    print(f"Validando skill: {skill_path}")
    print("-" * 50)
    
    valid, errors = validate_skill(skill_path)
    
    if not valid:
        print("✗ Errores encontrados:")
        for error in errors:
            print(f"  - {error}")
        print("\n✗ Empaquetado cancelado.")
        return 1
    
    print("✓ Validación exitosa")
    print("-" * 50)
    
    # Empaquetar
    output_dir = args.output.resolve() if args.output else None
    if package_skill(skill_path, output_dir):
        print("-" * 50)
        print("✓ Proceso completado")
        return 0
    
    return 1


if __name__ == "__main__":
    sys.exit(main())
