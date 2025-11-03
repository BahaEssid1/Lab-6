import os

# Folder to scan
folder = r"C:\Users\Baha\Desktop\Developpement full stack\Lab 6"  # Change if needed

# Get all files with their sizes
files = []
for filename in os.listdir(folder):
    file_path = os.path.join(folder, filename)
    if os.path.isfile(file_path):
        size = os.path.getsize(file_path)  # size in bytes
        files.append((filename, size))

# Sort files by size descending
files.sort(key=lambda x: x[1], reverse=True)

# Take top 5
top5 = files[:5]

print("[OK] 5 plus gros fichiers dans le dossier:")
for name, size in top5:
    print(f"{name} - {size / 1024:.2f} KB")
