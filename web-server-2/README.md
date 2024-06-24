## Building a CRUD API with Golang

In this video, we'll create a complete CRUD API using Golang, focusing on a movie server that allows us to list, create, update, and delete movies. Here's an overview:

### Key Concepts

- **No Database**: We'll use structs and slices to manage data within the server itself, rather than relying on a database.
- **Local Server**: The API will be served on `localhost:8000`.

### Tools and Libraries

- **Gorilla Mux**: We'll use the Gorilla Mux library for routing. Instructions on installing it will be provided.

### API Routes

We will implement five different routes:

1. **Get All Movies**
   - **Route**: `/movies`
   - **Method**: GET
   - **Function**: `getMovies()`
2. **Get Movie by ID**
   - **Route**: `/movies/{id}`
   - **Method**: GET
   - **Function**: `getMovie()`
3. **Create Movie**
   - **Route**: `/movies`
   - **Method**: POST
   - **Function**: `createMovie()`
4. **Update Movie**
   - **Route**: `/movies/{id}`
   - **Method**: PUT
   - **Function**: `updateMovie()`
5. **Delete Movie**
   - **Route**: `/movies/{id}`
   - **Method**: DELETE
   - **Function**: `deleteMovie()`

### HTTP Methods

- **GET**: Used for both getting all movies and getting a movie by ID.
- **POST**: Used for creating a new movie.
- **PUT**: Used for updating an existing movie.
- **DELETE**: Used for deleting a movie.

By following this structure, you'll have a clear understanding of the CRUD operations and how to implement them using Golang and Gorilla Mux.
