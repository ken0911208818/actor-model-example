# concurrency 

### 使用兩個waitGroup 以及一個channal進行併發
### 一個waitGroup控制 要開啟go 要開啟go goroutine 的數量
### 一個waitGroup控制 task的數量

### channal 負責傳遞資料 當task做完時 自動銷毀 channal
 