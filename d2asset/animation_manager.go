package d2asset

type animationManager struct {
	cache *cache
}

func createAnimationManager() *animationManager {
	return &animationManager{cache: createCache(AnimationBudget)}
}

func (sm *animationManager) loadAnimation(animationPath, palettePath string) (*Animation, error) {
	cachePath := animationPath + palettePath
	if animation, found := sm.cache.retrieve(cachePath); found {
		return animation.(*Animation).clone(), nil
	}

	dc6, err := loadDC6(animationPath, palettePath)
	if err != nil {
		return nil, err
	}

	animation, err := createAnimationFromDC6(dc6)
	if err != nil {
		return nil, err
	}

	sm.cache.insert(cachePath, animation.clone(), 1)
	return animation, err
}
