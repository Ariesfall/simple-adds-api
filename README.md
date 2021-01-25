# simple-adds-api

```sh
go run main.go

go test ./...
```

## 1 Introduction 
This task involves using https://the-odds-api.com/. You’ll need to create a free account for API access. Be sure to check the usage terms. There is an API request limit to consider (we won’t need you to provide your API key on sub mission). We need you to use either MongoDB or PostgreSQL and our preferred language is Go. It would be helpful for you to read the API documentation (available at https://the-odds-api.com/liveapi/guides/v3/) in full before starting the task. 

## 2 Spec 
The API documentation can be found here: https://the-odds-api.com/liveapi/guides/v3/ 
• The App should establish a database connection, the connection string should be passed via command line or environment variable 
• The API key should be passed to the app via command line or environment variable 
• On start-up: 
The app should store all available sports in the database 
The app should store all upcoming fixtures in the database 
• All matches returned from the ’sport=upcoming’ request will be consid ered in-play matches 
• Odds for matches which are not in-play should be updated every hour 
• Due to the API limits, odds for in-play matches should be requested inside a function which accepts parameter used to delay the next API request; we will only ask for 1X2 (h2h) odds from U.K bookmakers for the same reason 
• How the data is stored, both in memory and the database, is entirely up to you. 
• We appreciate that this could be a time-consuming task and so we are not expecting the perfect data storage method. 

### 2.1 Bonus 
• The App should keep track of in-play matches in memory: 
Odds for all in-play matches should be updated in real-time 
Current odds should be readily available e.g Data[Match].Odds[Bookmaker][Market] Do not keep previous prices in memory 
Odds changes should be sent to the database 
