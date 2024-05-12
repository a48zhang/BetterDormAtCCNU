FROM golang:1.22 AS builder
ADD  . src
WORKDIR src
RUN go build main
FROM python:latest
COPY --from=builder /go/src/main /main
COPY --from=builder /go/src/pdfgen/* /
#RUN apt update && apt install python3 -y && apt install python3-pip -y && pip3 install -r /pdfgen/requirements.txt
RUN pip3 install -r /requirements.txt
CMD /main & python3 /app.py