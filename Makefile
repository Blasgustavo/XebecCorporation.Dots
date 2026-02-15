# Makefile para automatizar commits

.PHONY: commit save

# Commit automático con detección de tipo
commit:
	@python scripts/auto_commit.py $(MESSAGE)

# Alias para guardar cambios
save: commit

# Commit con mensaje personalizado
commit-msg:
	@python scripts/auto_commit.py "$(MESSAGE)"

# Ver estado
status:
	@git status

# Ver último commit
log:
	@git log -1 --oneline
