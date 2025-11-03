import os
import time

# Exact path to your logs folder
folder = r"C:\Users\Baha\Desktop\Developpement full stack\Lab 6\logs"

# Time in seconds for 7 days
seven_days = 7 * 24 * 60 * 60
now = time.time()

# Check each file in the folder
for filename in os.listdir(folder):
    file_path = os.path.join(folder, filename)
    if os.path.isfile(file_path) and file_path.endswith(".log"):
        age = now - os.path.getmtime(file_path)
        if age > seven_days:
            print(f"[DELETE] {file_path}")
            os.remove(file_path)
        else:
            print(f"[KEEP] {file_path}")
