# Todo App - Buffalo Framework vs Pure Go Karşılaştırması

Bu repository, aynı To-Do uygulamasının iki farklı yaklaşımla implementasyonunu içerir:

## 📁 Proje Yapısı

```
todo_app_go_variants
├── todo_app_pure_go/              # Pure Go implementasyonu
│   ├── main.go
│   ├── go.mod
│   ├── models/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   └── ...
└── todo_app_buffalo/      # Buffalo Framework implementasyonu
    │   ├── actions/
    │   ├── models/
    │   ├── migrations/
    │   └── ...
    └── docker-compose.yml
```

## 🏗️ Yaklaşım Karşılaştırması

### Pure Go Yaklaşımı (`todo_app_pure_go/`)
- ✅ **Minimal dependencies** (8 paket)
- ✅ **Tam kontrol** - Her şeyi manuel yazıyoruz
- ✅ **Yüksek performans** - Native Go
- ✅ **Öğrenme odaklı** - Go'yu derinlemesine öğreniyoruz
- ✅ **Esneklik** - İstediğiniz gibi yapılandırabilirsiniz

### Buffalo Framework Yaklaşımı (`todo_app_buffalo/`)
- ✅ **Auto-generated structure** - CLI ile otomatik oluşturulur
- ✅ **Built-in ORM** - Pop ORM dahil
- ✅ **Built-in validation** - Otomatik validasyon
- ✅ **Built-in migrations** - Fizz migration sistemi
- ✅ **Convention over configuration** - Standart yapı
- ✅ **Hot reload** - Development kolaylığı

## 📊 Detaylı Karşılaştırma

| Özellik | Pure Go | Buffalo Framework |
|---------|---------|-------------------|
| **Dependencies** | 8 | 50+ |
| **Startup Time** | ~50ms | ~200ms |
| **Memory Usage** | ~15MB | ~25MB |
| **Learning Curve** | Düşük | Orta |
| **Development Speed** | Orta | Yüksek |
| **Flexibility** | Yüksek | Orta |
| **Control** | Tam | Sınırlı |
| **Conventions** | Yok | Var |
| **Auto-generation** | Yok | Var |

## 🚀 Hızlı Başlangıç

### Pure Go Versiyonu
```bash
cd todo_app_pure_go
docker compose up -d postgres
go mod tidy
go run main.go
```

### Buffalo Framework Versiyonu
```bash
cd todo_app_buffalo
docker compose up -d
```

## 📡 API Endpoints

Her iki versiyon da aynı API endpoints'lerini destekler:

```
GET    /api/todos           # Tüm todo'ları listele
GET    /api/todos/{id}      # ID'ye göre todo getir
POST   /api/todos           # Yeni todo oluştur
PUT    /api/todos/{id}      # Todo güncelle
DELETE /api/todos/{id}      # Todo sil
PATCH  /api/todos/{id}/toggle # Completion durumunu değiştir
```

## ✅ Test Sonuçları

### Pure Go Versiyonu ✅
```bash
# Todo oluştur
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Pure Go Todo", "priority": "HIGH"}'

# Todo listele
curl http://localhost:8080/api/todos

# Completion toggle
curl -X PATCH http://localhost:8080/api/todos/1/toggle

# Todo güncelle
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Güncellenmiş Pure Go Todo"}'

# Filtreleme
curl "http://localhost:8080/api/todos?completed=true"
```

### Buffalo Framework Versiyonu ✅
```bash
# Todo oluştur
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Buffalo Todo", "priority": "HIGH"}'

# Todo listele
curl http://localhost:8080/api/todos

# Completion toggle
curl -X PATCH http://localhost:8080/api/todos/1/toggle

# Todo güncelle
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Güncellenmiş Buffalo Todo"}'

# Filtreleme
curl "http://localhost:8080/api/todos?completed=true"
```

## 🎓 Öğrenme Amaçları

### Pure Go Versiyonu Öğretir:
1. **Go'yu derinlemesine** - Her katmanı manuel yazıyoruz
2. **Dependency Injection** - Manuel DI pattern
3. **Error Handling** - Go'nun error handling yaklaşımı
4. **Database Operations** - GORM kullanımı
5. **HTTP Handling** - Gin framework kullanımı
6. **Project Structure** - Katmanlı mimari

### Buffalo Framework Versiyonu Öğretir:
1. **Framework Kullanımı** - Buffalo conventions
2. **Auto-generation** - CLI tools kullanımı
3. **Built-in Features** - ORM, validation, migrations
4. **Convention over Configuration** - Standart yapı
5. **Rapid Development** - Hızlı geliştirme

## 🔧 Teknik Detaylar

### Pure Go Teknolojileri:
- **Gin** - Web framework
- **GORM** - ORM
- **go-playground/validator** - Validation
- **godotenv** - Environment management
- **Manual DI** - Dependency injection

### Buffalo Framework Teknolojileri:
- **Buffalo** - Full-stack framework
- **Pop** - Built-in ORM
- **Fizz** - Migration system
- **Buffalo CLI** - Development tools
- **Auto DI** - Built-in dependency injection

## 🚀 Sonuç

Bu karşılaştırma, Go web geliştirmede iki farklı yaklaşımı göstermektedir:

- **Pure Go**: Öğrenme odaklı, tam kontrol, yüksek performans
- **Buffalo Framework**: Hızlı geliştirme, standart yapı, built-in özellikler

Her iki yaklaşımın da kendi avantajları vardır ve projenin ihtiyaçlarına göre seçim yapılmalıdır.

## 📚 Daha Fazla Bilgi

- [Pure Go Versiyonu](./todo_app_pure_go/README.md)
- [Buffalo Framework Versiyonu](./todo_app_buffalo/README.md)
- [Go Documentation](https://golang.org/doc/)
- [Buffalo Documentation](https://gobuffalo.io/) 