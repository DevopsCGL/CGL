# Étape 1 : Compilation
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Ajouter les certificats SSL
RUN apk --no-cache add ca-certificates

# Copier les fichiers Go (mais pas .env)
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Compiler l’application Go
RUN go build -o app .

# Étape 2 : Image finale plus légère
FROM alpine:latest  

WORKDIR /app

# Ajouter les certificats SSL (nécessaire pour certaines connexions)
RUN apk --no-cache add ca-certificates

# Copier uniquement l’exécutable depuis l’étape 1
COPY --from=builder /app/app .

# Copie le fichier .env pour qu’il soit accessible dans le conteneur
COPY --from=builder /app/.env .env

CMD ["./app"]