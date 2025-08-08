# Task Tracker CLI

## Rules

Règles à respecter dans ce projet

- Pas de code dans main.go sauf la commande.
- Chaque fonction doit faire UNE chose.
- Pas de global state.
- Pas de fmt.Println sauvage, tout doit être testable.
- On testera dès qu’on aura une logique métier isolée.
- On documente avec des commentaires clairs au-dessus de chaque fonction publique.

## Architecture

- `task/`: contient le code métier lié aux tasks
- `cmd/`: contient le point d'entrée CLI
- `data/`: contient les données

## Ordre de réalisation

- Définition des besoins

  - Création du modèle `Task`. Création de la struct puis du constructeur `NewTask(description string) *Taks`

- Ajout des besoin métiers

  - `MarkDone()`,`SetStatus()` .. .
  - Les méthodes ne font pas de print ou d'écrire, elles doivent être tester

- Ajout du système de persistence
  -Stockage en JSON

- Implémentation de la CLI
  - dans le fichier `cmd/main`
