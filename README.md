# 📧 Email Verifier

A simple Email validation tool written in Go. Allows you to check the email format, as well as the validity of the domain name.

---

## Installation and launch
Download the executable file for your OS by clicking on the link:
```
https://github.com/spookysploit/email-verifier/releases/tag/email-verifier
```
### Launch:
Linux:\
```
./email-verifier_linux_x64
```
Windows:\
```
.\email-verifier_x64.exe
```
### Usage:
One email
```
Enter email(s) (comma-separated):
example@gmail.com
Email: example@gmail.com
Domain: gmail.com
- MX: true
- SPF: true
- DMARC: true
```
Several emails separated by commas
```
Enter email(s) (comma-separated):
example1@gmail.com,hesoyam@internet.com,aezakmi@yahoo.com,non-valid.mail.ru
Email: example1@gmail.com
Domain: gmail.com
- MX: true
- SPF: true
- DMARC: true

Email: hesoyam@internet.com
Domain: internet.com
- MX: true
- SPF: true
- DMARC: true

Email: aezakmi@yahoo.com
Domain: yahoo.com
- MX: true
- SPF: true
- DMARC: true

non-valid.mail.ru IS INVALID (bad format)
```
