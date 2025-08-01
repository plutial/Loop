package ecs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	// Game packages
	"github.com/plutial/Loop/gfx"
	"github.com/plutial/Loop/physics"
)

type GameMapData struct {
	// Tile size
	TileWidth  int `json:"tilewidth"`
	TileHeight int `json:"tileheight"`

	// Tile Layer Size
	LayerWidth  int `json:"width"`
	LayerHeight int `json:"height"`

	// Tile layers
	TileLayers []TileLayerData `json:"layers"`
}

type TileLayerData struct {
	Data []int  `json:"data"`
	Name string `json:"name"`
}

type TileTag bool

func (world *World) LoadMap(path string) {
	// Open the json file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Map
	var gameMapData GameMapData

	json.Unmarshal(data, &gameMapData)

	// Load the tile texture
	tileTexture := gfx.NewTexture("assets/res/GrassTiles.png")

	// Put the tiles into the world
	for y := range gameMapData.LayerHeight {
		for x := range gameMapData.LayerWidth {
			tileSourceId := gameMapData.TileLayers[0].Data[y*gameMapData.LayerWidth+x]

			// If the tile does not exist, then do not load it
			if tileSourceId == 0 {
				continue
			}

			// Create an entity
			id := world.NewEntity()

			// Add the tile tag
			AddComponent[TileTag](world, id)

			// Create the sprite component
			sprite := AddComponent[gfx.Sprite](world, id)

			*sprite = gfx.NewSprite(tileTexture)

			// Source rectangle
			textureSize := physics.NewVector2(
				float64(tileTexture.Bounds().Dx()),
				float64(tileTexture.Bounds().Dy()),
			)

			sprite.Source.Position.X = float64((tileSourceId-1)%int(int(textureSize.X)/16)) * 16
			sprite.Source.Position.Y = float64((tileSourceId-1)/int(int(textureSize.Y)/16)) * 16

			sprite.Source.Size = physics.NewVector2(16, 16)

			// Destination rectangle
			sprite.Destination.Position = physics.NewVector2(float64(x)*16, float64(y)*16)
			sprite.Destination.Size = physics.NewVector2(16, 16)

			// Create the physics body
			body := AddComponent[physics.Body](world, id)

			position := physics.NewVector2(float64(x)*16, float64(y)*16)
			size := physics.NewVector2(16, 16)

			*body = physics.NewBody(position, size)
		}
	}
}
