#!/usr/bin/python3
import requests, json

HOST = "http://127.0.0.1:8096/itop/webservices/rest.php?version=1.0"

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
    "json_data": json_str
}


# secure_rest_services
def get():
    r = requests.post(HOST, data=json_data)
    return r


if __name__ == "__main__":
    result = get()
    print(json.dumps(result.json()))