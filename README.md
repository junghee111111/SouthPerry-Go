# 🍁 MapleStory v1.2.6 Private Server (Go)

**A resurrection of the classic MapleStory experience — written from scratch in pure Go.**  
This project revives the legendary 2007-era MapleStory v1.2.6 client using a fully custom, opcode-driven TCP server.

> “No Netty. No Spring. Just packets, opcodes, and pain.”

---

## 🎯 Version

- **Client Version:** KMS 1.2.6
- **Tested With:** Original 1.2.6 KMS client (manually patched to connect to `127.0.0.1:8484`)

## 🎯 How to start?
First,

> docker compose -f ./.docker/db/docker-compose.yml up -d

Second,

> docker compose -f ./.docker/app/docker-compose.app.yml up --build