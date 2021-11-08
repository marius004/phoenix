# https://hub.docker.com/r/ubuntu/postgres
FROM ubuntu/postgres:latest

# SET POSTGRES ENV VARIABLES TODO!

# Install dependencies
RUN apt-get update
RUN apt-get install -y wget git gcc

# Setup isolate 
RUN apt-get update -y && \
    apt-get install -y libcap-dev asciidoc-base && \
    cd ~/Downloads && \
    git clone https://github.com/marius004/isolate && \
    cd isolate && \
    make && \
    make install

# Setup go
RUN wget -P /tmp https://dl.google.com/go/go1.16.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /tmp/go1.16.6.linux-amd64.tar.gz
RUN rm /tmp/go1.16.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/phoenix

COPY . . 

CMD ["go", "run", "."]