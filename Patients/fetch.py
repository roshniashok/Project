"""This python fetcher fetches the posted data from the Cassandra Database
using a HTTP Get request from the localhost. It dumps the data into a json file. """

import requests
import time
import json
from pprint import pprint


r = requests.get("http://localhost:8017/patients")

dataa=json.loads(r.text)

with open('data.json','w')as f:
       data=json.dump(dataa,f,indent=4,sort_keys=True)
       pprint(data)
