package lenconv

func KmToMi(km Kilometer) Mile {
	return Mile(km * .62)
}

func MiToKm(mi Mile) Kilometer {
	return Kilometer(mi * 1.61)
}
