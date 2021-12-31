# syntax=docker/dockerfile:1

# pull latest golang version
FROM golang

# setup workdir in the image
WORKDIR /build

# # Fetch dependency/mod files
# COPY go.mod ./
# COPY go.sum ./

# copy app source code into container workdir
COPY . ./


# download and copy dependencies
RUN go mod download

# copy app source code into container workdir
COPY . ./

# create a binary executable for the application under the image workdir/app
RUN go build -o ./app

# expose port 5000
EXPOSE 5000

# run the executable to start the image
CMD [ "./app" ]
