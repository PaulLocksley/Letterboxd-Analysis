package main

func writeCacheResults(movies []movie) {
	if movieCacheLock {
		return
	}

	movieCacheLock = true
	for i := range movies {
		if _, ok := movieCache[movies[i].ID]; ok {
			continue
		}
		movieCache[movies[i].ID] = movies[i].Crew
	}

	saveCache()
	movieCacheLock = false
}
