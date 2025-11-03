import requests

# Website to check
url = "https://azure.microsoft.com/en-us/"

try:
    response = requests.get(url, timeout=5)  # timeout in seconds
    if response.status_code == 200:
        print(f"[OK] {url} is reachable!")
    else:
        print(f"[ERROR] {url} returned status code {response.status_code}")
except requests.exceptions.RequestException as e:
    print(f"[ERROR] Cannot reach {url}: {e}")
