#!/usr/bin/env bash

# Disable, stop, and remove unit file
sudo /bin/systemctl disable fury.service
sudo /bin/systemctl stop fury.service
sudo rm /etc/systemd/system/fury.service

# Reload all unit files and reset failed
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl reset-failed
