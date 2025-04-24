# 🍁 MapleStory v1.2.6 Private Server (Go)

**A resurrection of the classic MapleStory experience — written from scratch in pure Go.**  
This project revives the legendary 2007-era MapleStory v1.2.6 client using a fully custom, opcode-driven TCP server.

### ‼️이 리포지토리는 개인 학습용으로만 유지되고 있습니다.
본 리포지토리의 코드를 기반으로 실제 사설 서버를 운영하는 행위는 금지되며,
이에 따른 모든 법적 책임은 사용자 본인에게 있습니다.

### ‼️This repository is maintained strictly for personal learning purposes only.
Operating a private server using the code from this repository is strictly prohibited.
All legal responsibilities for any such use lie solely with the user.

---

## 🎯 Version

- **Client Version:** KMS 1.2.6
- **Tested With:** Original 1.2.6 KMS client (manually patched to connect to `127.0.0.1:8484`)

## 🎯 How to start?
First,

> docker compose -f ./.docker/db/docker-compose.yml up -d

Second,

> docker compose -f ./.docker/app/docker-compose.app.yml up --build