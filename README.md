# PMS - Personal Management System

Build Dev and Run
go build -ldflags="-X main.isBuild=dev" && ./backend

Build Prodaction and Run
go build -ldflags="-X main.isBuild=prod" && ./backend