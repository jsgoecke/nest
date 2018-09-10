package nest

/**
 * Helper function to find a Thermostat in the Devices struct by its name
 * Returns nil if no Thermostat was found
 */
func (d *Devices) FindThermostat(name string) *Thermostat {
	for _, t := range d.Thermostats {
		if t.Name == name {
			return t
		}
	}

	return nil
}
