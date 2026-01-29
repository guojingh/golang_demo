import json
import requests


def send_post_request_status_1(data):
    """Send a POST request to the specified API endpoint."""
    API_URL = 'http://10.138.170.248:8801/mcp/inner/list'
    HEADERS = {'Content-Type': 'application/json'}
    response = requests.get(API_URL, headers=HEADERS, params=data)
    return response.text

print(send_post_request_status_1({"status":"1"}))