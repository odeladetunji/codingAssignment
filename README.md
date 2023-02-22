# codingAssignment

### Clone the Project and move to the main branch.

The root directory contain a docker compose file  (docker-compose-yml)
Make sure to have docker installed in your machine.
The docker compose file contains a PostreSQL docker image.
Run the following docker command on the terminal from the root directory of the project.
   
   >  docker compose up
    
This will initialize the PostreSQL DB.

move to the src/account-service directory and run the command.
  
  >  go run .

This should start the account service;

on another terminal move to the src/transaction-service directory and run the command.
  >  go run .

This should start the Transaction Service.

Upon starting the services three customers will be initialized in the DB so as to have some customers in the system as stated by assignment

The code was written in a way to role back transactions upon failure during account creation.

For Api Testing follow the Url Below (Postman Collection)
https://documenter.getpostman.com/view/2630186/2s93CLrtX8

