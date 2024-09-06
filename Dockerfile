FROM golang:1.20-alpine AS builder 

WORKDIR /app 

RUN  apt-get update -y \
&& apt-get upgrade -y \
&& apt-get install -y gcc defualt-libmysqlclient-dev pkg-config \
&& rm -rf  /var/lib/apt/lists/* 

#to copy .txt file at the root containing  installation content NOT SURE ABOUT THIS ONE 
#Since it takes two args 
COPY init.txt .

#cmd to install the packages present inside of the go module 
RUN  go mod download 
RUN  go build -o main 

#cmd to install the dependencies present in the init.txt file 
RUN go get 

#now that we have downloaded them we can 
CMD ["go run","main.go"]
