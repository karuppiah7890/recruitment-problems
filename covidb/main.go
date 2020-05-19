package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/karuppiah7890/recruitment-problems/covidb/pkg/models"
)

func main() {
	fmt.Println("Welcome to COVIDB! Choose an option from the menu")
	var choice int
	exit := false
	people := models.People{}

	for !exit {
		fmt.Println("\n\nMenu")
		fmt.Println("1. Add information about positively tested person")
		fmt.Println("2. List information about positively tested person")
		fmt.Println("3. Exit")
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Println("Enter information about the person")
			person := getInformationAboutPerson(people)
			people[person.ID] = person
			fmt.Println("Successfully stored the information about the person")
		case 2:
			fmt.Println("Listing the information about all the people")
			listOutInformationAboutPeople(people)
		case 3:
			exit = true
		default:
			fmt.Printf("Unknown option %d\n", choice)
		}
	}
}

func listOutInformationAboutPeople(people models.People) {
	for _, person := range people {
		fmt.Println(person)
	}
}

func getInformationAboutPerson(people models.People) models.Person {
	gender := getGender()
	age := getAge()
	area := getArea()
	city := getCity()
	relationships := getRelationships(people)
	id := uuid.New()

	person := models.Person{
		ID:            id,
		Gender:        gender,
		Age:           age,
		Area:          area,
		City:          city,
		Relationships: relationships,
	}

	return person
}

func getGender() string {
	fmt.Print("Gender: ")
	gender := readLine()
	return gender
}

func getAge() int {
	var age int
	fmt.Print("Age: ")
	fmt.Scanf("%d", &age)
	return age
}

func getArea() string {
	fmt.Print("Area: ")
	area := readLine()
	return area
}

func getCity() string {
	fmt.Print("City: ")
	city := readLine()
	return city
}

func getRelationships(people models.People) models.Relationships {
	var choice string
	fmt.Print("Does the user have any relationships with existing patients? [Y/y/Yes/yes/N/n/No/no] ")
	fmt.Scanf("%s", &choice)

	if choice == "N" || choice == "n" || choice == "No" || choice == "no" {
		return nil
	}

	var relationships models.Relationships

	for true {
		fmt.Println("Type information about the relationship")
		relationship, err := getRelationship()
		if err != nil {
			if _, ok := err.(InvalidUUID); ok {
				fmt.Printf("%s. Continuing...\n", err.Error())
				continue
			}
			panic(err)
		}
		if _, exists := people[relationship.PersonID]; !exists {
			fmt.Printf("No person with the ID %s exists. Continuing...\n", relationship.PersonID)
			continue
		}
		relationships = append(relationships, relationship)

		fmt.Print("Does the user have one more relationship? [Y/y/Yes/yes/N/n/No/no] ")
		fmt.Scanf("%s", &choice)

		if choice == "N" || choice == "n" || choice == "No" || choice == "no" {
			break
		}
	}

	return relationships
}

// InvalidUUID denotes an error for invalid UUID
type InvalidUUID struct {
	InvalidInput string
	ParseError   string
}

func (invalidUUID InvalidUUID) Error() string {
	return fmt.Sprintf("The ID %s has an invalid format", invalidUUID.InvalidInput)
}

func getRelationship() (models.Relationship, error) {
	fmt.Print("Name of the relationship: ")
	relationshipName := readLine()

	fmt.Print("ID of the person: ")
	personID := readLine()

	personUUID, err := uuid.Parse(personID)
	if err != nil {
		return models.Relationship{}, InvalidUUID{
			InvalidInput: personID,
			ParseError:   err.Error(),
		}
	}

	relationship := models.Relationship{
		Name:     relationshipName,
		PersonID: personUUID,
	}

	return relationship, nil
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}
