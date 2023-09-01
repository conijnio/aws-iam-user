package adapters

var registeredAdapters = map[string]map[string]IUserAdapter{}

func prepareRegionMap(region string) {
	if _, exists := registeredAdapters[region]; !exists {
		registeredAdapters[region] = make(map[string]IUserAdapter)
	}
}

func RegisterAdapter(region string, profile string, adapter IUserAdapter) {
	prepareRegionMap(region)
	registeredAdapters[region][profile] = adapter
}

func HasAdapter(region string, profile string) bool {
	if _, exists := registeredAdapters[region]; !exists {
		return false
	}

	if _, exists := registeredAdapters[region][profile]; !exists {
		return false
	}

	return true
}

func GetAdapter(region string, profile string) IUserAdapter {
	return registeredAdapters[region][profile]
}
