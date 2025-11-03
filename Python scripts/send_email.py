import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

# === Configuration ===
smtp_server = "smtp.gmail.com"  # Replace with your SMTP server
smtp_port = 587                  # Usually 587 for TLS
sender_email = "essidbaha24@gmail.com"
password = "ywkjpcnvakvbpnhr"   # Use app password if using Gmail
receiver_email = "essidbaha18@gmail.com"

subject = "ALERT: Test Email"
message = "This is a test alert message from your script to baha ."

# === Create the email ===
msg = MIMEMultipart()
msg["From"] = sender_email
msg["To"] = receiver_email
msg["Subject"] = subject

msg.attach(MIMEText(message, "plain"))

# === Send the email ===
try:
    with smtplib.SMTP(smtp_server, smtp_port) as server:
        server.starttls()  # Enable TLS
        server.login(sender_email, password)
        server.send_message(msg)
    print("[OK] Email sent successfully!")
except Exception as e:
    print("[ERROR] Failed to send email:", e)
