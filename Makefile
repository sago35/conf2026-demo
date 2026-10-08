all: build

build:
	mkdir -p out
	tinygo build -o out/conf2025badge-xiao-ble.uf2 --target xiao-ble --size short .

flash:
	tinygo flash --target xiao-ble --size short .

monitor:
	tinygo monitor --target xiao-ble

clean:
	rm -rf out

.PHONY: all build flash monitor clean
