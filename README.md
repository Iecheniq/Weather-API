# Weather-API
A simple weather API using Beego

To run the API follow the steps below:    

-Change all import routes "github.com/iecheniq/weather/(package_name)" to "weather/(package_name)".      
-Change "localhost" to "weather" in all database routes.      
-Change "force" and "verbose" variables to "true" in main.go and in test files to create the    
 database tables then change both variables to "false".    
-cd Path/To/API/folder.        
-docker-compose up  
