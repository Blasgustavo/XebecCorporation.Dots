#!/usr/bin/env python3
"""
Script de commit commits
Detecta el tipo de cambio y genera el mensaje apropiado
Soporta args: automático con conventional auto-commit [--no-push] [mensaje]
"""

import subprocess
import sys
import os
from pathlib import Path

# Tipos de commit con emojis
COMMIT_TYPES = {
    "feat": {"emoji": "✨", "description": "Nueva funcionalidad"},
    "fix": {"emoji": "🐛", "description": "Corrección de bug"},
    "docs": {"emoji": "📚", "description": "Documentación"},
    "chore": {"emoji": "🔧", "description": "Tareas de mantenimiento"},
    "refactor": {"emoji": "♻️", "description": "Refactorización"},
    "test": {"emoji": "✅", "description": "Tests"},
    "style": {"emoji": "💎", "description": "Estilo/formato"},
    "perf": {"emoji": "⚡", "description": "Rendimiento"},
    "ci": {"emoji": "👷", "description": "CI/CD"},
    "build": {"emoji": "📦", "description": "Build/dependencias"},
    "revert": {"emoji": "⏪", "description": "Revertir"},
}

def get_git_changes():
    """Obtiene los archivos cambiados"""
    result = subprocess.run(
        ["git", "status", "--porcelain"],
        capture_output=True,
        text=True
    )
    return result.stdout.strip()

def detect_change_type(changes):
    """Detecta el tipo de cambio basándose en los archivos"""
    changes_lower = changes.lower()
    
    # Detección por archivos
    if "docs/" in changes or "readme" in changes_lower:
        return "docs", "Actualizar documentación"
    
    if "cmd/" in changes or "internal/ui" in changes or "internal/os" in changes or "internal/actions" in changes:
        return "feat", "Agregar funcionalidad al CLI"
    
    if ".go" in changes and ("test" in changes_lower or "_test.go" in changes_lower):
        return "test", "Agregar tests"
    
    if "go.mod" in changes or "go.sum" in changes or "package" in changes_lower:
        return "build", "Actualizar dependencias"
    
    if ".gitignore" in changes or "scripts/" in changes:
        return "chore", "Actualizar configuración"
    
    # Default
    return "feat", "Actualizar proyecto"

def generate_commit_message(commit_type, description):
    """Genera el mensaje de commit con formato conventional"""
    emoji = COMMIT_TYPES[commit_type]["emoji"]
    return f"{emoji} {commit_type}: {description}"

def commit(message=None, do_push=True):
    """Hace el commit"""
    # Verificar si hay cambios
    changes = get_git_changes()
    if not changes:
        print("[INFO] No hay cambios para commitear")
        return False
    
    # Si no hay mensaje, detectar tipo
    if not message:
        commit_type, desc = detect_change_type(changes)
        message = generate_commit_message(commit_type, desc)
    
    # Agregar cambios
    subprocess.run(["git", "add", "-A"], check=True)
    
    # Hacer commit
    result = subprocess.run(
        ["git", "commit", "-m", message],
        capture_output=True,
        text=True
    )
    
    if result.returncode == 0:
        # Mostrar resultado
        result = subprocess.run(
            ["git", "log", "-1", "--oneline"],
            capture_output=True,
            text=True
        )
        print(f"[OK] Commit: {result.stdout.strip()}")
        
        # Push si se requiere
        if do_push:
            push_result = subprocess.run(
                ["git", "push"],
                capture_output=True,
                text=True
            )
            if push_result.returncode == 0:
                print("[OK] Push completado")
            else:
                print(f"[WARN] Push falló: {push_result.stderr}")
        
        return True
    else:
        print(f"[ERROR] {result.stderr}")
        return False

def main():
    message = None
    do_push = True
    
    # Parsear argumentos
    for arg in sys.argv[1:]:
        if arg == "--no-push":
            do_push = False
        elif not arg.startswith("-"):
            message = arg
    
    commit(message, do_push)

if __name__ == "__main__":
    main()
