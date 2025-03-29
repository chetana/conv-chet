package app

import (
	"context"
	"log"
	"os" // Import du package os

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitializeFirestore() {
	ctx := context.Background()

	// Récupérer le chemin du fichier de clé depuis une variable d'environnement
	keyPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if keyPath == "" {
		log.Fatal("GOOGLE_APPLICATION_CREDENTIALS variable d'environnement non définie")
	}

	// Utiliser un fichier de clé de compte de service
	sa := option.WithCredentialsFile(keyPath)

	// Initialiser l'application Firebase
	conf := &firebase.Config{
		ProjectID: os.Getenv("GCP_PROJECT_ID"), // Récupérer l'ID du projet GCP
	}

	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de l'application Firebase: %v\n", err)
	}

	// Initialiser le client Firestore
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de Firestore: %v\n", err)
	}

	log.Println("Firestore initialisé avec succès")
}
