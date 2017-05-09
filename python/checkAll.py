# -*- coding:utf-8 -*-
import requests
import json

if __name__ == "__main__":
    data = {"sid" : "flash", "checkDate" : "2017-05-09", "cityId" : "0"}
    headers = {"x-header-id" : "111"}
    # post a form
    r = requests.post("www.baidu.com", data=data, headers = headers)
    print r.text

