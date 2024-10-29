FROM archlinux:latest

RUN pacman -Syu --noconfirm && pacman -S --noconfirm go 

WORKDIR /app

COPY . .

RUN go build ip.go

CMD ["./ip"]
   

