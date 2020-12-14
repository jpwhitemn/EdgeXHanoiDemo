#!/bin/sh
## Go to Device Modbus DS and start it
echo "Starting device modbus - with temperature probe"
cd /home/jim/forks/device-modbus-go/cmd
./device-modbus &

## Go to Device SNMP DS and start it
echo "Starting device SNMP - with Patlite"
cd /home/jim/forks/device-snmp-go/cmd
./device-snmp-go &

echo "Device Services all started"

wait
echo "EdgeX demo done"

