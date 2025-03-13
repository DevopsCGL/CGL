package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5" // Importer chi pour gérer les routes
	"github.com/joho/godotenv" // Importer godotenv
)

func main() {
	// Charger les variables d'environnement depuis le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur de chargement du fichier .env")
	}

	// Vérification que la variable d'environnement est bien lue
	port := os.Getenv("PORT") // Récupère la variable d'environnement PORT
	if port == "" {
		log.Fatal("PORT non défini dans le fichier .env")
	}

	log.Println("🎉 Le fichier .env a été chargé avec succès")

	// Créer un nouveau routeur avec chi
	router := chi.NewRouter()

	// Ajouter une route de test
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🚀 L'application fonctionne sur le port " + port))
	})

	// Démarrer le serveur HTTP
	log.Printf("🚀 Le serveur tourne sur le port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
