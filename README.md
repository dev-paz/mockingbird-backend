# Backend for Mockingbird

Mockingbird lets you record separate videos of you and your frinds singining. 

The app will then stitch them together to make a single music video with everyone singing together in perfect harmony. 



# How it works

The backend is hosted on heroku, written in Go with a progres database. 

Firebase auth is used for login and accounts. 

I use this could video editing software whcih I host on my own aws server for all of the video editing and rendering https://www.openshot.org/cloud-api/

I also setup this VOD solution using AWS to convert, compress and stream the rendered videos quickly https://www.youtube.com/watch?v=d-XCfp97pX0 



