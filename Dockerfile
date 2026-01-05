# =========================
# Étape 1 : Build
# =========================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod ./

# Télécharger les dépendances
RUN go mod download

# Copier tout le code source
COPY . .

# Compiler le binaire
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o virements ./cmd/virements

# =========================
# Étape 2 : Runtime service
# =========================
FROM alpine:3.19

# Créer un utilisateur non-root
RUN adduser -D appuser

WORKDIR /app

# Copier le binaire depuis le builder
COPY --from=builder /app/virements .

USER appuser

EXPOSE 8080

CMD ["./virements"]
