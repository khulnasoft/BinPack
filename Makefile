OWNER = khulnasoft
PROJECT = binpack

TOOL_DIR = .tool
BINPACK = $(TOOL_DIR)/binpack
TASK = $(TOOL_DIR)/task

.DEFAULT_GOAL := make-default

## Bootstrapping targets #################################

$(BINPACK):
	@-mkdir $(TOOL_DIR)
# we don't have a release of binpack yet, so build off of the current branch
# curl -sSfL https://raw.githubusercontent.com/$(OWNER)/$(PROJECT)/main/install.sh | sh -s -- -b $(TOOL_DIR)
	go build -o $(TOOL_DIR)/$(PROJECT) ./cmd/$(PROJECT)

.PHONY: task
$(TASK) task: $(BINPACK)
	$(BINPACK) install task

.PHONY: ci-bootstrap-go
ci-bootstrap-go:
	go mod download

.PHONY: ci-bootstrap-tools
ci-bootstrap-tools: $(BINPACK)
	$(BINPACK) install -vvv

# this is a bootstrapping catch-all, where if the target doesn't exist, we'll ensure the tools are installed and then try again
%:
	make $(TASK)
	$(TASK) $@

## Shim targets #################################

.PHONY: make-default
make-default: $(TASK)
	@# run the default task in the taskfile
	@$(TASK)

# for those of us that can't seem to kick the habit of typing `make ...` lets wrap the superior `task` tool
TASKS := $(shell bash -c "$(TASK) -l | grep '^\* ' | cut -d' ' -f2 | tr -d ':' | tr '\n' ' '" ) $(shell bash -c "$(TASK) -l | grep 'aliases:' | cut -d ':' -f 3 | tr '\n' ' ' | tr -d ','")

.PHONY: $(TASKS)
$(TASKS): $(TASK)
	@$(TASK) $@

help: $(TASK)
	@$(TASK) -l
