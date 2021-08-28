<h1 align="center">
  GoTify Profile Tracker (Backend)
</h1>

<h4 align="center">A REST API web server built in Go with endpoints for viewing your Friend's Spotify activity!</h4>
<h5 align="center">Uses my Go package <a href="https://github.com/nathanjukes/GoTify-BuddyListt">Spotify-BuddyList</a> for grabbing the data!</h5>


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

- curl http://localhost:8081/user/spotify:user:{SpotifyID}


## Personal use

<h3>Environment Variables needed</h3>

- '<b>Cookie</b>' - sd_pc cookie from Spotify's Web Player
    - After you log into the Spotify web player, you need to find the cookie named 'sp_dc' and get the value of it
    - This cookie expires every year, so if used in production, it will be best to automate the collection of it

- '<b>Client ID</b>' - 
    
- '<b>Client Secret</b>' - 
    
- '<b>Access Token</b>' -
    
- '<b>Refresh Token</b>' - 
    
## Change Log

- Released 28/08/21

### License
[MIT](https://github.com/nathanjukes/GoTify-Profile-Tracker/blob/master/LICENSE.md)
