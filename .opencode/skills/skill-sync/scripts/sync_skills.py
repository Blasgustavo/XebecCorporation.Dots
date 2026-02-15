#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Script de sincronización para skills de OpenCode.

Este script:
- Valida la estructura de todos los skills
- Verifica frontmatter de cada SKILL.md
- Genera reporte de estado
- Sugiere correcciones si hay errores
- Integra con el orquestador y agentes del proyecto
"""

import argparse
import json
import os
import re
import sys
from pathlib import Path
from typing import Any

# Configurar UTF-8 para stdout en Windows
if sys.platform == 'win32':
    import io
    sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8', errors='replace')

# Símbolos compatibles con Windows
SYMBOL_OK = "[OK]"
SYMBOL_WARN = "[!]"
SYMBOL_ERROR = "[X]"
SYMBOL_INFO = "[i]"


# Colores para terminal (solo si soporta ANSI)
class Colors:
    try:
        GREEN = '\033[92m'
        YELLOW = '\033[93m'
        RED = '\033[91m'
        BLUE = '\033[94m'
        RESET = '\033[0m'
        BOLD = '\033[1m'
        ENABLED = True
    except Exception:
        GREEN = YELLOW = RED = BLUE = RESET = BOLD = ''
        ENABLED = False


def print_success(msg: str):
    prefix = f"{Colors.GREEN}{SYMBOL_OK}{Colors.RESET} " if Colors.ENABLED else f"{SYMBOL_OK} "
    print(f"{prefix}{msg}")


def print_warning(msg: str):
    prefix = f"{Colors.YELLOW}{SYMBOL_WARN}{Colors.RESET} " if Colors.ENABLED else f"{SYMBOL_WARN} "
    print(f"{prefix}{msg}")


def print_error(msg: str):
    prefix = f"{Colors.RED}{SYMBOL_ERROR}{Colors.RESET} " if Colors.ENABLED else f"{SYMBOL_ERROR} "
    print(f"{prefix}{msg}")


def print_info(msg: str):
    prefix = f"{Colors.BLUE}{SYMBOL_INFO}{Colors.RESET} " if Colors.ENABLED else f"{SYMBOL_INFO} "
    print(f"{prefix}{msg}")


def parse_frontmatter(content: str) -> tuple[dict[str, str], str]:
    """Parsea el frontmatter YAML de un archivo markdown."""
    pattern = r'^---\s*\n(.*?)\n---\s*\n(.*)$'
    match = re.match(pattern, content, re.DOTALL)
    
    if not match:
        return {}, content
    
    yaml_content = match.group(1)
    body = match.group(2)
    
    # Parse simple YAML (solo clave: valor)
    frontmatter = {}
    for line in yaml_content.split('\n'):
        if ':' in line:
            key, value = line.split(':', 1)
            frontmatter[key.strip()] = value.strip()
    
    return frontmatter, body


def validate_skill(skill_path: Path, verbose: bool = False) -> dict[str, Any]:
    """Valida un skill y retorna el resultado."""
    result = {
        'name': skill_path.name,
        'path': str(skill_path),
        'valid': True,
        'errors': [],
        'warnings': [],
        'info': {}
    }
    
    # Verificar que existe SKILL.md
    skill_md = skill_path / 'SKILL.md'
    if not skill_md.exists():
        result['valid'] = False
        result['errors'].append('Falta SKILL.md')
        return result
    
    # Verificar frontmatter
    try:
        content = skill_md.read_text(encoding='utf-8')
        frontmatter, body = parse_frontmatter(content)
        
        if not frontmatter:
            result['valid'] = False
            result['errors'].append('Falta frontmatter (---)')
        else:
            # Verificar campos requeridos
            if 'name' not in frontmatter:
                result['valid'] = False
                result['errors'].append('Falta campo "name" en frontmatter')
            elif frontmatter['name'] != skill_path.name:
                result['warnings'].append(
                    f'El nombre en frontmatter "{frontmatter["name"]}" '
                    f'no coincide con el directorio "{skill_path.name}"'
                )
            
            if 'description' not in frontmatter:
                result['warnings'].append('Falta campo "description" en frontmatter')
            
            # Verificar integración con agentes
            if 'agents' in frontmatter:
                agents = frontmatter['agents']
                result['info']['agents'] = agents
                if verbose:
                    print_info(f"  Agentes registrados: {agents}")
            
            result['info']['frontmatter'] = frontmatter
    
    except Exception as e:
        result['valid'] = False
        result['errors'].append(f'Error al leer SKILL.md: {e}')
    
    # Verificar estructura opcional
    for subdir in ['scripts', 'references', 'assets']:
        subdir_path = skill_path / subdir
        if subdir_path.exists():
            result['info'][subdir] = True
    
    return result


def get_agents_from_project() -> list[str]:
    """Obtiene la lista de agentes disponibles del proyecto."""
    # Definición de agentes del proyecto XebecCorporation.Dots
    return [
        'terminal-config',
        'shell-config', 
        'tools-install',
        'explore',
        'analysis',
        'generator'
    ]


def sync_skills(base_path: Path, verbose: bool = False, fix: bool = False) -> int:
    """Sincroniza todos los skills en el directorio base."""
    skills_dir = base_path / '.opencode' / 'skills'
    
    if not skills_dir.exists():
        print_error(f"Directorio de skills no encontrado: {skills_dir}")
        return 1
    
    # Obtener todos los directorios de skills
    skill_dirs = [d for d in skills_dir.iterdir() if d.is_dir()]
    
    if not skill_dirs:
        print_warning("No se encontraron skills en el directorio")
        return 0
    
    print(f"{Colors.BOLD}Sincronizando {len(skill_dirs)} skills...{Colors.RESET}\n")
    
    valid_count = 0
    error_count = 0
    warning_count = 0
    
    all_results = []
    
    for skill_dir in sorted(skill_dirs):
        if verbose:
            print(f"{Colors.BOLD}Verificando: {skill_dir.name}{Colors.RESET}")
        
        result = validate_skill(skill_dir, verbose)
        all_results.append(result)
        
        if result['valid']:
            valid_count += 1
            print_success(f"  {skill_dir.name}")
            
            if result['warnings']:
                for warning in result['warnings']:
                    print_warning(f"    {warning}")
                    warning_count += 1
        else:
            error_count += 1
            print_error(f"  {skill_dir.name}")
            for error in result['errors']:
                print_error(f"    {error}")
    
    # Resumen
    print(f"\n{Colors.BOLD}=== Resumen ==={Colors.RESET}")
    print(f"Total de skills: {len(skill_dirs)}")
    print(f"{Colors.GREEN}Válidos: {valid_count}{Colors.RESET}")
    
    if warning_count > 0:
        print(f"{Colors.YELLOW}Advertencias: {warning_count}{Colors.RESET}")
    
    if error_count > 0:
        print(f"{Colors.RED}Errores: {error_count}{Colors.RESET}")
    
    # Información de integración con agentes
    print(f"\n{Colors.BOLD}=== Agentes del Proyecto ==={Colors.RESET}")
    agents = get_agents_from_project()
    for agent in agents:
        print(f"  • {agent}")
    
    # Skills que integran con agentes
    skills_with_agents = [r for r in all_results if r['info'].get('agents')]
    if skills_with_agents:
        print(f"\n{Colors.BOLD}=== Skills con Integración ==={Colors.RESET}")
        for r in skills_with_agents:
            agents_list = r['info']['agents']
            print(f"  • {r['name']}: {', '.join(agents_list)}")
    
    # Guardar reporte
    report_path = base_path / '.opencode' / 'skills' / 'sync_report.json'
    report_data = {
        'total': len(skill_dirs),
        'valid': valid_count,
        'errors': error_count,
        'warnings': warning_count,
        'skills': all_results,
        'available_agents': agents
    }
    report_path.write_text(json.dumps(report_data, indent=2, ensure_ascii=False), encoding='utf-8')
    print(f"\nReporte guardado en: {report_path}")
    
    return 0 if error_count == 0 else 1


def main():
    parser = argparse.ArgumentParser(
        description='Sincroniza y valida skills de OpenCode'
    )
    parser.add_argument(
        '--verbose', '-v',
        action='store_true',
        help='Muestra detalles de cada skill'
    )
    parser.add_argument(
        '--fix',
        action='store_true',
        help='Intenta corregir errores automáticamente'
    )
    parser.add_argument(
        '--path',
        type=str,
        default='.',
        help='Ruta base del proyecto (default: .)'
    )
    
    args = parser.parse_args()
    
    base_path = Path(args.path).resolve()
    exit_code = sync_skills(base_path, args.verbose, args.fix)
    
    sys.exit(exit_code)


if __name__ == '__main__':
    main()
