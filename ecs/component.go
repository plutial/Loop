package ecs

import (
	"fmt"
	"log"

	"github.com/plutial/Loop/util"
)

// Register a component with the component type
func RegisterComponent[T any](world *World) {
	// Create a slice for the component at the given index
	componentSet := util.NewSparseSet[T]()
	world.ComponentPool[util.GetType[T]()] = &componentSet
}

// Get the address of the component slice
func GetComponentSet[T any](world *World) *util.SparseSet[T] {
	// Get the address of the component slice
	componentSetAddress, ok := world.ComponentPool[util.GetType[T]()].(*util.SparseSet[T])

	if !ok {
		message := fmt.Sprintf("Component type %v not found", util.GetType[T]())
		panic(message)
	}

	return componentSetAddress
}

// Check if an entity has a component
func HasComponent[T any](world *World, id int) bool {
	componentSet := GetComponentSet[T](world)

	_, ok := componentSet.Get(id)

	// Return the check
	return ok
}

// Add a component to an entity
func AddComponent[T any](world *World, id int) *T {
	// Get the component slice
	componentSet := GetComponentSet[T](world)

	// If the entity already has the component, return the address of the component
	if HasComponent[T](world, id) {
		address, _ := componentSet.GetAddress(id)
		return address
	}

	// Add the entity
	var temp T
	componentSet.Add(id, temp)
	address, _ := componentSet.GetAddress(id)

	// Return the address of the component
	return address
}

// Remove a component from an entity
func RemoveComponent[T any](world *World, id int) {
	componentSet := GetComponentSet[T](world)
	componentSet.Delete(id)
}

// Get the address of the component
func GetComponent[T any](world *World, id int) *T {
	componentSet := GetComponentSet[T](world)

	address, ok := componentSet.GetAddress(id)

	// Check if the entity has the component
	if !ok {
		// Send an error message
		message := fmt.Sprintf(
			"Entity %v is either not alive and/or does not have the component %v",
			id, util.GetType[T](),
		)
		log.Fatal(message)
		return nil
	}

	// Return the address of the component
	return address
}

// Returns a slice of entity ids which have component A
func GetEntities[A any](world *World) []int {
	entities := make([]int, 0)

	for id := range world.Size {
		// Check that entity is alive
		if !world.IsEntityAlive(id) {
			continue
		}

		// Add to the slice if the entity has all the required components
		if HasComponent[A](world, id) {
			entities = append(entities, id)
		}
	}

	return entities
}

// Returns a slice of entity ids which have component A and B
func GetEntities2[A, B any](world *World) []int {
	entities := make([]int, 0)

	for id := range world.Size {
		// Check that entity is alive
		if !world.IsEntityAlive(id) {
			continue
		}

		// Add to the slice if the entity has all the required components
		if HasComponent[A](world, id) && HasComponent[B](world, id) {
			entities = append(entities, id)
		}
	}

	return entities
}
