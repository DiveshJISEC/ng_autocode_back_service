# Technical Guide: Creating a New Web Service

## Overview
This guide outlines the process of creating a new web service using our standard architecture based on Go and PostgreSQL. The guide follows the structure implemented in the NG Backend Server project.

## Table of Contents
1. [Project Structure](#project-structure)
2. [Prerequisites](#prerequisites)
3. [Implementation Steps](#implementation-steps)
4. [Configuration](#configuration)
5. [Build and Deployment](#build-and-deployment)
6. [Testing](#testing)

## Project Structure
```
project/
├── app/               # Application resources
│   ├── bin/          # Compiled binaries
│   ├── conf/         # Configuration files
│   ├── docs/         # Documentation
│   └── scripts/      # Build scripts
├── cmd/              # Main application entry points
├── common/           # Shared utilities
├── internal/         # Core business logic
│   ├── api/         # API layer
│   ├── core/        # Business logic
│   ├── infra/       # Infrastructure layer
│   └── startup/     # Application startup
├── pkg/             # Reusable packages
└── web/            # Frontend resources
```

## Prerequisites
1. Go 1.24.3 or higher
2. PostgreSQL database
3. IDE (VS Code recommended)
4. Git for version control

## Implementation Steps

### 1. Module Structure and Hierarchy
The application follows a strict layered architecture with interface-implementation pairs at each layer:

```
internal/core/
├── moduleFolder/                # e.g., fd/
    ├── portModule.go           # Module interface definition
    ├── serviceGroup1/          # e.g., book/
    │   ├── portBook.go        # Group interface & implementation
    │   └── service/           
    │       ├── portServices.go # Service interface
    │       ├── bookServices.go # Service implementation
    │       ├── model/         
    │       │   ├── db/        # Database models (GORM entities)
    │       │   └── dto/       # Data Transfer Objects (API contracts)
    │       └── repo/         
    │           ├── portRepo.go # Repository interface
    │           └── bookRepo.go # Repository implementation
    └── serviceGroup2/          # e.g., list/
        └── ... (same pattern)
```

Key Architecture Rules:
1. **Interface-Implementation Pattern**:
   - Service and Repository layers have interface-implementation pairs:
     - `port*.go` - Interface definition
     - `*Services.go`/`*Repo.go` - Implementation
   - Group layer combines interface and implementation in single `port*.go` file
   - New functionality is added to existing files, not new files

2. **Service Groups**:
   - Independent functional units (e.g., 'book', 'list')
   - Each has complete set of interface-implementation pairs
   - New features either:
     - Extend existing service group (modify implementation files)
     - Create new service group (new folder with complete structure)

3. **Data Model Separation**:
   - `dto/` - API contracts and request/response models
   - `db/` - Database entities with GORM tags

### 2. Extending vs Creating Services

1. **Extending Existing Services**:
   - Add new methods to existing interface files (`port*.go`)
   - Implement in corresponding implementation files
   - Never create new files for new operations

2. **Creating New Service Group**:
   - Create new folder under module (e.g., `fd/transaction/`)
   - Follow complete interface-implementation pattern
   - Implement all required layers (web, service, repo)

3. **Data Models**:
   - Place DTOs in `model/dto/` for API contracts
   - Place DB models in `model/db/` for GORM entities
   - Example DTO:
     ```go
     type RequestDTO struct {
         UserId      string `json:"FML_USER_ID" binding:"required"`
         SessionCode string `json:"FML_SSSN_ID" binding:"required"`
     }
     ```

### 3. Layer Implementation Example (FD Order Book)

Each layer follows interface-implementation pattern:

1. **Web/Group Layer**:
```go
// portBook.go - Both interface and implementation
type BookServices interface {
    GetFDOrderBook(c *gin.Context)
}

func (f *bookServices) GetFDOrderBook(c *gin.Context) {
    // HTTP handling & service layer calls
}
```

2. **Service Layer**:
```go
// portServices.go - Interface
type ServiceLayer interface {
    GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) ([]*dto.FDOrderBookResponse, error)
}

// bookServices.go - Implementation
func (g *serviceItem) GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) ([]*dto.FDOrderBookResponse, error) {
    // Business logic & data layer calls
}
```

3. **Repository Layer**:
```go
// portRepo.go - Interface
type DataLayer interface {
    GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) ([]*dto.FDOrderBookResponse, error)
}

// bookRepo.go - Implementation
func (g *dbSt) GetFDOrderBook(c context.Context, request *dto.FDOrderBookRequest) ([]*dto.FDOrderBookResponse, error) {
    // Database operations
}
```

Key Points:
- Each layer has its interface and implementation files
- New methods are added to existing files, not new files
- Clear separation of concerns between layers
- Data transformations happen at layer boundaries

### 5. Add API Endpoints
1. Update `internal/api/router.go`
2. Add new routes in appropriate group
3. Implement request validation
4. Add proper error handling

Example Route Addition:
```go
v1 := router.Group("/v1")
module := v1.Group("/module")
{
    module.POST("/endpoint", moduleGrp.HandleEndpoint)
}
```

### 6. Frontend Integration
1. Add new UI components in `web/`
2. Update API endpoints in `script.js`
3. Add necessary styles in `styles.css`

## Configuration

### 1. Environment Configuration
Update configuration files in `app/conf/`:
```yaml
server:
  host: 0.0.0.0
  port: 8086
  
repo:
  databases:
    postgres:
      read:
        user: postgres
        password: admin
        host: localhost
        port: 5432
        db: ServiceDB
```

### 2. Logging Configuration
Configure logging in `pkg/logger/logger.go`:
```yaml
log:
  path: "log/app.log"
  level: -1  # Debug level
```

## Build and Deployment

### 1. Local Development
```powershell
# Debug build
.\app\scripts\buildWinDebug.bat

# Run application
.\run.bat
```

### 2. Production Deployment
```powershell
# Release build
.\app\scripts\buildWinRelease.bat
```

### 3. Docker Deployment
- Use provided Docker configurations in `app/docker/`
- Build and run using Docker Compose

## Testing

### 1. Unit Testing
- Create tests in `test/` directory
- Follow Go testing conventions
- Test all public interfaces

### 2. Integration Testing
- Test API endpoints
- Verify database operations
- Check error handling

### Best Practices
1. **Code Organization**
   - Follow clean architecture principles
   - Maintain separation of concerns
   - Use dependency injection

2. **Error Handling**
   - Use common error types from `common/http/apiError.go`
   - Implement proper logging
   - Return appropriate HTTP status codes

3. **Security**
   - Validate all inputs
   - Implement proper authentication/authorization
   - Use HTTPS in production
   - Sanitize database queries

4. **Performance**
   - Implement caching where appropriate
   - Use database indexes
   - Monitor resource usage

## Monitoring and Maintenance

### 1. Health Checks
- Implement `/health` endpoint
- Monitor system metrics
- Set up logging and alerting

### 2. Versioning
- Use semantic versioning
- Document API changes
- Maintain backward compatibility

## Conclusion
Following this guide ensures consistency across services and maintains the quality standards of our application architecture. Always refer to existing implementations in the codebase as examples.

For detailed API documentation, refer to `app/docs/api_desc.txt`.
