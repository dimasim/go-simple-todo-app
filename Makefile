# Nama untuk file binary produksi kita
BINARY_NAME=go-simple-todo-app

## run: Menjalankan aplikasi di mode development dengan air.
run:
	air

## build: Mengompilasi aplikasi untuk produksi.
build:
	go build -o $(BINARY_NAME) ./main.go

## clean: Membersihkan file hasil build.
clean:
	go clean
	rm -f $(BINARY_NAME)