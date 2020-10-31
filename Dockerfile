FROM golang:1.14.4-alpine AS builder

# 1. Precompile the entire go standard library into the first Docker cache layer: useful for other projects too!
RUN CGO_ENABLED=0 GOOS=linux go install -v -installsuffix cgo -a std

# 2. Prepare and enter the src folder
WORKDIR /go/src/github.com/OpenPEC/

# 3. Download and precompile all third party libraries, ignoring errors (some have broken tests or whatever)
ADD go.mod .
ADD go.sum .
RUN go mod download -x
RUN go list -m all | tail -n +2 | cut -f 1 -d " " | awk 'NF{print $0 "/..."}' | CGO_ENABLED=0 GOOS=linux xargs -n1 go build -v -installsuffix cgo -i; echo done

# 4. Add the sources
ADD . .

# 5. Compile! Should only compile our sources since everything else is precompiled
RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo -o /go/src/github.com/OpenPEC/ -ldflags "-s -w" /go/src/github.com/OpenPEC/

# 6. Put everything in a SMOL container that weighs a few MBs
FROM scratch
COPY --from=builder /go/src/github.com/OpenPEC/ /go/src/github.com/OpenPEC/
CMD ["/go/src/github.com/OpenPEC/OpenPEC"]

EXPOSE 9090

#sudo docker build -t openpec .
#sudo docker run --network="host" --publish 9090:9090 --name openpec --rm openpec