import subprocess

# Use the exact service name
service_name = "CloudflareWARP"

try:
    # Restart the service
    subprocess.run(["powershell", "Restart-Service", service_name], check=True)
    print(f"[OK] {service_name} restarted successfully.")
except subprocess.CalledProcessError:
    print(f"[ERROR] Failed to restart {service_name}.")
