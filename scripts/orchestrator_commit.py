#!/usr/bin/env python3
"""
Orquestador de commits - Automatiza el proceso de commit
Invoca el skill de conventional commits automáticamente
"""

import subprocess
import sys
import os
from pathlib import Path

# Colores (sin emojis para Windows compatibility)
GREEN = '[OK]'
YELLOW = '[INFO]'
BLUE = '[XEBEC]'
NC = ''

# Tipos de commit (del skill) - sin emojis para Windows
COMMIT_TYPES = {
    "feat": {"emoji": "[FEAT]", "description": "Nueva funcionalidad", "paths": ["cmd/", "internal/", "main.go"]},
    "fix": {"emoji": "[FIX]", "description": "Bug fix", "paths": ["fix", "bug", "error", "issue"]},
    "docs": {"emoji": "[DOCS]", "description": "Documentacion", "paths": ["docs/", "README", "readme"]},
    "chore": {"emoji": "[CHORE]", "description": "Mantenimiento", "paths": ["scripts/", "Makefile", ".gitignore"]},
    "build": {"emoji": "[BUILD]", "description": "Build/dependencias", "paths": ["go.mod", "go.sum", "vendor/"]},
    "test": {"emoji": "[TEST]", "description": "Tests", "paths": ["_test.go", "test"]},
    "refactor": {"emoji": "[REFACTOR]", "description": "Refactorizacion", "paths": ["refactor"]},
    "style": {"emoji": "[STYLE]", "description": "Estilo", "paths": ["format", "style"]},
    "perf": {"emoji": "[PERF]", "description": "Rendimiento", "paths": ["perf", "performance", "optimize"]},
    "ci": {"emoji": "[CI]", "description": "CI/CD", "paths": [".github/", "ci", "workflow"]},
}

def print_banner():
    """Muestra el banner del orquestador"""
    print("  _                _   _             _   ____  ")
    print(" | |    ___  __ _| |_| |_ ___     __| | / _ \\ ")
    print(" | |   / _ \\/ _' |  _| __/ __|   / _  || | | |")
    print(" | |__|  __/ (_| | | | |_\\__ \\  | (_| || |_| |")
    print(" |_____\\___|\\__,_|_|  \\__|___/   \\__,_|\\___/ ")
    print()
    print("  Orquestador de Commits - XEBEC CORPORATION")
    print()

def get_changes():
    """Obtiene los cambios pendientes"""
    result = subprocess.run(
        ["git", "status", "--porcelain"],
        capture_output=True,
        text=True
    )
    return result.stdout.strip()

def detect_commit_type(changes):
    """Detecta el tipo de commit basado en los archivos cambiados"""
    changes_lower = changes.lower()
    
    # Orden de prioridad (docs primero porque es específico)
    if any(p in changes_lower for p in COMMIT_TYPES["docs"]["paths"]):
        return "docs"
    
    if any(p in changes_lower for p in COMMIT_TYPES["feat"]["paths"]):
        return "feat"
    
    if any(p in changes_lower for p in COMMIT_TYPES["fix"]["paths"]):
        return "fix"
    
    if any(p in changes_lower for p in COMMIT_TYPES["build"]["paths"]):
        return "build"
    
    if any(p in changes_lower for p in COMMIT_TYPES["test"]["paths"]):
        return "test"
    
    if any(p in changes_lower for p in COMMIT_TYPES["chore"]["paths"]):
        return "chore"
    
    # Default
    return "feat"

def generate_message(commit_type, description=""):
    """Genera el mensaje de commit con el formato del skill"""
    emoji = COMMIT_TYPES[commit_type]["emoji"]
    desc = description or COMMIT_TYPES[commit_type]["description"]
    return f"{emoji} {commit_type}: {desc}"

def commit_with_skill(message=None):
    """Hace el commit invocando el skill de conventional commits"""
    changes = get_changes()
    
    if not changes:
        print("[INFO] No hay cambios para commitear")
        return False
    
    # Detectar tipo si no hay mensaje
    if not message:
        commit_type = detect_commit_type(changes)
        message = generate_message(commit_type)
        print(f"[DETECTED] Tipo: {commit_type} ({COMMIT_TYPES[commit_type]['description']})")
    
    # Mostrar cambio detectado
    print(f"\n[COMMIT] {message}")
    
    # Agregar y commitear
    subprocess.run(["git", "add", "-A"], check=True)
    result = subprocess.run(
        ["git", "commit", "-m", message],
        capture_output=True,
        text=True
    )
    
    if result.returncode == 0:
        result = subprocess.run(
            ["git", "log", "-1", "--oneline"],
            capture_output=True,
            text=True
        )
        print(f"\n[OK] Commit completado:")
        print(f"      {result.stdout.strip()}")
        return True
    else:
        print(f"\n[ERROR] {result.stderr}")
        return False

def main():
    print_banner()
    
    # Obtener mensaje si se proporciona
    message = None
    if len(sys.argv) > 1:
        message = " ".join(sys.argv[1:])
    
    success = commit_with_skill(message)
    sys.exit(0 if success else 1)

if __name__ == "__main__":
    main()
