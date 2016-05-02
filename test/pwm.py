import RPi.GPIO as GPIO 
import time
  
GPIO.setmode(GPIO.BCM)  
  
GPIO.setup(16, GPIO.OUT)
GPIO.setup(20, GPIO.OUT)
GPIO.setup(21, GPIO.OUT)
  
r = GPIO.PWM(16, 50)    
g = GPIO.PWM(20, 50)    
b = GPIO.PWM(21, 50)    
  
r.start(0)             
g.start(0)             
b.start(0)             
       
d = 3 
for l in range (0, 3):               
        for dc in range(0, 101, 5):
            r.ChangeDutyCycle(dc)
            g.ChangeDutyCycle(dc/d)
            time.sleep(0.1)
        for dc in range(100, -1, -5):
            r.ChangeDutyCycle(dc)
            g.ChangeDutyCycle(dc/d)
            time.sleep(0.1)           
  
r.stop()                
g.stop()                
b.stop()                
  
GPIO.cleanup()
