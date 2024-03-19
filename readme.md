# BOT NOI

# วิธีติดตั้ง

เลือก file path ที่ต้องการจะเก็บไฟล์ไว้

```
cd filepath
```

clone repo

```
git clone https://github.com/Timemi11/BackEnd_botnoi_linebot.git
```

เปลี่ยน directory ไปที่ BackEnd_botnoi_linebot

```
cd BackEnd_botnoi_linebot
```

set linebot
```
1.นำ Channel Secret ไปใส่ตรงสตริงตัวแรก บรรทัดที่ 15
2.นำ Channel access token ไปใส่ตรงสตริงตัวสอง บรรทัดที่ 16
```

ติดตั้ง dependency เผื่อไม่มี go.mod

```
go mod init
```

start server

```
go run main
```

เปิด cmd อีกตัวแล้วพิมพ์

```
ngrok http localhost:5000
```

นำ url https ตรง Forwarding ไปใส่ใน Webhook URL เช่น

```
https://a50a-2001-fb1-17e-579e-9dd8-ac00-9457-a3ab.ngrok-free.app/chat
หรือตาม format นี้
[https//....]/chat
```
