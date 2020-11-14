package ecs

// A System implements logic for processing entities possessing components of
// the same aspects as the system. A System should iterate over its Entities on
// `Update`, in any way suitable for the current implementation.
//
// By convention, systems provide an Add method for adding entities and their
// associated components to the system; e.g.
//
//    Add(basic *ecs.BasicEntity, collision *CollisionComponent, space *SpaceComponent)
type System interface {
	// Update updates the system. It is invoked by the engine once every frame,
	// with dt being the duration since the previous update.
	Update(dt float32)

	// Remove removes the given entity from the system.
	Remove(e BasicEntity)
}

// SystemAddByInterfacer is a system that also implements the AddByInterface method
type SystemAddByInterfacer interface {
	System

	// AddByInterface allows you to automatically add entities based on the
	// interfaces that the entity implements. It should add the entity passed
	// as o to the system after casting it to the correct interface.
	AddByInterface(o Identifier)
}

// Initializer provides initialization of systems.
type Initializer interface {
	// New initializes the given System, and may be used to initialize some
	// values beforehand, like storing a reference to the World.
	New(*World)
}

// systems implements a sortable list of `System`. It is indexed on
// `System.Priority()`.
type systems []System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
