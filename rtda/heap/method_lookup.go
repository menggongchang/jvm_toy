package heap

func LookUpMethodInclass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.Name() == name && method.Descriptor() == descriptor {
				return method
			}
		}
	}
	return nil
}

func LookUpMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.Name() == name && method.Descriptor() == descriptor {
				return method
			}
		}
		method := LookUpMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
