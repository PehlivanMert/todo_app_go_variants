# Todo App REST API - Türkçe Dokümantasyon

Katmanlı mimari ile geliştirilmiş, production-ready Go REST API projesi.

## 🚀 Hızlı Erişim

- [Proje Yapısı](#proje-yapısı)
- [Kurulum](#kurulum)
- [API Endpoints](#api-endpoints)
- [Docker ile Çalıştırma](#docker-ile-çalıştırma)
- [Swagger Dokümantasyonu](#swagger-dokümantasyonu)
- [Go Projesi Detaylı Açıklama](#go-projesi-detaylı-açıklama)

## 📁 Proje Yapısı

```
todo_app_pure_go/
├── main.go                     # Ana uygulama giriş noktası
├── main_test.go               # Test dosyası
├── go.mod & go.sum            # Go modül dosyaları
├── .env                       # Ortam değişkenleri
├── Dockerfile                 # Docker image
├── docker-compose.yml         # Docker Compose
├── README.md                  # İngilizce dokümantasyon
├── README_TR.md               # Türkçe dokümantasyon
├── .gitignore                 # Git ignore
├── config/
│   └── database.go            # Veritabanı konfigürasyonu
├── models/
│   └── todo.go                # Todo veri modeli
├── dto/
│   ├── todo_request.go        # İstek DTO'ları
│   └── todo_response.go       # Yanıt DTO'ları
├── repository/
│   ├── todo_repository.go     # Repository arayüzü
│   └── todo_repository_impl.go # Repository uygulaması
├── service/
│   ├── todo_service.go        # Service arayüzü
│   └── todo_service_impl.go   # Service uygulaması
├── controller/
│   └── todo_controller.go     # HTTP controller
├── middleware/
│   ├── cors.go                # CORS middleware
│   └── logger.go              # Logging middleware
├── utils/
│   ├── response.go            # Yanıt yardımcıları
│   └── validator.go           # Validasyon yardımcıları
├── routes/
│   └── routes.go              # Route tanımları
└── docs/
    └── docs.go                # Swagger dokümantasyonu
```

## 🛠️ Kurulum

### Gereksinimler
- Go 1.21 veya üzeri
- PostgreSQL veritabanı
- Git
- Docker (opsiyonel, geliştirme için)

### Adım 1: Projeyi İndirin
```bash
git clone
cd todo_app_pure_go
```

### Adım 2: Bağımlılıkları Yükleyin
```bash
go mod tidy
```

### Adım 3: Ortam Değişkenlerini Ayarlayın
`.env` dosyası oluşturun:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todoapp
DB_SSLMODE=disable
PORT=8080
GIN_MODE=debug
```

### Adım 4: PostgreSQL Veritabanını Kurun
```sql
CREATE DATABASE todoapp;
```

### Adım 5: Uygulamayı Çalıştırın
```bash
go run main.go
```

Uygulama `http://localhost:8080` adresinde çalışacaktır.

### Adım 6: Swagger Dokümantasyonuna Erişin
```
http://localhost:8080/swagger/index.html
```

## 📚 API Endpoints

### Swagger UI
İnteraktif API dokümantasyonuna erişmek için:
```
http://localhost:8080/swagger/index.html
```

### Base URL
```
http://localhost:8080/api
```

### Endpoints

#### 1. Tüm Todo'ları Listele
```http
GET /api/todos
```

**Query Parametreleri:**
- `completed` (opsiyonel): Tamamlanma durumuna göre filtrele (true/false)
- `priority` (opsiyonel): Önceliğe göre filtrele (LOW, MEDIUM, HIGH)
- `limit` (opsiyonel): Sayfa başına öğe sayısı (varsayılan: 10, maksimum: 100)
- `offset` (opsiyonel): Atlanacak öğe sayısı (varsayılan: 0)

**Örnek:**
```bash
curl -X GET "http://localhost:8080/api/todos?completed=false&priority=HIGH&limit=5&offset=0"
```

#### 2. ID'ye Göre Todo Getir
```http
GET /api/todos/{id}
```

**Örnek:**
```bash
curl -X GET "http://localhost:8080/api/todos/1"
```

#### 3. Todo Oluştur
```http
POST /api/todos
```

**İstek Gövdesi:**
```json
{
  "title": "Proje dokümantasyonunu tamamla",
  "description": "Todo API için kapsamlı dokümantasyon yaz",
  "priority": "HIGH"
}
```

**Örnek:**
```bash
curl -X POST "http://localhost:8080/api/todos" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Proje dokümantasyonunu tamamla",
    "description": "Todo API için kapsamlı dokümantasyon yaz",
    "priority": "HIGH"
  }'
```

#### 4. Todo Güncelle
```http
PUT /api/todos/{id}
```

**İstek Gövdesi:**
```json
{
  "title": "Güncellenmiş başlık",
  "description": "Güncellenmiş açıklama",
  "completed": true,
  "priority": "MEDIUM"
}
```

#### 5. Todo Sil
```http
DELETE /api/todos/{id}
```

#### 6. Todo Tamamlanma Durumunu Değiştir
```http
PATCH /api/todos/{id}/toggle
```

## 🐳 Docker ile Çalıştırma

### Hızlı Başlangıç
```bash
# Tüm servisleri başlat (uygulama + veritabanı)
docker compose up -d

# Sadece veritabanını başlat
docker compose up -d postgres

# Tüm servisleri durdur
docker compose down

# Servisleri durdur ve veritabanı verilerini sil
docker compose down -v

# Logları görüntüle
docker compose logs -f

# Yeniden derle ve başlat
docker compose up --build -d
```

### Manuel Docker Komutları
```bash
# PostgreSQL container'ını başlat
docker run -d --name postgres-todo \
  -e POSTGRES_DB=todoapp \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -p 5432:5432 \
  postgres:15-alpine

# Container'ı durdur
docker stop postgres-todo

# Container'ı sil
docker rm postgres-todo

# Tüm container'ları temizle
docker system prune -a
```

### Uygulamayı Test Edin
```bash
curl http://localhost:8080/health
```

## 📖 Swagger Dokümantasyonu

### Swagger UI Erişimi
Swagger UI'ya erişmek için:
```
http://localhost:8080/swagger/index.html
```

### Swagger Özellikleri
- ✅ **İnteraktif Test**: Endpoint'leri doğrudan tarayıcıdan test edebilme
- ✅ **İstek/Yanıt Örnekleri**: Örnek istek ve yanıtları görme
- ✅ **Parametre Doğrulama**: Gerekli ve opsiyonel parametreleri anlama
- ✅ **Hata Kodları**: Tüm olası hata yanıtlarını görme
- ✅ **Kimlik Doğrulama**: Gelecekteki kimlik doğrulama için hazır

### Swagger ile Yapabilecekleriniz
1. **İnteraktif Dokümantasyon Görüntüleme**: `http://localhost:8080/swagger/index.html` adresini ziyaret edin
2. **API Endpoint'lerini Test Etme**: Swagger UI kullanarak tüm endpoint'leri doğrudan test edin
3. **İstek/Yanıt Şemalarını Görme**: Detaylı istek ve yanıt formatlarını inceleyin
4. **OpenAPI Spec İndirme**: OpenAPI spesifikasyonu `/swagger/doc.json` adresinde mevcut

## 🏗️ Go Projesi Detaylı Açıklama

### Neden Bu Yapıyı Kullandık?

#### 1. **Katmanlı Mimari (Layered Architecture)**
```
Controller → Service → Repository → Database
```

**Neden?**
- **Sorumluluk Ayrımı**: Her katmanın belirli bir sorumluluğu var
- **Test Edilebilirlik**: Her katmanı ayrı ayrı test edebiliriz
- **Bakım Kolaylığı**: Değişiklikler tek bir yerde yapılır
- **Yeniden Kullanılabilirlik**: Katmanlar başka projelerde kullanılabilir

#### 2. **SOLID Prensipleri**

**S - Single Responsibility (Tek Sorumluluk)**
```go
// Her struct/func tek bir işi yapar
type TodoRepository interface {
    Create(todo *models.Todo) (*models.Todo, error)
    GetByID(id uint) (*models.Todo, error)
    // ...
}
```

**O - Open/Closed (Açık/Kapalı)**
```go
// Interface ile extension'a açık, modification'a kapalı
type TodoService interface {
    CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error)
    // Yeni metodlar eklenebilir ama mevcut kod değişmez
}
```

**L - Liskov Substitution (Liskov Yerine Geçme)**
```go
// Interface implementation'ları birbirinin yerine geçebilir
todoRepo := repository.NewTodoRepository(db) // PostgreSQL
// todoRepo := repository.NewMockRepository() // Test için
```

**I - Interface Segregation (Arayüz Ayrımı)**
```go
// Küçük, özel interface'ler
type TodoRepository interface {
    Create(todo *models.Todo) (*models.Todo, error)
    GetByID(id uint) (*models.Todo, error)
    // Sadece todo işlemleri
}
```

**D - Dependency Inversion (Bağımlılık Tersine Çevirme)**
```go
// Yüksek seviye modüller düşük seviye modüllere bağımlı değil
type TodoController struct {
    todoService service.TodoService // Interface'e bağımlı
}
```

#### 3. **Dependency Injection (Bağımlılık Enjeksiyonu)**

```go
// main.go'da bağımlılıkları enjekte ediyoruz
func main() {
    // Database bağlantısı
    db, err := config.ConnectDatabase(dbConfig)
    
    // Repository oluştur
    todoRepo := repository.NewTodoRepository(db)
    
    // Service oluştur
    todoService := service.NewTodoService(todoRepo)
    
    // Controller oluştur
    todoController := controller.NewTodoController(todoService)
}
```

**Neden?**
- **Loose Coupling**: Katmanlar arası gevşek bağlantı
- **Testability**: Mock objeler kullanabiliriz
- **Flexibility**: Farklı implementation'lar kullanabiliriz

#### 4. **DTO Pattern (Data Transfer Object)**

```go
// Request DTO - Gelen veriyi temsil eder
type CreateTodoRequest struct {
    Title       string   `json:"title" validate:"required,min=1,max=100"`
    Description *string  `json:"description" validate:"omitempty,max=500"`
    Priority    Priority `json:"priority" validate:"omitempty,oneof=LOW MEDIUM HIGH"`
}

// Response DTO - Giden veriyi temsil eder
type TodoResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description *string   `json:"description"`
    Completed   bool      `json:"completed"`
    Priority    Priority  `json:"priority"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

**Neden?**
- **Güvenlik**: İç veri yapısını gizler
- **Esneklik**: API değişikliklerini kolaylaştırır
- **Validasyon**: Gelen veriyi doğrular

#### 5. **Repository Pattern**

```go
// Interface - Soyutlama
type TodoRepository interface {
    Create(todo *models.Todo) (*models.Todo, error)
    GetByID(id uint) (*models.Todo, error)
    GetAll(completed *bool, priority *Priority, limit, offset int) ([]*models.Todo, error)
    Update(id uint, todo *models.Todo) (*models.Todo, error)
    Delete(id uint) error
    ToggleComplete(id uint) (*models.Todo, error)
    GetTotalCount(completed *bool, priority *Priority) (int64, error)
}

// Implementation - Somut uygulama
type TodoRepositoryImpl struct {
    db *gorm.DB
}
```

**Neden?**
- **Database Independence**: Veritabanı değişikliklerini kolaylaştırır
- **Testing**: Mock repository kullanabiliriz
- **Abstraction**: Veritabanı detaylarını gizler

#### 6. **Service Layer (İş Mantığı Katmanı)**

```go
type TodoServiceImpl struct {
    todoRepo repository.TodoRepository
}

func (s *TodoServiceImpl) CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
    // 1. Validasyon
    if errors := utils.ValidateStruct(req); len(errors) > 0 {
        return nil, errors.New("validation failed: " + errors[0])
    }
    
    // 2. İş mantığı
    if req.Priority == "" {
        req.Priority = models.MEDIUM
    }
    
    // 3. Model oluştur
    todo := &models.Todo{
        Title:       req.Title,
        Description: req.Description,
        Priority:    req.Priority,
        Completed:   false,
    }
    
    // 4. Repository'ye kaydet
    createdTodo, err := s.todoRepo.Create(todo)
    if err != nil {
        return nil, err
    }
    
    // 5. Response DTO'ya dönüştür
    return s.todoToResponse(createdTodo), nil
}
```

**Neden?**
- **Business Logic**: İş kurallarını burada uygularız
- **Validation**: Veri doğrulama
- **Transformation**: DTO dönüşümleri
- **Transaction Management**: İşlem yönetimi

#### 7. **Controller Layer (HTTP Katmanı)**

```go
type TodoController struct {
    todoService service.TodoService
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
    // 1. Request'i parse et
    var req dto.CreateTodoRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
        return
    }
    
    // 2. Service'i çağır
    todo, err := tc.todoService.CreateTodo(&req)
    if err != nil {
        utils.InternalServerErrorResponse(c, "Failed to create todo: "+err.Error())
        return
    }
    
    // 3. Response gönder
    utils.CreatedResponse(c, todo, "Todo created successfully")
}
```

**Neden?**
- **HTTP Handling**: HTTP isteklerini yönetir
- **Request/Response**: İstek ve yanıt formatlarını yönetir
- **Error Handling**: HTTP hata kodlarını döner
- **Middleware Integration**: Middleware'leri kullanır

#### 8. **Middleware Pattern**

```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        
        c.Next()
    }
}
```

**Neden?**
- **Cross-cutting Concerns**: Tüm endpoint'lerde ortak işlemler
- **Reusability**: Tekrar kullanılabilir
- **Separation of Concerns**: Ana iş mantığından ayrı

#### 9. **Error Handling (Hata Yönetimi)**

```go
// Standardized error responses
func BadRequestResponse(c *gin.Context, message string) {
    ErrorResponse(c, http.StatusBadRequest, message, "Bad Request")
}

func NotFoundResponse(c *gin.Context, message string) {
    ErrorResponse(c, http.StatusNotFound, message, "Not Found")
}

func InternalServerErrorResponse(c *gin.Context, message string) {
    ErrorResponse(c, http.StatusInternalServerError, message, "Internal Server Error")
}
```

**Neden?**
- **Consistency**: Tutarlı hata yanıtları
- **Debugging**: Hata ayıklama kolaylığı
- **User Experience**: Kullanıcı dostu hata mesajları

#### 10. **Configuration Management (Konfigürasyon Yönetimi)**

```go
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func LoadDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnv("DB_PORT", "5432"),
        User:     getEnv("DB_USER", "postgres"),
        Password: getEnv("DB_PASSWORD", "password"),
        DBName:   getEnv("DB_NAME", "todoapp"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}
```

**Neden?**
- **Environment Flexibility**: Farklı ortamlar için farklı ayarlar
- **Security**: Hassas bilgileri environment'tan alır
- **Maintainability**: Konfigürasyon değişikliklerini kolaylaştırır

### Go'da Önemli Kavramlar

#### 1. **Interface'ler**
```go
// Interface - Sözleşme tanımlar
type TodoService interface {
    CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error)
}

// Implementation - Sözleşmeyi yerine getirir
type TodoServiceImpl struct {
    todoRepo repository.TodoRepository
}

func (s *TodoServiceImpl) CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
    // Implementation
}
```

#### 2. **Pointer Receivers**
```go
// Pointer receiver - Struct'ı değiştirebilir
func (s *TodoServiceImpl) CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
    // s.todoRepo'yu değiştirebilir
}

// Value receiver - Struct'ı kopyalar
func (s TodoServiceImpl) GetInfo() string {
    // s'yi değiştiremez, kopya üzerinde çalışır
}
```

#### 3. **Error Handling**
```go
// Go'da error'lar değer olarak döner
func CreateTodo(todo *models.Todo) (*models.Todo, error) {
    if err := db.Create(todo).Error; err != nil {
        return nil, err // Error'u yukarı ilet
    }
    return todo, nil
}

// Error'u kontrol et
if err != nil {
    // Hata durumunu handle et
    return err
}
```

#### 4. **Goroutines ve Channels**
```go
// Goroutine - Concurrent execution
go func() {
    log.Printf("Server starting on port %s", port)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Failed to start server: %v", err)
    }
}()

// Channel - Goroutine'ler arası iletişim
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit // Signal bekle
```

#### 5. **Struct Tags**
```go
type Todo struct {
    ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    Title       string    `json:"title" gorm:"not null;size:100"`
    Description *string   `json:"description" gorm:"size:500"`
    Completed   bool      `json:"completed" gorm:"default:false"`
    Priority    Priority  `json:"priority" gorm:"type:varchar(10);default:'MEDIUM'"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
```

**Neden?**
- **JSON Serialization**: JSON formatını belirler
- **Database Mapping**: Veritabanı kolonlarını belirler
- **Validation**: Validasyon kurallarını belirler

### Test Etme

```bash
# Tüm testleri çalıştır
go test -v

# Coverage ile test et
go test -cover

# Benchmark testleri
go test -bench=.
```

### Production'a Hazırlık

1. **Environment Variables**: Hassas bilgileri environment'tan al
2. **Logging**: Structured logging kullan
3. **Health Checks**: Sağlık kontrolü endpoint'leri
4. **Graceful Shutdown**: Düzgün kapatma
5. **Error Handling**: Kapsamlı hata yönetimi
6. **Validation**: Gelen veriyi doğrula
7. **Security**: CORS, rate limiting, authentication
8. **Monitoring**: Metrics ve monitoring
9. **Documentation**: API dokümantasyonu
10. **Docker**: Containerization

Bu yapı sayesinde:
- ✅ **Maintainable**: Bakımı kolay kod
- ✅ **Testable**: Test edilebilir kod
- ✅ **Scalable**: Ölçeklenebilir kod
- ✅ **Readable**: Okunabilir kod
- ✅ **Flexible**: Esnek kod

