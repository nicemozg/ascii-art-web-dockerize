Ascii-art-web

Description

Ascii-art-web is a web application written in Go that allows users to create ASCII art using different banners (shadow, standard, thinkertoy). It provides a graphical user interface (GUI) for generating ASCII art from text input.

Authors

Kuanysh Anuarbekov and Nurgisa Serikkaliyev

Usage: how to run

Clone the repository: git clone https://zero.academie.one/git/kanuarbe/ascii-art-web
Navigate to the project directory: cd ascii-art-web
Run the server: go run main.go
Open your web browser and go to http://localhost:8080 to access the application.
Implementation details: algorithm
Ascii-art-web uses Go's standard HTTP package to create a server. It defines two HTTP endpoints:

GET /: Renders the main page with a form for text input and banner selection.
POST /ascii-art: Receives data (text and banner selection) from the client, generates ASCII art using the specified banner, and returns the result.
The application follows the following steps:

Define HTML templates for the main page and 404 error page.
Create HTTP handlers to handle GET and POST requests.
Parse form data from the client-side to extract text and banner selection.
Generate ASCII art using the selected banner.
Return the ASCII art as a response to the client.

Instructions

Ensure you have Go installed on your system.
Clone the repository.
Navigate to the project directory.
Run the server using go run main.go.
Access the application in your web browser at http://localhost:8080.
Enter text and select a banner, then click the submit button to generate ASCII art.

Allowed Packages

Ascii-art-web only uses standard Go packages and does not rely on any third-party libraries.

Feel free to reach out if you have any questions or need further assistance!