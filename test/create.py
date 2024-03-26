#! python
import requests

print("start create")
s = requests.post("http://localhost:8080/student", data={'name': 'paolo', 
                                                         'lastname':'magnani', 
                                                         'year': 2004, 'month':7, 
                                                         'day':30})
print(f"[student]: {s.status_code} \n {s.headers} \n {s.content}\n")

s = requests.post("http://localhost:8080/class", data={'year': 2, 
                                                       'section': "I", 
                                                       'schoolyear': 2024, 
                                                       'idmajor': 2})
print(f"[class]: {s.status_code} \n {s.headers} \n {s.content}\n")

s = requests.post("http://localhost:8080/major", data={'name': 'pippo'})
print(f"[major]: {s.status_code} \n {s.headers} \n {s.content}\n")

s = requests.post("http://localhost:8080/studentclass", data={'ids': 1, 'idc': 1})
print(f"[student-class]: {s.status_code} \n {s.headers} \n {s.content}")

print("end create")