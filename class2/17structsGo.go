package main

import "fmt"
func main() {
	var aLanguage struct{
		name string
		compiled bool
		openSource bool
		imperative bool
		oo string
		staticallyTyped bool
		stronglyTyped bool
		version float32
	}
	aLanguage.name = "Go"
	aLanguage.compiled = true
	aLanguage.openSource = true
	aLanguage.imperative = true
	aLanguage.oo = "It is not object-oriented in the normal sense like Java and C++ because it does not have the concept of classes and inheritance."
	aLanguage.staticallyTyped = true
	aLanguage.version = 1.16
	fmt.Println("Name: ",   aLanguage.name)
	fmt.Println("Compiled: ",   aLanguage.compiled)
	fmt.Println("Open Source: ",   aLanguage.openSource)
	fmt.Println("Imperative: ",   aLanguage.imperative)
	fmt.Println("Object Oriented: ",   aLanguage.oo)
	fmt.Println("Statically Typed: ",   aLanguage.staticallyTyped)
	fmt.Println("Version: ",   aLanguage.version)

}
