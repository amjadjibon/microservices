# microservices
Golang backend microservices

## Auth Service

### Environment
```bash
cp auth/conf/.env.example auth/conf/.env
```

### Update environment variables
```bash
DATABASE_DSN=""
```

### Migrate
```bash
make auth-migrate-up
```

### Run
```bash
make auth-run
```
