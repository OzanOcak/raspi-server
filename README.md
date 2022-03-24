brew install nmap

sudo nmap -sS -p 22 192.168.1.1-255 | grep -B4 -A2 "22/tcp open”

ssh pi@192.168.1.18

go version

mkdir gohttpserver
cd gohttpserver
go mod init github.com/OzanOcak/gohttserverp

go build -o server

./server

curl http://localhost:3000/world


