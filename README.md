Employee Management System

To launch the docker services, navigate to the docker directory and run "docker compose up -d".
Assuming you have docker installed and running, this will launch the Azure SQL database and the Mailhog SMTP server.

To run the server, navigate back to the EMS folder and type "go run server.go".
This will start the server and allolw you to hit the /login and /employee endpoints through postman or other means.

Issues/Unfinished parts:
The services running via docker don't actually do anything, I've been unable to connect to the database and haven't 
tried to create the email notification

When querying the API for employee information, no results will be found as the database isn't connected.

