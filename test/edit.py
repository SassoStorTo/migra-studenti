#! python
import requests

print("start edit")
s = requests.put("http://localhost:8080/student", data={'id': '1',
                                                         'name': 'pipppo', 
                                                         'lastname':'cocaina', 
                                                         'year': 2004, 'month':7, 
                                                         'day':30})
print(f"[student]: {s.status_code} \n {s.headers} \n {s.content}\n")

s = requests.put("http://localhost:8080/class", data={'id': 1,
                                                       'year': 2, 
                                                       'section': "TAR", 
                                                       'schoolyear': 2024, 
                                                       'idmajor': 2})
print(f"[class]: {s.status_code} \n {s.headers} \n {s.content}\n")

s = requests.put("http://localhost:8080/major", data={'id':1, 'name': 'topo'})
print(f"[major]: {s.status_code} \n {s.headers} \n {s.content}\n")

print("end edit")