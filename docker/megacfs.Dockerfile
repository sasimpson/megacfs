FROM golang

ENV GO111MODULE=on
ENV BUILDPATH="build/"

#copy over source tree
WORKDIR /app
COPY . .

RUN make install

