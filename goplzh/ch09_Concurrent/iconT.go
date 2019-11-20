package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

func loadIcon(name string) image.Image {
	// 省略 . . .
	return nil
}

func loadIcons()  {
	icons = map[string]image.Image{
		"spades.png": loadIcon("spades.png"),
	}
}

func Icon(name string) image.Image {
	if icons == nil{
		loadIcon(name)
	}
	return icons[name]
}

var mu1 sync.Mutex

func Icon1(name string) image.Image {
	mu1.Lock()
	defer mu1.Unlock()
	if icons[name] == nil{
		loadIcon(name)
	}
	return icons[name]
}

var mu2 sync.RWMutex

func Icon2(name string) image.Image {
	mu2.RLock()
	if icons != nil {
		mu2.RUnlock()
		return icons[name]
	}
	mu2.RUnlock()

	// acquire an exclusive lock
	mu2.Lock()
	if icons == nil{
		loadIcon("all")
	}
	mu2.Unlock()
	return icons[name]
}

var loadIconsOnce sync.Once

func loadAllA()  {

}
func Icon3(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
