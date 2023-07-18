#!/bin/bash

# root check
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi


echo "Stopping service..."
systemctl stop jetson-fan-ctl
echo "Successfully stopped service"

echo "Disabling startup service..."
systemctl disable jetson-fan-ctl
echo "Successfully disabled service for startup"

echo "Removing config at /etc/jetson-fan-ctl.conf ..."
rm /etc/jetson-fan-ctl.conf 2>/dev/null
echo "Successfully removed config at /etc/jetson-fan-ctl.conf"

echo "Removing service at /lib/systemd/system/ ..."
rm /lib/systemd/system/jetson-fan-ctl.service 2>/dev/null
echo "Successfully removed service at /lib/systemd/system/"

echo "Removing jetson-fan-ctl from /usr/bin/..."
rm /usr/bin/jetson-fan-ctl 2>/dev/null
echo "Successfully removed jetson-fan-ctl from /usr/bin/"

echo "Reloading systemctl daemon..."
systemctl daemon-reload
echo "Successfully reloaded systemctl daemon"

echo -e "\033[34m"jetson-fan-ctl is now uninstalled."\033[0m"