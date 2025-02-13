# go baseimage
FROM golang:1.23

# create app directory
RUN  mkdir /app

# set app directory as a working directory
WORKDIR /app

COPY . /app 

RUN go mod download

CMD ["go", "run", "./cmd/api"]

