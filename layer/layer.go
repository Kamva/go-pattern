package layerpattern

// Layer is a design pattern to make layered services,...
// e.g validation->sanitizer->guard->core
// all layers implements same interface (e.g Service interface)
type Layer interface {
	SetRoot(root interface{})
	SetNext(next interface{})
}

// SetLayers set layers in order of provided Layers.
// finally returns the root of services.
func SetLayers(services ...Layer) interface{} {
	if len(services) == 0 {
		return nil
	}
	root := services[0]

	var prev Layer
	for _, current := range services {
		current.SetRoot(root)

		if prev != nil {
			prev.SetNext(current)
		}

		prev = current
	}

	return root
}
