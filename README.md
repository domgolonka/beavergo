<p align="center"> 
  <img src="assets/logo.png" width="300" height="300" alt="Foretoken" /></p>
  <h2 align="center">BeaverGo</h2>
  <p align="center">A go client for Beaver https://github.com/Clivern/Beaver/</p>

<p align="center">
    <a href="https://github.com/domgolonka/beavergo/issues/new/choose">Report Bug</a>
    ·
    <a href="https://github.com/domgolonka/beavergo/issues/new/choose">Request Feature</a>
</p>


## About

BeaverGo is a go client for Beaver https://github.com/Clivern/Beaver/

# Usage

### Installation 

It easy to use, all you need to do is import it in your project
    
    import "github.com/domgolonka/beavergo"
    

### Some of the commands are: 

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
