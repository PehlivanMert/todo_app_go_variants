# Todo App - Buffalo Framework

Bu proje, **Buffalo Framework** kullanılarak geliştirilmiş bir To-Do uygulamasıdır. Pure Go yaklaşımıyla karşılaştırmak için oluşturulmuştur.

## 🚀 Teknolojiler

- **Go 1.23+**
- **Buffalo Framework** - Full-stack web framework
- **Pop ORM** - Database ORM (Buffalo'nun parçası)
- **PostgreSQL** - Veritabanı
- **Fizz** - Migration sistemi
- **Buffalo CLI** - Development tools

## 📁 Proje Yapısı

```
todo_app_buffalo/
├── actions/           # Controllers (Buffalo'da actions)
│   ├── app.go        # Ana uygulama konfigürasyonu
│   ├── home.go       # Ana sayfa handler
│   └── todos.go      # Todo CRUD işlemleri
├── models/           # Data models
│   ├── models.go     # Database connection
│   └── todo.go       # Todo model
├── migrations/       # Database migrations
├── config/          # Konfigürasyon dosyaları
├── database.yml     # Database konfigürasyonu
├── .env             # Environment variables
└── cmd/app/         # Uygulama entry point
```

## 🏗️ Buffalo vs Pure Go Karşılaştırması

### Buffalo Framework Yaklaşımı:
- ✅ **Auto-generated structure**
- ✅ **Built-in ORM (Pop)**
- ✅ **Built-in validation**
- ✅ **Built-in migrations**
- ✅ **Convention over configuration**
- ✅ **Hot reload**
- ✅ **CLI tools**

### Pure Go Yaklaşımı:
- ✅ **Minimal dependencies**
- ✅ **Full control**
- ✅ **Better learning experience**
- ✅ **Higher performance**
- ✅ **More flexibility**

## 🚀 Kurulum ve Çalıştırma

### 1. Gereksinimler
```bash
# Buffalo CLI kurulumu
go install github.com/gobuffalo/cli/cmd/buffalo@latest

# PostgreSQL kurulumu (Docker ile)
docker run --name postgres-buffalo -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todoapp -p 5432:5432 -d postgres:15
```

### 2. Proje Kurulumu
```bash
# Dependencies yükleme
go mod tidy

# Database migration
buffalo pop migrate

# Uygulamayı çalıştırma
buffalo dev
```

### 3. Docker ile Çalıştırma
```bash
# Build
docker build -t todo-app-buffalo .

# Run
docker run -p 8080:8080 --link postgres-buffalo:postgres todo-app-buffalo
```

## 📡 API Endpoints

### Todo Endpoints
```
GET    /api/todos           # Tüm todo'ları listele
GET    /api/todos/{id}      # ID'ye göre todo getir
POST   /api/todos           # Yeni todo oluştur
PUT    /api/todos/{id}      # Todo güncelle
DELETE /api/todos/{id}      # Todo sil
PATCH  /api/todos/{id}/toggle # Completion durumunu değiştir
```

### Query Parameters
```
?page=1&per_page=10        # Pagination
?completed=true            # Filter by completion
?priority=HIGH             # Filter by priority
```

## 🔧 Buffalo Framework Özellikleri

### 1. Auto-generated CRUD
```go
// Buffalo otomatik olarak CRUD işlemlerini oluşturur
type TodosResource struct {
    buffalo.Resource
}

// Otomatik route mapping:
// GET /todos -> List()
// GET /todos/{id} -> Show()
// POST /todos -> Create()
// PUT /todos/{id} -> Update()
// DELETE /todos/{id} -> Destroy()
```

### 2. Built-in Validation
```go
func (t *Todo) Validate(tx *pop.Connection) (*validate.Errors, error) {
    return validate.Validate(
        &validators.StringIsPresent{Field: t.Title, Name: "Title"},
        &validators.StringLengthInRange{Field: t.Title, Name: "Title", Min: 1, Max: 100},
    ), nil
}
```

### 3. Built-in ORM (Pop)
```go
// Database connection otomatik olarak context'te mevcut
tx, ok := c.Value("tx").(*pop.Connection)

// CRUD işlemleri
tx.Create(todo)
tx.Find(todo, id)
tx.Update(todo)
tx.Destroy(todo)
```

### 4. Migration Sistemi
```bash
# Migration oluşturma
buffalo pop generate migration create_todos

# Migration çalıştırma
buffalo pop migrate

# Migration geri alma
buffalo pop migrate down
```

## 🎯 Buffalo CLI Komutları

```bash
# Yeni proje oluşturma
buffalo new app-name --api --db-type=postgres

# Development server
buffalo dev

# Database işlemleri
buffalo pop migrate
buffalo pop migrate down
buffalo pop generate migration migration_name

# Test çalıştırma
buffalo test

# Build
buffalo build
```

## 📊 Performans Karşılaştırması

| Özellik | Buffalo | Pure Go |
|---------|---------|---------|
| **Startup Time** | ~200ms | ~50ms |
| **Memory Usage** | ~25MB | ~15MB |
| **Dependencies** | 50+ | 8 |
| **Learning Curve** | Orta | Düşük |
| **Flexibility** | Orta | Yüksek |
| **Development Speed** | Yüksek | Orta |

## 🔍 Test Etme

### 1. Todo Oluşturma
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buffalo ile Todo",
    "description": "Buffalo framework kullanarak oluşturuldu",
    "priority": "HIGH"
  }'
```

### 2. Todo Listeleme
```bash
curl http://localhost:8080/api/todos
```

### 3. Todo Güncelleme
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Güncellenmiş Todo",
    "completed": true
  }'
```

### 4. Completion Toggle
```bash
curl -X PATCH http://localhost:8080/api/todos/1/toggle
```

## 🎓 Öğrenme Notları

### Buffalo Framework Avantajları:
1. **Hızlı Geliştirme**: Auto-generated CRUD
2. **Convention over Configuration**: Standart yapı
3. **Built-in Tools**: ORM, validation, migrations
4. **Hot Reload**: Development kolaylığı

### Buffalo Framework Dezavantajları:
1. **Framework Bağımlılığı**: Buffalo kurallarına uymak zorunda
2. **Daha Fazla Dependency**: 50+ paket
3. **Daha Az Kontrol**: Framework kısıtlamaları
4. **Öğrenme Maliyeti**: Buffalo'yu öğrenmek gerekli

## 🚀 Sonuç

Bu proje, **Buffalo Framework**'ün Go web geliştirmedeki rolünü göstermektedir. Pure Go yaklaşımıyla karşılaştırıldığında:

- **Buffalo**: Hızlı geliştirme, standart yapı, built-in özellikler
- **Pure Go**: Tam kontrol, yüksek performans, minimal dependencies

Her iki yaklaşımın da kendi avantajları vardır ve projenin ihtiyaçlarına göre seçim yapılmalıdır.
