FROM golang:1.15

RUN mkdir /hai
ADD . /hai
WORKDIR /hai/app
RUN go build -o main 
CMD ["/hai/app/main"]

EXPOSE 8080
