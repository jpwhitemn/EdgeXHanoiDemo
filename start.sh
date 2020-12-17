#!/bin/sh
## Go to Device Modbus DS and start it
echo "Starting device modbus - with temperature probe"
cd /home/jim/forks/device-modbus-go/cmd
./device-modbus >/dev/null 2>&1 &

## Go to Device SNMP DS and start it
echo "Starting device SNMP - with Patlite"
cd /home/jim/forks/device-snmp-go/cmd
./device-snmp-go >/dev/null 2>&1 &

## Go to App Service Influx and start it
echo "Starting App Service Influx"
cd /home/jim/forks/app-service-influx/cmd
./app-service-influx >/dev/null 2>&1 &

echo "Device Services and App Service all started"

#wait
#echo "EdgeX demo done"

