FROM golang:1.16.6-buster as build

RUN mkdir /app
WORKDIR /app

COPY . .

RUN cd pattern-extractor && CGO_ENABLED=0 go build -o /pattern-extractor -ldflags '-extldflags "-static"' -tags timetzdata ./main

FROM scratch

COPY --from=build /pattern-extractor /bin/pattern-extractor

CMD ["/bin/pattern-extractor"]
