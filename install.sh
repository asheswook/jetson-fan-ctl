#!/bin/bash

# root check
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

echo "Installing jetson-fan-ctl to /usr/bin/..."
rm /usr/bin/jetson-fan-ctl 2>/dev/null
cp build/jetson-fan-ctl /usr/bin/
echo "Successfully installed jetson-fan-ctl to /usr/bin/"

echo "Registering service to /lib/systemd/system/..."
cp build/jetson-fan-ctl.service /lib/systemd/system/
chmod 644 /lib/systemd/system/jetson-fan-ctl.service
echo "Successfully registered service to /lib/systemd/system/"

echo "Creating config at /etc/jetson-fan-ctl.conf ..."
rm /etc/jetson-fan-ctl.conf 2>/dev/null
cp build/jetson-fan-ctl.conf /etc/
chmod 664 /etc/jetson-fan-ctl.conf
echo "Successfully created config at /etc/jetson-fan-ctl.conf"

echo "Starting service..."
systemctl daemon-reload
systemctl start jetson-fan-ctl
echo "Successfully started service"

echo "Enabling service for startup..."
systemctl enable jetson-fan-ctl
echo "Successfully enabled service for startup"

echo -e "\033[34m"Congratulations! jetson-fan-ctl is now installed and running.
echo ""
echo "To change the fan speed," 
echo -e edit /etc/jetson-fan-ctl.conf and restart the service with:"\033[0m"
echo -e "\033[32m"sudo systemctl restart jetson-fan-ctl"\033[0m"