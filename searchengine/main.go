package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Name  string
	Email string
}

var DataBase = []User{
	{Name: "William S. Hopkins", Email: "WilliamSHopkins@teleworm.us"},
	{Name: "Mary G. Hensley", Email: "MaryGHensley@armyspy.com"},
	{Name: "Dustin S. Wagner", Email: "DustinSWagner@jourrapide.com"},
	{Name: "Peter A. Mize", Email: "PeterAMize@dayrep.com"},
	{Name: "Peter A. Mize", Email: "PeterAMize@dayrep.com"},
	{Name: "Elia J. Libby", Email: "EliaJLibby@jourrapide.com"},
	{Name: "Judy D. Chiaramonte", Email: "JudyDChiaramonte@teleworm.us"},
	{Name: "Denise S. Piscitelli", Email: "DeniseSPiscitelli@teleworm.us"},
	{Name: "Dixie R. Platz", Email: "DixieRPlatz@teleworm.us"},
	{Name: "Dixie R. Platz", Email: "DixieRPlatz@teleworm.us"},
	{Name: "Samuel T. Kight", Email: "SamuelTKight@teleworm.us"},
	{Name: "Larry M. Larrabee", Email: "LarryMLarrabee@dayrep.com"},
	{Name: "Gregory N. Payne", Email: "GregoryNPayne@rhyta.com"},
	{Name: "Vincent J. White", Email: "VincentJWhite@armyspy.com"},
	{Name: "Eddie Reed", Email: "eddie.reed@example.com"},
	{Name: "Julian Reed", Email: "julian.reed@example.com"},
	{Name: "Marie Matthews", Email: "marie.matthews@example.com"},
	{Name: "Kirk Henderson", Email: "kirk.henderson@example.com"},
	{Name: "Leonard Baker", Email: "leonard.baker@example.com"},
	{Name: "Julian Frazier", Email: "julian.frazier@example.com"},
	{Name: "Troy Davidson", Email: "troy.davidson@example.com"},
	{Name: "Ashley Nichols", Email: "ashley.nichols@example.com"},
	{Name: "Allen Webb", Email: "allen.webb@example.com"},
	{Name: "Warren Gomez", Email: "warren.gomez@example.com"},
	{Name: "Mike Kuhn", Email: "mike.kuhn@example.com"},
	{Name: "Sharlene Wright", Email: "sharlene.wright@example.com"},
	{Name: "Derek Flores", Email: "derek.flores@example.com"},
	{Name: "Bradley Jensen", Email: "bradley.jensen@example.com"},
	{Name: "Leon Dixon", Email: "leon.dixon@example.com"},
	{Name: "Kitty Chapman", Email: "kitty.chapman@example.com"},
	{Name: "Anita Murray", Email: "anita.murray@example.com"},
	{Name: "Daisy Harris", Email: "daisy.harris@example.com"},
	{Name: "Dwayne Knight", Email: "dwayne.knight@example.com"},
	{Name: "Sara Mills", Email: "sara.mills@example.com"},
	{Name: "Genesis Harvey", Email: "sara.mills@example.com"},
}

type Worker struct {
	users []User
	ch    chan *User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if strings.Contains(user.Email, email) {
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]
	ch := make(chan *User)
	log.Printf("Looking for %s", email)
	go NewWorker(DataBase[:7], ch).Find(email)
	go NewWorker(DataBase[7:14], ch).Find(email)
	go NewWorker(DataBase[14:21], ch).Find(email)
	go NewWorker(DataBase[21:28], ch).Find(email)
	go NewWorker(DataBase[28:], ch).Find(email)

	// if channel returns an user I show him
	// But if in 1 second any user has been returned I assume that user was not found
	for {
		select {
		case user := <-ch:
			log.Printf("The email %s is owned by %s", user.Email, user.Name)
		case <-time.After(100 * time.Millisecond):
			return
		}
	}
}
