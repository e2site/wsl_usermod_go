target = override

init: build
build: FORCE ; go run ./cmd/ --config="test.json" --checkExistFile

FORCE: