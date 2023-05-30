## Setting Up

Make sure you have a postgres database running, set via:

```
export DATABASE_URL=postgresql://postgres:postgres@localhost:5432/example
```

Run the following:

```
go mod download
go run github.com/prisma/prisma-client-go migrate dev --name init
```

