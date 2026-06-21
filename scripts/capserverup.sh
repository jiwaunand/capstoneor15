#!/usr/bin/env bash

podman stop capstonegoserver
podman container rm capstonegoserver
podman volume prune
podman build -t capstonegoserver ./images/capstonegoserver/
podman run -d -p 127.0.0.1:8080:8080 --name capstonegoserver capstonegoserver
