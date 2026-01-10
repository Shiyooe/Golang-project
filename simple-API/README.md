## Cara Menjalankan

1. Inisialisasi module:
```bash
go mod init simple-api
```

2. Jalankan aplikasi:
```bash
go run cmd/api/main.go
```

3. Test API:
```bash
# Get all users
curl http://localhost:8080/users

# Create user
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Bob Wilson","email":"bob@example.com"}'
```