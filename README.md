It's a Go client for Beaver

https://github.com/Clivern/Beaver/



# Usage

#### Some of the commands are: 

Token and url are in your beaver config. Make sure to include the protocol and port (if any).
Example: `http://localhost:8080`

`chat := NewConnect(token, url) `

`health := chat.HealthCheck()`

`chat.CreateConfig(key, value)`

`chat.GetConfig(key)`

`chat.UpdateConfig(value)`

`chat.DeleteConfig(key)`

`chat.GetChannel(channelname)`

`chat.CreateChannel(channel, type)`

`chat.UpdateChannel(channel, type)`

`chat.PublishChannel(channel, data)`

`chat.BroadcastChannel(channels, data)`

`chat.DeleteChannel(channel) `

`chat.CreateClient(channel []string)`

`chat.GetClient(id)`

`chat.SubscribeClient(channels, id)`

`chat.UnsubscribeClient(channels, id)`

`chat.DeleteClient(id)`
