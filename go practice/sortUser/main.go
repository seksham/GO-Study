
package main

import (
    "fmt"
    "sort"
)

type User struct {
    name string
    age  int
}

func main() {
    users := []User{
        {"Eve", 28},
        {"Charlie", 35},
        {"Bob", 25},
        {"Grace", 40},
        {"David", 20},
        {"Jack", 31},
        {"Alice", 30},
        {"Frank", 22},
        {"Ivy", 27},
        {"Henry", 33},
        {"Alice", 25},  // Adding another Alice to demonstrate sorting by age for same name
    }

    // Sort by Name, then by Age
    sort.Slice(users, func(i, j int) bool {
        if users[i].name == users[j].name {
            return users[i].age < users[j].age
        }
        return users[i].name < users[j].name
    })

    fmt.Println("Sorted by Name, then by Age:")
    for _, user := range users {
        fmt.Printf("%s: %d\n", user.name, user.age)
    }
}