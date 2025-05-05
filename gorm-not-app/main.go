package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	ID      uint64 `gorm:"primary_key"`
	Title   string `gorm:"size255"`
	Content string
}

const dsn = "host=localhost user=bartukurnaz dbname=notesdb port=5432 sslmode=disable"

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to database")
	}
	fmt.Println("connected to database")

	db.AutoMigrate(&Note{})
	createNote(db, "Learn GORM", "I am doing PostgreSQL CRUD operations with GORM.")
	getAllNotes(db)
	updateNote(db, 1, "GORM Learned", "GORM and PostgreSQL connection was established successfully.")
	getNoteByID(db, 1)
	deleteNote(db, 1)
}

func createNote(db *gorm.DB, title, content string) {
	note := Note{Title: title, Content: content}
	result := db.Create(&note)
	if result.Error != nil {
		fmt.Println("error creating note", result.Error)
		return
	}
	fmt.Println("created note", note)

}

func getAllNotes(db *gorm.DB) {
	var notes []Note
	db.Find(&notes)
	fmt.Println("All notes:")
	for _, note := range notes {
		fmt.Printf("ID: %d | Title: %s | Content: %s\n", note.ID, note.Title, note.Content)
	}
}

func getNoteByID(db *gorm.DB, id uint) {
	var note Note
	result := db.First(&note, id)
	if result.Error != nil {
		fmt.Println("note not found :", result.Error)
		return
	}
	fmt.Printf("ID: %d | Title: %s | Content: %s\n", note.ID, note.Title, note.Content)
}

func updateNote(db *gorm.DB, id uint, newTitle, newContent string) {
	var note Note
	if err := db.First(&note, id).Error; err != nil {
		fmt.Println("note not found.")
		return
	}
	note.Title = newTitle
	note.Content = newContent
	db.Save(&note)
	fmt.Println("note updated")
}
func deleteNote(db *gorm.DB, id uint) {
	var note Note
	if err := db.First(&note, id).Error; err != nil {
		fmt.Println("No notes found to delete.")
	}
}
