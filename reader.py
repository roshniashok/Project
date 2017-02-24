import csv
import json


f = open( 'dataset.csv', 'r' )

reader = csv.DictReader( f)
out = json.dumps( [ row for row in reader ],indent=4)
messages=stream_array(out)
for message in messages:
	handle_message(message)

