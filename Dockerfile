FROM golang:1.19.4-alpine

### Indica cual es la carpeta donde trabajará en el contenedor de docker
WORKDIR /home/courses/go/goweb-crud

### COPY [.]->maquina host   [/home/courses/go/goweb-crud]-> contendor
COPY . .

RUN go mod download && go mod verify

RUN go build -o goweb-crud

EXPOSE 3000

### CMD ["/home/courses/go/goweb-crud/goweb-crud"]
CMD ["./goweb-crud"]