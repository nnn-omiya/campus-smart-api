# campus-smart-api

Dockerを含めたapiサーバー

## path構成

- POST /device/create
  - {type: "controller", name: "dev_test"}
- POST /device/login
  - {name: "dev_test", address: "10.8.100.121"}
- POST /device/status
  - {name: "dev_test", status: "200"}
- POST /device/error
  - {name: "dev_test", type: "1", message: "Internet error"}
- POST /api/device_control
  - { power: true, detail: { mode: 0, temperature: 0, direction: 0, volume: 0 }}
- GET /api/device_status
  - ?type=controller&(pointing=1)
