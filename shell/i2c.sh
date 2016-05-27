#!/bin/bash


echo "21" > /sys/class/gpio/export
echo "in" > /sys/class/gpio/gpio21/direction


i2cset -y 1 0x20 0x00 0x80
i2cset -y 1 0x20 0x0c 0x80
i2cset -y 1 0x20 0x04 0x80
i2cset -y 1 0x20 0x08 0x80
i2cset -y 1 0x20 0x06 0x80

PREVIOUS=
while true
do
  sleep 2
  TIP=$(cat /sys/class/gpio/gpio21/value)
#  if [ "$TIP" != "$PREVIOUS" ]; then
    echo "changed $PREVIOUS > $TIP"
#    i2cget -y 1 0x20 0x07
#    i2cget -y 1 0x20 0x10
    i2cget -y 1 0x20 0x12
 # fi
  PREVIOUS=$TIP
done

