package config

import (
	"context"
	"time"
)

func (c *Config) AnimateTUI(ctx context.Context) {
	frames, delay := c.Character.GetCurrentImages(c.SpriteInformation.State, c.SpriteInformation.Location)
	if len(*frames) == 0 {
		return
	}

	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond * 5)
	defer ticker.Stop()

	i := 0

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			frames, _ := c.Character.GetCurrentImages(c.SpriteInformation.State, c.SpriteInformation.Location)
			if len(*frames) == 0 {
				return
			}

			frame := (*frames)[i]
			i = (i + 1) % len(*frames)

			c.App.App.QueueUpdateDraw(func() {
				c.App.Approot.View.Content.Sub.Sprite.UpdateSprite(frame)

				c.UpdateTUIinformation()
			})
		}
	}
}
