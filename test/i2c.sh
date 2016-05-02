#!/bin/bash

ADDRESS="0x20"

IODIRA="0x00"
OLATA="0x14"
GPPUA="0x0C"

# Set direction: 7 is input rest are output
i2cset -y 1 $ADDRESS $IODIRA 0x80
i2cset -y 1 $ADDRESS $GPPUA 0x80

for i in {1..7}
do
    sleep 0.1
    i2cset -y 1 $ADDRESS $OLATA 0x0$i
done

i2cset -y 1 $ADDRESS $OLATA 0x00
