package serializer

type TypeRegistry interface {
	Register(tis ...Type)
	Find(id TypeID) (ctor Constructor, exists bool)
}

func NewTypeRegistry() TypeRegistry {
	return typeRegistry{
		make(map[TypeID]Constructor),
	}
}

type typeRegistry struct {
	memory map[TypeID]Constructor
}

func (tr typeRegistry) Register(ts ...Type) {
	for _, t := range ts {
		tr.memory[t.ID] = t.Constructor
	}
}

func (tr typeRegistry) Find(id TypeID) (ctor Constructor, exists bool) {
	ctor, exists = tr.memory[id]
	return ctor, exists
}
