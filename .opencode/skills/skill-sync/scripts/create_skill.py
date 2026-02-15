#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Crea estructura básica de un nuevo skill con integración opcional de agentes.

Usage:
    python create_skill.py <nombre> [descripcion]
    python create_skill.py <nombre> [descripcion] --integrate
    python create_skill.py <nombre> --agents terminal-config shell-config
"""

import argparse
import os
import sys
from pathlib import Path

# Configurar UTF-8 para stdout en Windows
if sys.platform == 'win32':
    import io
    sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8', errors='replace')

# Símbolos compatibles con Windows
SYMBOL_OK = "[OK]"
SYMBOL_WARN = "[!]"


# Agentes disponibles del proyecto
AVAILABLE_AGENTS = [
    'terminal-config',
    'shell-config',
    'tools-install',
    'explore',
    'analysis',
    'generator'
]


def create_skill(name: str, description: str = "", integrate: bool = False, agents: list[str] = None):
    """Crea un nuevo skill con estructura básica."""
    base = Path(".opencode/skills")
    skill_dir = base / name
    
    if skill_dir.exists():
        print(f"{SYMBOL_WARN} El skill '{name}' ya existe en {skill_dir}")
        response = input("¿Deseas sobrescribirlo? (s/N): ").strip().lower()
        if response != 's':
            print("Cancelado.")
            return None
    
    skill_dir.mkdir(parents=True, exist_ok=True)
    
    # Validar agentes si se proporcionan
    valid_agents = []
    if agents:
        for agent in agents:
            if agent in AVAILABLE_AGENTS:
                valid_agents.append(agent)
            else:
                print(f"{SYMBOL_WARN} Agente desconocido: {agent}")
                print(f"    Agentes disponibles: {', '.join(AVAILABLE_AGENTS)}")
    
    # Construir frontmatter
    frontmatter = f"""---
name: {name}
description: {description or "Skill generado automáticamente."}
"""
    
    if integrate or valid_agents:
        frontmatter += "agents:\n"
        if valid_agents:
            for agent in valid_agents:
                frontmatter += f"  - {agent}\n"
        else:
            # Agregar todos los agentes disponibles si es integración pero no se especificaron
            for agent in AVAILABLE_AGENTS[:3]:  # Por defecto, agregar los 3 primeros
                frontmatter += f"  - {agent}\n"
    
    frontmatter += "---\n"
    
    content = f"""{frontmatter}
# {name.title()}

{description or "Skill generado automáticamente."}

## Cuándo Usar Este Skill

Usa este skill cuando:
- [Describe cuándo se debe usar este skill]

## Instrucciones

[Agrega las instrucciones específicas para este skill]

## Ejemplos

### Ejemplo 1

```bash
# Ejemplo de uso
```

## Integración con Agentes

"""
    
    # Agregar información de agentes si aplica
    if integrate or valid_agents:
        agent_list = valid_agents if valid_agents else AVAILABLE_AGENTS[:3]
        content += f"""
Este skill está configurado para trabajar con los siguientes agentes:

| Agente | Descripción |
|--------|-------------|
"""
        for agent in agent_list:
            desc = {
                'terminal-config': 'Configuración de terminal (Alacritty)',
                'shell-config': 'Configuración de shell (Nushell, Starship)',
                'tools-install': 'Instalación de herramientas',
                'explore': 'Exploración de código',
                'analysis': 'Análisis de configuraciones',
                'generator': 'Generación de archivos'
            }.get(agent, agent)
            content += f"| `{agent}` | {desc} |\n"
    
    content += f"""

---

Base directory for this skill: file:///C:/Users/qty94/Documents/XebecCorporation.Dots/.opencode/skills/{name}
"""
    
    (skill_dir / "SKILL.md").write_text(content, encoding="utf-8")
    print(f"{SYMBOL_OK} Skill '{name}' creado en {skill_dir}")
    
    # Crear estructura de directorios opcionales
    for subdir in ['scripts', 'references', 'assets']:
        (skill_dir / subdir).mkdir(exist_ok=True)
    
    # Crear script de ejemplo si hay integración
    if integrate or valid_agents:
        example_script = skill_dir / "scripts" / "example.py"
        example_script.write_text('''#!/usr/bin/env python3
"""Script de ejemplo para el skill."""

def main():
    print("Skill ejecutado correctamente")

if __name__ == "__main__":
    main()
''', encoding="utf-8")
        print(f"{SYMBOL_OK} Script de ejemplo creado en {example_script}")
    
    # Información de próximo paso
    print(f"""
[NEXT STEPS] Próximos pasos:
   1. Edita el SKILL.md con las instrucciones específicas
   2. Ejecuta sincronización: python .opencode/skills/skill-sync/scripts/sync_skills.py
   3. Reinicia el contexto de OpenCode
""")
    
    if integrate or valid_agents:
        print(f"[LINK] Integración con agentes configurada")
    
    return skill_dir


def main():
    parser = argparse.ArgumentParser(
        description='Crea un nuevo skill para OpenCode'
    )
    parser.add_argument(
        'name',
        help='Nombre del skill (usa minúsculas y guiones)'
    )
    parser.add_argument(
        'description',
        nargs='?',
        default='',
        help='Descripción del skill'
    )
    parser.add_argument(
        '--integrate', '-i',
        action='store_true',
        help='Configurar integración con agentes del proyecto'
    )
    parser.add_argument(
        '--agents', '-a',
        nargs='+',
        help='Lista de agentes con los que integrar (espacio separado)'
    )
    
    args = parser.parse_args()
    
    name = args.name.lower().replace(" ", "-")
    
    create_skill(
        name, 
        args.description, 
        args.integrate, 
        args.agents
    )


if __name__ == "__main__":
    main()
