# Utiliser une image officielle de Node.js
FROM node:14

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers package.json et package-lock.json
COPY package.json .
COPY package-lock.json .

# Installer les dépendances
RUN npm install

# Copier tous les fichiers de l'application dans le conteneur
COPY . .

# Exposer le port 3000 pour l'application React
EXPOSE 3000

# Lancer l'application React
CMD ["npm", "start"]
