wasm:
	tinygo build -o main.wasm -target wasi -scheduler none -no-debug .