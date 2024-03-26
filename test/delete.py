#! python

import requests

print("start delete")

s = requests.delete("http://localhost:8080/student", data={'id': 1})
print(f"[student]: {s.status_code} \n {s.headers} \n {s.content}\n")
s = requests.delete("http://localhost:8080/class", data={'id': 1})
print(f"[class]: {s.status_code} \n {s.headers} \n {s.content}\n")
s = requests.delete("http://localhost:8080/major", data={'id': 1})
print(f"[major]: {s.status_code} \n {s.headers} \n {s.content}")
s = requests.delete("http://localhost:8080/studentclass", data={'ids': 1, 'idc': 1})
print(f"[student-class]: {s.status_code} \n {s.headers} \n {s.content}")

print("end delete")