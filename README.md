<h1 align="center">
  GoTify Profile Tracker (Backend)
</h1>

<h4 align="center">A REST API web server built in Go with endpoints for viewing your Friend's Spotify activity!</h4>
<h5 align="center">Uses my Go package <a href="https://github.com/nathanjukes/GoTify-BuddyList">Spotify-BuddyList</a> for grabbing the data!</h5>


<p align="center">
  <a href="https://github.com/nathanjukes/GoTify-Profile-Tracker">
      <img src="https://scrutinizer-ci.com/g/pH7Software/pH7-Social-Dating-CMS/badges/build.png?b=master">
  </a>
  <a href="https://github.com/nathanjukes/GoTify-Profile-Tracker">
    <img src="https://img.shields.io/badge/version-v1.0-blue">
  </a>
  <a href="https://github.com/nathanjukes/GoTify-Profile-Tracker/blob/master/LICENSE.md">
    <img src="https://img.shields.io/github/license/Naereen/StrapDown.js.svg">
  </a>
</p>

## What is it & why?

GoTify is a collection of Spotify utilities written in Go by myself. GoTify Profile Tracker is a REST API web server that provides an implementation of my <a href="https://github.com/nathanjukes/GoTify-BuddyList">GoTify BuddyList package</a> and scales up past the Spotify endpoints to retain large sets of data. The main use of this API is to view the listening activity of your friends in a well-formatted manner, there is also a <a href="https://github.com/nathanjukes/GoTify-Profile-Tracker-Frontend">frontend Vue.js application</a> attached to this that utilises the API. 

## Example Queries

- curl http://localhost:8081/activity

### Example Response

Array containing this object:

```json
{
   "timestamp":1630585689720,
   "time":"2021-09-02T13:28:09.72+01:00",
   "user_id":"spotify:user:21dnoih5aa7vkozsvczwoe4za",
   "user":{
      "uri":"spotify:user:21dnoih5aa7vkozsvczwoe4za",
      "name":"Lars.On.Spotify",
      "imageUrl":"https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=541218032718353&height=50&width=50&ext=1632658424&hash=AeQSx5CUG26hrnhmIXA"
   },
   "track_id":"spotify:track:7yq4Qj7cqayVTp3FF9CWbm",
   "track":{
      "uri":"spotify:track:7yq4Qj7cqayVTp3FF9CWbm",
      "name":"Riptide",
      "imageUrl":"http://i.scdn.co/image/ab67616d0000b273d3ce97395ff522b0d70c1094",
      "album_id":"spotify:album:6rIbiUMmZJfqJRnXhVxFvg",
      "album":{
         "uri":"spotify:album:6rIbiUMmZJfqJRnXhVxFvg",
         "name":"Dream Your Life Away"
      },
      "artist_id":"spotify:artist:10exVja0key0uqUkk6LJRT",
      "artist":{
         "uri":"spotify:artist:10exVja0key0uqUkk6LJRT",
         "name":"Vance Joy"
      }
   }
}
```

- curl http://localhost:8081/user/spotify:user:{SpotifyID}

### Example Response

Array containing this object:

```json
{
   "timestamp":1630585689720,
   "time":"2021-09-02T13:28:09.72+01:00",
   "user_id":"spotify:user:21dnoih5aa7vkozsvczwoe4za",
   "user":{
      "uri":"spotify:user:21dnoih5aa7vkozsvczwoe4za",
      "name":"Lars.On.Spotify",
      "imageUrl":"https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=541218032718353&height=50&width=50&ext=1632658424&hash=AeQSx5CUG26hrnhmIXA"
   },
   "track_id":"spotify:track:7yq4Qj7cqayVTp3FF9CWbm",
   "track":{
      "uri":"spotify:track:7yq4Qj7cqayVTp3FF9CWbm",
      "name":"Riptide",
      "imageUrl":"http://i.scdn.co/image/ab67616d0000b273d3ce97395ff522b0d70c1094",
      "album_id":"spotify:album:6rIbiUMmZJfqJRnXhVxFvg",
      "album":{
         "uri":"spotify:album:6rIbiUMmZJfqJRnXhVxFvg",
         "name":"Dream Your Life Away"
      },
      "artist_id":"spotify:artist:10exVja0key0uqUkk6LJRT",
      "artist":{
         "uri":"spotify:artist:10exVja0key0uqUkk6LJRT",
         "name":"Vance Joy"
      }
   }
}
```

## Personal use

<h3>Environment Variables needed</h3>
<h4>To set the environment variables, insert them into the init() method of the main.go file</h4>

- '<b>Cookie</b>' - sd_pc cookie from Spotify's Web Player
    - After you log into the Spotify web player, you need to find the cookie named 'sp_dc' and get the value of it
    - This cookie expires every year, so if used in production, it will be best to automate the collection of it

- '<b>Client ID</b>' - Follow the 'Register your App' tutorial <a href="https://developer.spotify.com/documentation/general/guides/app-settings/">here!</a> 
    
- '<b>Client Secret</b>' - Follow the 'Register your App' tutorial <a href="https://developer.spotify.com/documentation/general/guides/app-settings/">here!</a> 
    
- '<b>Access Token</b>' - Follow this: <a href="https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow">Link!</a>. Once you have the Access and Refresh token, insert them into the init() method. Be sure to request the 'user-follow-modify' scope.
    
- '<b>Refresh Token</b>' - Follow this: <a href="https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow">Link!</a>. Once you have the Access and Refresh token, insert them into the init() method. Be sure to request the 'user-follow-modify' scope.
    
## Change Log

- Released 28/08/21

### License
[MIT](https://github.com/nathanjukes/GoTify-Profile-Tracker/blob/master/LICENSE.md)
