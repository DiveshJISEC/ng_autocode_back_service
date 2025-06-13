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

### 1. Create Module Structure
- Create a new module under `internal/core/`
- Follow the existing pattern:
  ```
  moduleFolder/
  ├── portModule.go       # Module interface
  ├── service/           # Service implementation
  │   ├── model/        # Data models
  │   │   ├── db/      # Database models
  │   │   └── dto/     # Data transfer objects
  │   └── repo/        # Repository layer
  ```

### 2. Define Data Models
1. Create DTO models in `model/dto/`
2. Create DB models in `model/db/`
3. Implement GORM tags for database mapping

Example DTO:
```go
type RequestDTO struct {
    UserId      string `json:"FML_USER_ID" binding:"required"`
    SessionCode string `json:"FML_SSSN_ID" binding:"required"`
}
```

### 3. Implement Repository Layer
1. Create repository interface in `repo/portRepo.go`
2. Implement database operations in `repo/implementationRepo.go`
3. Use GORM for database operations

### 4. Implement Service Layer
1. Define service interface in `portServices.go`
2. Implement business logic in service files
3. Handle data transformation between DTO and DB models

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
