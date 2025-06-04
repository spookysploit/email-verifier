# 📧 Email Verifier

Простая программа для проверки валидности Email написанная на Go. Позволяет проверить формат email, а также валидность доменного имени.

---

## Установка и запуск
Скачайте исполняемый файл под вашу ОС перейдя по ссылке:
```
https://github.com/spookysploit/email-verifier/releases/tag/email-verifier
```
### Запуск:
Linux:\
```
./email-verifier_linux_x64
```
Windows:\
```
.\email-verifier_x64.exe
```
### Использование:
Один email
```
Enter email(s) (comma-separated):
example@gmail.com
Email: example@gmail.com
Domain: gmail.com
- MX: true
- SPF: true
- DMARC: true
```
Несколько через запятую
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

