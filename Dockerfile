FROM golang:1.9

ARG PROJECT="github.com/netm4ul/netm4ul"
ARG FULL_PATH=/go/src/${PROJECT}
ARG EXECUTABLE=${FULL_PATH}/netm4ul
RUN useradd netm4ul

COPY . ${FULL_PATH}
WORKDIR ${FULL_PATH}

# RUN go get -u github.com/golang/dep/...
# RUN make
USER netm4ul
RUN mkdir -p ${FULL_PATH}

RUN go build . 
CMD [${EXECUTABLE}]
