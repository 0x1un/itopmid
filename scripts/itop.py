#!/usr/bin/python3
import requests, json

HOST = "http://140.246.60.181:8096/itop/webservices/rest.php?version=1.3"

json_str2 = json.dumps({
    "operation": "core/update",
    "class": "UserRequest",
    "comment": "close ticket",
    "key": {
        "status": "closed"
    },
    "output_fields": "status",
    "fields": {
        "status": "resolved"
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
    "auth_user": "",
    "auth_pwd": ".",
    "json_data": json_str2
}


# secure_rest_services
def get():
    r = requests.post(HOST, data=json_data)
    return r


if __name__ == "__main__":
    result = get()
    print(json.dumps(result.json()))