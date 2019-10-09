package main

import (
	"testing"
	"reflect"
)

func TestAddContact(t *testing.T)  {
	c := Contact{1, "newContact", "Oybek"}

	contact := addContact(&c)

	if contact != c {
		t.Errorf("some errors have in add contact function")
	}
}

func TestListContact(t *testing.T)  {
	var expectedList []Contact
	expectedList = append(expectedList, Contact{1,"newContact", "Oybek"})

	contactList := listContact()

	if !reflect.DeepEqual(contactList, expectedList) {
		t.Errorf("some errors have in contact list function")
	}
}

func TestUpdateContact(t *testing.T) {
	c := Contact{1, "updateContact", "Oybek"}

	contact, str := updateContact(1, &c)

	if contact != c || str != "" {
		t.Errorf("%v in update contact", str)
	}
}

func TestDeleteContact(t *testing.T)  {
	msg := deleteContact(1)

	if msg != "" {
		t.Errorf("%v in delete contact", msg)
	}
}
