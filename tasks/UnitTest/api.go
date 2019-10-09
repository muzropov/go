package main

var gContactList []Contact

type Contact struct {
	id int
	name string
	created_by string
}

func listContact() []Contact  {
	return gContactList
}

func addContact(c *Contact) Contact {
	gContactList = append(gContactList, *c)

	return *c
}

func updateContact(id int, c *Contact) (Contact, string)  {
	for index, contact := range gContactList {
		if contact.id == id {
			gContactList[index] = *c

			return *c, ""
		}
	}

	return Contact{}, "not found contact"
}

func deleteContact(id int) string {
	for index, contact := range gContactList {
		if contact.id == id {
			gContactList = append(gContactList[:index], gContactList[index+1:]...)

			return ""
		}
	}

	return "not found contact"
}