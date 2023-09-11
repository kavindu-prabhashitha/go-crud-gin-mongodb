#### Go language folder structure for web app/sever ####

1. You can place the readme text here, like setup steps of the project , or any other information about the project that you think would be helpful for other developers

2. Create ".env" file, where we will keep all the environment related constants, like database username, password , secret key of jwt if any, etc

3. Create the main.go file in the root, this will serve as the project/webserver entry point

4. create a "Makefile" this will contain some basic commands to run the project

5. create these folders:
    1. Controllers : contains controller that accepts a request and call a particular service to process it
    2. Services : contains services that has the logic for doing all the process like manipulating DB and giving back desired results
    3. Models : contains the structs for storing data in DB or fetching data from DB
    4. DB : keep your DB related operation, these can be used in services to fetch / update / insert
        *** Place DB connection filesas well inside this folder ****
    5. Routes : keep all your routes here and pass the name of controller
    6. Config : keep all your configuration things here , fetching variables from .env file or even db config can be placed here
    7. Constants: keep all constants here, can categorise them into separate files
    8. Utils : keep commonly used functions here like capitalizing first letter, etc