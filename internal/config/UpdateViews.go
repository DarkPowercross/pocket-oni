package config

func (c *Config) UpdateTUIinformation() {
	header := c.App.Approot.View.Header
	content := c.App.Approot.View.Content.Sub.SpriteInformation.SubInformationViews
	information := c.SpriteInformation

	header.HeaderThought.SetText(information.GetThought())
	header.HeaderWeather.SetText(information.GetWeather())
	header.HeaderFeedback.SetText(information.Message.GetMessage())
	header.HeaderCharacter.SetText(information.GetState())

	content.SpriteHunger.SetText(information.GetFood())
	content.SpriteHappiness.SetText(information.GetHappiness())
	content.SpriteWater.SetText(information.GetWater())
	content.SpriteLocation.SetText(information.GetLocation())
	content.SpriteHealth.SetText(information.GetHealth())
	content.SpriteComfort.SetText(information.GetComfort())
}

