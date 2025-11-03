import subprocess

# Specify the username
username = "Baha"  # Change to the user you want to check

try:
    # 'tasklist /V' lists all running processes with their username on Windows
    result = subprocess.run(["tasklist", "/V"], capture_output=True, text=True, check=True)
    processes = [line for line in result.stdout.splitlines() if username in line]

    count = len(processes)
    print(f"[OK] User '{username}' has {count} running processes.")
except subprocess.CalledProcessError:
    print("[ERROR] Failed to count processes.")
