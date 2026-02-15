#!/bin/bash
# Script para hacer commits automáticos con conventional commits
# Uso: ./commit.sh "mensaje del commit"

# Colores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Verificar si hay cambios
if [ -z "$(git status --porcelain)" ]; then
    echo -e "${YELLOW}No hay cambios para commitear${NC}"
    exit 0
fi

# Si se proporciona mensaje, usarlo
if [ -n "$1" ]; then
    COMMIT_MSG="$1"
else
    echo -e "${YELLOW}Proporciona un mensaje de commit o usa --auto para detección automática${NC}"
    exit 1
fi

# Agregar todos los cambios
git add -A

# Hacer commit
git commit -m "$COMMIT_MSG"

# Mostrar resultado
echo -e "${GREEN}Commit realizado:${NC}"
git log -1 --oneline
