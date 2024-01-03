#!/bin/bash

go build -o bookings cmd/web/*.go

./bookings -dbname=bookings -dbuser=dansogbey -cache=false -production=false  