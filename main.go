package main



func main() {
	app := fiber.New()
	// User paths
	app.Post("/user", createuser)  				// Create new user
	app.Get("/user/:id?", getuser) 				// Get exsisting user by user id
	
	// Post paths
	app.Post("/posts", createuser)  			// Create new post
	app.Get("/posts/:id?", getuser)				// Get existing post by post id
	app.Get("/posts/users/:id?", getuser)		// Get all the posts of a user using user id
	
	app.Listen(port)
  }