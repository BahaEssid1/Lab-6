import subprocess

# Path to your script to schedule
script_path = r"C:\Users\Baha\Desktop\Developpement full stack\Lab 6\clean_logs.py"

# Task name
task_name = "Lab6DailyTask"

# Command to create the scheduled task at 2 AM daily
command = f'schtasks /Create /SC DAILY /TN "{task_name}" /TR "python \\"{script_path}\\"" /ST 02:00 /F'

try:
    result = subprocess.run(command, shell=True, check=True, capture_output=True, text=True)
    print("[OK] Scheduled task created successfully!")
    print(result.stdout)
except subprocess.CalledProcessError as e:
    print("[ERROR] Failed to create scheduled task:", e.stderr)
