#!/usr/bin/python3
import requests, json

HOST = "http://127.0.0.1:8000/webservices/rest.php?version=1.3"

json_str2 = json.dumps({
    "operation": "core/update",
    "class": "UserRequest",
    "comment": "close ticket",
    "key": {
        "ref": "R-000006",
        "status": "resolved"
    },
    "output_fields": "status",
    "fields": {
        "status": "new"
    }
})

json_str = json.dumps({
    "operation":
    "core/get",
    "class":
    "UserRequest",
    "key":
    "SELECT UserRequest WHERE operational_status = 'ongoing'",
    "output_fields":
    # "request_type,servicesubcategory_name,urgency,origin,caller_id_friendlyname,impact,title,description",
    "*",
})
json_data = {
    "auth_user": "admin",
    "auth_pwd": "goodluck@123.",
    "json_data": json_str2
}


# secure_rest_services
def get():
    r = requests.post(HOST, data=json_data)
    return r


if __name__ == "__main__":
    result = get()
    print(json.dumps(result.json()))
    # print(json.loads(json_str2))