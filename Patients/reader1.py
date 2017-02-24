import csv
import time
import json
import sys
import requests
from collections import OrderedDict
import urllib2

exampleFile = open('dataset.csv')
exampleReader = csv.DictReader(exampleFile)
for row in exampleReader:
	sorted_row=OrderedDict(sorted(row.items(),key=lambda item: exampleReader.fieldnames.index(item[0])))
	print('Row #' + str(exampleReader.line_num) + ' ' + str(sorted_row))
	time.sleep(1)
	json.dump(sorted_row,sys.stdout)
	sys.stdout.write('\n')
	time.sleep(1)
	#w = requests.post(url='http://localhost:8005/patients/new', data=json.dumps(sorted_row))
	req = urllib2.Request('http://localhost:8005/patients/new')
	req.add_header('Content-Type', 'application/json')
	response = urllib2.urlopen(req, json.dumps(sorted_row))
	time.sleep(1)
