

Part 1
created one GoLang program and one React script for frontend
Did not use any generation tools
Have not programmed in GoLang or React previously
Used Postman, chrome Webserver to server the react page and a chrome extension to get around CORS issue of running both React and GoLang Server on same PC.

tested the endpoints using Postman:
	DELETE 	localhost:8080/app/people/2
	PUT		localhost:8080/app/people/1
	POST	localhost:8080/app/people
			{
			"Id": "6", 
			"Name": "Dan",
			"Age": "93",
			"Balance": "234",
			"Email": "dude@onethirdavenue.com",
			"Address": "2 Langtry Lodge  Parkgate"
			}
	GET 	localhost:8080/app/people
	GET 	localhost:8080/app/people?sort=Name
	GET 	localhost:8080/app/people?sort=Email
	
Added additional functionality to the frontend to allow sort, create and delete

Part 2
Created an alternative GoLang program showing how structures can be extended to a more relational model. Tested this with Postman

I have extended both the frontend and restAPI so handle extra functionality. 
	
	
