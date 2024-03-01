# ASCII-Art-Web

Description
-----------

This is a Go program that starts a server at port 8080 and handles GET and POST requests for an ASCII art generator web application.


How To use
----------

1. Clone the repo in your PC.
2. Navigate to the main directory.
3. Type `go run .` in your terminal and press ENTER.
4. Navigate to `http://localhost:8080`
5. Type the desired text in the textfield.
6. Choose the font.
7. Press the Submit or Download button. || *You may press Download without submitting the input first.*


How it works
------------

This is a Go program that starts a web server at http://localhost:8080 and handles two routes:
the home page ("/") and a route for generating ASCII art from user input ("/asciiart").

The home page displays an HTML form that allows the user to input text and choose a font for generating ASCII art. The "/asciiart" route handles a POST request, validates the user input, reads the corresponding font file from the "banners" directory, and generates ASCII art using the provided text and font.

The generated ASCII art is then displayed on the home page using a Go HTML template.


#### Authors

Anton Urban & Daniil Leonov
Kood/Johvi 2023


How it works
------------

This is a Go program that starts a web server at http://localhost:8080 and handles two routes:
the home page ("/") and a route for generating ASCII art from user input ("/asciiart").

The home page displays an HTML form that allows the user to input text and choose a font for generating ASCII art. The "/asciiart" route handles a POST request, validates the user input, reads the corresponding font file from the "banners" directory, and generates ASCII art using the provided text and font.

The generated ASCII art is then displayed on the home page using a Go HTML template.


#### Authors

Anton Urban & Daniil Leonov
Kood/Johvi 2023