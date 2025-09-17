# ⚙️ Cài đặt và chạy Go

Tài liệu này hướng dẫn cách cài đặt ngôn ngữ lập trình Go trên **Windows**, **macOS**, và **Linux**.  
Bạn có thể xem thêm tài liệu chính thức tại: [https://go.dev/doc/install](https://go.dev/doc/install)

---

## 📦 1. Cài đặt Go

### Windows
1. Truy cập trang tải Go: [https://go.dev/dl/](https://go.dev/dl/)
2. Tải file cài đặt `.msi` cho Windows (64-bit).
3. Chạy file `.msi` và làm theo hướng dẫn.
4. Sau khi cài đặt, mở **Command Prompt** hoặc **PowerShell** và kiểmtra:

```powershell
go version
```

### MacOS
Cách 1: Dùng Homebrew (khuyên dùng)

```bash
brew install go
```

Cách 2: Cài đặt thủ công
1. Tải gói .pkg tại https://go.dev/dl/.
2. Mở file và làm theo hướng dẫn cài đặt.
3. Kiểm tra phiên bản:

```bash
go version
```

### Linux
Ubuntu / Debian

```bash
sudo apt update
sudo apt install -y golang
```

Fedora / CentOS / RHEL

```bash
sudo dnf install golang -y
```

Cài đặt thủ công (tất cả bản Linux)

1. Tải file .tar.gz tại https://go.dev/dl/.
2. Giải nén vào /usr/local:

```bash
sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz
```

3. Thêm Go vào PATH (thêm dòng này vào ~/.bashrc hoặc ~/.zshrc):

```bash
export PATH=$PATH:/usr/local/go/bin
```

4. Reload shell:

```bash
source ~/.bashrc
```

5. Kiểm tra:

```bash
go version
```

---

## 🏃 2. Chạy chương trình Go đầu tiên

### 1. Tạo một file mới hello.go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```


### 2. Chạy chương trình:

```bash
go run hello.go
```

### 3. Biên dịch ra file thực thi:

```bash
go build hello.go
./hello   # Trên macOS/Linux
hello.exe # Trên Windows
```

---

## 🛠 3. Kiểm tra môi trường Go

- Kiểm tra GOPATH:

```bash
go env GOPATH
```

- Kiểm tra cấu hình đầy đủ:

```bash
go env
```