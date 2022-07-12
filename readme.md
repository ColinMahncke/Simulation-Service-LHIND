# Simulation Service Sensors LHIND

The Simulation Service generates random data for different sensors and uploads them to Kafka 

## Functions  
- With the Simulation Service you can configurate your different "fake"-sensors. They can give out data for AreaCounting Sensors and LineCounting Sensors.
- There is also a configuration for the ramdomness, whether it should calculate different data or the same data everytime.
- The is also a variable with which one allows you to change the time between incoming data
- Furthermore the user can configurate the randomness between the score of data. (max- and min_rand)

---
- ->The data will be uploaded on a Kafka host, which is configurable


## GenerateEvents

The Generic GenerateEvents makes up two different functions. The GenerateAreaCount and the GenerateLineCount. Both of them generate random data. 
- GenerateAreaCount:    Counts at a moment every                   object in range of the                     sensor

- GenerateLineCount:    Counts the whole time the                     object which cross the line                    in both directions

The value gets connectet to other Informations (SensorName, measurement_timestamp, unit) and then the whole information is uploaded to kafka.



## Environment variables
- interval
- area_sensors
- line_sensors
- time
- host
- area_min_rand
- area_max_rand
- line_min_rand
- line_max_rand



