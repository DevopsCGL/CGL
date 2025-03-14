package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

func main() {
	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur de chargement du fichier .env")
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")

	// Vérifier les variables d'environnement
	if port == "" || dbURL == "" {
		log.Fatal("PORT ou DATABASE_URL non défini dans le fichier .env")
	}

	log.Println("✅ Fichier .env chargé avec succès")

	// Connexion à la base PostgreSQL
	db, err = sqlx.Connect("pgx", dbURL)
	if err != nil {
		log.Fatal("❌ Erreur de connexion à la base :", err)
	}
	log.Println("🎉 Connexion à la base réussie")

	// Créer le routeur
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Route pour récupérer toutes les tâches dans la base
	router.Get("/getAllTaches", func(w http.ResponseWriter, r *http.Request) {
		var taches []struct {
			ID          int    `db:"id" json:"id"`
			Date        string `db:"date" json:"date"`
			Description string `db:"description" json:"description"`
		}

		// Récupérer les tâches depuis la base de données
		err := db.Select(&taches, "SELECT * FROM tache")
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des tâches", http.StatusInternalServerError)
			return
		}

		// Définir le header pour spécifier que la réponse est en JSON
		w.Header().Set("Content-Type", "application/json")

		// Encoder les tâches en JSON et les envoyer dans la réponse
		err = json.NewEncoder(w).Encode(taches)
		if err != nil {
			http.Error(w, "Erreur lors de l'encodage JSON", http.StatusInternalServerError)
		}
	})

	// Route pour ajouter une tâche
	router.Post("/addTache", func(w http.ResponseWriter, r *http.Request) {
		// Définir une structure pour la tâche
		type Tache struct {
			Date        string `json:"date"`
			Description string `json:"description"`
		}

		// Décoder le corps de la requête
		var tache Tache
		err := json.NewDecoder(r.Body).Decode(&tache)
		if err != nil {
			http.Error(w, "Erreur de décodage du JSON", http.StatusBadRequest)
			return
		}

		// Insérer la nouvelle tâche dans la base de données
		_, err = db.Exec("INSERT INTO tache (date, description) VALUES ($1, $2)", tache.Date, tache.Description)
		if err != nil {
			http.Error(w, "Erreur lors de l'ajout de la tâche", http.StatusInternalServerError)
			return
		}

		// Répondre avec un message de succès
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Tâche ajoutée avec succès"})
	})

	router.Delete("/deleteTache/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		log.Println("Tentative de suppression de la tâche avec ID:", id)

		result, err := db.Exec("DELETE FROM tache WHERE id = $1", id) // ✅ Remplacement de '?' par '$1'
		if err != nil {
			log.Println("Erreur SQL:", err)
			http.Error(w, "Erreur lors de la suppression de la tâche", http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			log.Println("Aucune tâche trouvée avec cet ID")
			http.Error(w, "Tâche non trouvée", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Tâche supprimée avec succès"})
	})

	// Lancer le serveur
	log.Printf("🚀 Serveur lancé sur le port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
