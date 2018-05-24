"""The python reader reads every line in the Dataset row after row and sends the
data in a HTTP post request to the Cassandra Database every second making
it  highly efficient and adaptable for real-time use."""


import csv
import time
import json
import sys
import requests
from collections import OrderedDict


exampleFile = open('dataset.csv')
exampleReader = csv.DictReader(exampleFile)

for row in exampleReader:
        sorted_row=OrderedDict(sorted(row.items(),key=lambda item: exampleReader.fieldnames.index(item[0])))
        print('Row #' + str(exampleReader.line_num) + ' ' + str(sorted_row))
        time.sleep(1)
        json.dump(sorted_row,sys.stdout)
        sys.stdout.write('\n')
        time.sleep(1)
	resp=requests.post('http://localhost:8006/patients/new', data=sorted_row)
	resp.text
	time.sleep(1)
