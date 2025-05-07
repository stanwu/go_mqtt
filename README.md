# Go MQTT 專案

## 簡介

此專案是一個使用 Go 語言開發的 MQTT 客戶端範例，展示如何連接到 MQTT broker，訂閱主題，並定期發佈 CPU 使用率的訊息。

## 功能

- 連接到 MQTT broker
- 訂閱指定的主題
- 每 5 秒發佈一次 CPU 使用率的訊息

## 使用方式

1. 確保已安裝 Go 環境。
2. 安裝必要的依賴：
   ```bash
   go mod tidy
   ```
3. 執行程式：
   ```bash
   go run main.go
   ```
4. 編譯程式：
   ```bash
   go build
   ```

## 注意事項

- 預設的 MQTT broker 是 `tcp://broker.emqx.io:1883`，可以根據需求修改程式碼中的 `broker` 常數。
- 預設的主題是 `test/topic123`，可以根據需求修改程式碼中的 `topic` 常數。
- 預編譯版本為 macOS Intel 版本

## 授權

此專案採用 MIT 授權條款。