"""The rules are based on the patient’s state whether the patient is active
or not and also depends on the Patient’s Heart Rate and Blood Pressure threshold value.
It throws an Alert if the values are not ideally normal."""

import json
from pprint import pprint
import time
with open('data.json') as json_data:

d = json.load(json_data)
   json_data.close()

   for patient in d["patients"]:
       if (patient['heartrate'] > 72 and patient['state'] == 0) :
               print ("HeartRate :")
               print patient['heartrate']
               print ("State is : Resting state ")
               print "Alert!"
               time.sleep(1)
       else:
               print ("HeartRate :")
               print patient['heartrate']
               print ("State is : Active State ")
               print "No issues"
