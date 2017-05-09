# -*- coding:utf-8 -*-
import requests
import json
# import json

# url = 'https://api.github.com/some/endpoint'
# payload = {'some': 'data'}
# r = requests.post(url, data=json.dumps(payload))

# or

# url = 'https://api.github.com/some/endpoint'
# payload = {'some': 'data'}
# r = requests.post(url, json=payload)

if __name__ == "__main__":
    params = {"_msg_type" : 11020, "district" : "11", "exempt_flag" : 1, "driver_id" : 12345, "uid":54322}
    data = {"params" : json.dumps(params)}
    headers = {"didi-header-rid" : "111"}
    r = requests.post("http://localhost:3219/gulfstream/credit/collectEvent", data=data, headers = headers)
    print r.text

