import zipfile
import os

# Path to the folder you want to back up
# On Windows, for example: C:\Users\Baha\Desktop\Developpement full stack\Lab 6\logs
folder_to_backup = r"C:\Users\Baha\Desktop\Developpement full stack\Lab 6\logs"

# Output ZIP file
# On Windows, you can save it in the same folder or anywhere else
output_zip = r"C:\Users\Baha\Desktop\Developpement full stack\Lab 6\backup.zip"

# Create the ZIP archive
with zipfile.ZipFile(output_zip, 'w', zipfile.ZIP_DEFLATED) as zipf:
    for root, dirs, files in os.walk(folder_to_backup):
        for file in files:
            file_path = os.path.join(root, file)
            # Keep folder structure relative to the folder_to_backup
            arcname = os.path.relpath(file_path, folder_to_backup)
            zipf.write(file_path, arcname)

print(f"Backup created at {output_zip}")
