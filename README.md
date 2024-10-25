# Project SapuSapu

A full stack project consisting of Golang as backend and probably React or HTMX *(still haven't decided yet brrr)* . It will scrape data from either existing public website or scrap data from open APIs  and take the URLs, store them in SQLite database. I will then either use those data and pass them through REST API as JSON for FE to consume or I will use those to render on the server and send the HTML to the FE. 

*** Disclosure: This project will not be open for public use once completed since I'm only building this as a way for me to build something with Go. ***

## Progress

Currently I am still going back and forth wheter I should scrape from gogo or nyaa.si. Both have their pros and cons. If I scrape from gogo, there will be a risk of them making changes to their UI thus making my scraper unusable and I have to update it. If I use nyaa.si, there will be some difficulties in searching and parsing the results as the naming there is quite weird.... I have developed 1 API which is getAnime. A POST request can be made and it will return the first page of the search then return them in JSON. I have also set up some tables to be use later on.

## TODO
For getAnime, the definition of complete is, able to store the scraped data into DB. If the animes already exist, skip the scrape part and get the data from db instead.
