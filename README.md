<<<<<<< README.md
# Transactions - Virements Service (Go)

## Run locally
```bash
# go run ./cmd/server
=======
# CI Go Virements - Documentation CI/CD

Ce projet contient un service Go pour gérer des virements bancaires et un pipeline CtLab pour **build, test et linting automatique**.

## Structure du projet

CI_Go_test/
├─ cmd/
│ └─ virements/
│ └─ main.go # Point d'entrée du service
├─ internal/
│ └─ virment/
│ ├─ model.go # Logique métier des virements
│ └─ model_test.go # Tests unitaires du modèle
├─ go.mod # Module Go
├─ go.sum # Dépendances Go
├─ Dockerfile # Build + runtime
├─ docker-compose.yml # Lancer service et tests localement
├─ .gitlab-ci.yml # Pipeline CI/CD
└─ .golangci.yml # Configuration du linting

🚀 Fonctionnalités

Création d’un virement

Consultation d’un virement par ID

Statuts du virement (PENDING, COMPLETED, FAILED…)

API REST en Go

Tests unitaires avec go test

Conteneurisation avec Docker

---

## 2️⃣ GitLab CI/CD

Le pipeline est défini dans `.gitlab-ci.yml` et contient trois jobs principaux :

### 2.1 Build

- Compile le projet Go.
- Vérifie les dépendances (`go mod tidy`).
- Conserve `go.sum` comme artifact pour les jobs suivants.
- Se déclenche uniquement sur :
  - Merge requests
  - Branche `main`

### 2.2 Unit Tests

- Exécute tous les tests unitaires Go (`go test ./...`).
- Génère un **rapport de couverture** (`coverage.out`) et un **rapport JUnit XML** (report.xml`) pour GitLab.
- Dépend du job `build`.

### 2.3 GolangCI Lint

- Analyse statique du code Go pour détecter :
  - Erreurs potentielles (`govet`, `errcheck`, `staticcheck`)
  - Code inutilisé (`unused`)
  - Problèmes de style (`revive`)
  - Simplifications possibles (`gosimple`)
  - Assignations inefficaces (`ineffassign`)
- Configuration centralisée dans `.golangci.yml`.
- Génère un **rapport JSON** (`lint-report.json`) pour GitLab.
- Dépend également du job `build`.

---

## 3️⃣ Docker

### 3.1 Dockerfile

- **Multi-stage build** :
  1. **Builder** : compile le binaire Go `virements`.
  2. **Runtime** : image Alpine légère, exécute le binaire en utilisateur non-root.
- Port exposé : `8080`.

### 3.2 docker-compose.yml

- Permet de lancer :
  - **virements-service** : exécute le service Go
  - **virements-tests** : exécute les tests unitaires dans un conteneur

Commandes principales :

```bash
docker-compose build   # Construire les conteneurs
docker-compose up      # Lancer le service et les tests

# 1️⃣ Vérifier les dépendances et compiler le projet
go mod tidy
go build ./cmd/virements

# 2️⃣ Exécuter les tests unitaires
go test ./...

# 3️⃣ Lancer le linting avec GolangCI-Lint
golangci-lint run

# 4️⃣ Docker - Construire les conteneurs
docker-compose build

L’API démarre sur :
http://localhost:8080

📡 Endpoints API

{
  "from": "ACC1",
  "to": "ACC2",
  "currency": "MAD",
}

🔍 Récupérer un virement
GET /api/v1/virements/{id}


⚠️ Utiliser l’ID retourné lors de la création.
  "reference": "REF123"
  "amount": 100,
POST /api/v1/virements
➕ Créer un virement





>>>>>>> README.md
