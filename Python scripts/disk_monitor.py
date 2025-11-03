import shutil

# Path to check
path = "/"  # On Windows, you can use "C:\\"

total, used, free = shutil.disk_usage(path)
percent_free = (free / total) * 100

print(f"Disk free: {percent_free:.2f}%")

if percent_free < 10:
    print("[ALERT] Disk space is below 10%!")
else:
    print("[OK] Disk space is sufficient.")
