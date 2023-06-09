import requests
from requests.auth import HTTPDigestAuth
import json
import os

def my_function(filename,jData):
  f = open(filename, "a")
  f.write(str(jData))
  f.close()

  #open and read the file after the appending:
  f = open(filename, "r")
  print(f.read())

# Replace with the correct URL
url = "http://api.weatherstack.com/current?access_key=36c6ac47321e29ae46c7da6315a67e58&query=New%20York"

# It is a good practice not to hardcode the credentials. So ask the user to enter credentials at runtime
#myResponse = requests.get(url,auth=HTTPDigestAuth(raw_input("username: "), raw_input("Password: ")), verify=True)
myResponse = requests.get(url)
print (myResponse.status_code)

# For successful API call, response code will be 200 (OK)
if(myResponse.ok):

    # Loading the response data into a dict variable
    # json.loads takes in only binary or string variables so using content to fetch binary content
    # Loads (Load String) takes a Json file and converts into python data structure (dict or list, depending on JSON)
    jData = json.loads(myResponse.content)

    print("The response contains {0} properties".format(len(jData)))
    print("\n")
    #print(myResponse.content)
    for key in jData:
      print (key)
      print (jData[key])
else:
  # If response code is not ok (200), print the resulting http error code with description
    myResponse.raise_for_status()
my_function("output.json",jData)
