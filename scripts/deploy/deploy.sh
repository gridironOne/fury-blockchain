#!/usr/bin/env bash

# Copy unit file to /etc/systemd/system/
sudo cp fury.service /etc/systemd/system/

# Reload all unit files
sudo /bin/systemctl daemon-reload

# Enable and start the service
sudo /bin/systemctl enable fury.service
sudo /bin/systemctl restart fury.service
