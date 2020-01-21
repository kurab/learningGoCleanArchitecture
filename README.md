## Directories
```
├── config
│   ├── config.go
│   └── config.yml
├── domain
│   └── model
│       ├── property.go
│       └── user.go
├── infrastructure
│   ├── api
│   │   ├── handler
│   │   │   ├── appHandler.go
│   │   │   ├── propertyHandler.go
│   │   │   └── userHandler.go
│   │   ├── router
│   │   │   └── router.go
│   │   └── validator
│   │       └── validator.go
│   └── datastore
│       ├── dbMySQL.go
│       ├── dbPostgreSQL.go
│       ├── propertyRepository.go
│       └── userRepository.go
├── interface
│   ├── controllers
│   │   ├── propertyController.go
│   │   └── userController.go
│   └── presenters
│       └── propertyPresenter.go
├── registry
│   └── registry.go
├── usecase
│   ├── presenter
│   │   └── propertyPresenter.go
│   ├── repository
│   │   ├── propertyRepository.go
│   │   └── userRepository.go
│   └── service
│       ├── propertyService.go
│       └── userService.go
└── main.go (UPDATED!)
```

## Database docker file
```yml:docker-compose.yml
version: '3'
services:
    myserver:
        image: mysql:5.7
        ports:
            - 3306:3306
        hostname: mysql57_server
        command: --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
        environment:
            MYSQL_DATABASE: goSample
            MYSQL_ROOT_PASSWORD: secret
            MYSQL_USER: user
            MYSQL_PASSWORD: secret
        volumes:
             - gosample_local_data:/var/lib/mysql
    pgserver:
        image: postgres:12.1-alpine
        ports:
            - 5432:5432
        hostname: ps121_server
        environment:
            POSTGRES_DB: goSample
            POSTGRES_USER: user
            POSTGRES_PASSWORD: secret
            POSTGRES_INITDB_ARGS: "--encoding=UTF-8 --locale=ja_JP.UTF-8"
        volumes:
             - gosamplepg_local_data:/var/lib/postgresql/data
        restart: always
        user: root

volumes:
    gosample_local_data:
        driver: local
    gosamplepg_local_data:
        driver: local
```

