## eaglescollectedsourcedatservice

##intro

This module starts two servers:

1. file server to give hhtp access to file in datasource folder
2. http server to serve out the content of all files

##to run

```
go build dataserverservice.go && ./dataserverservice

```
check results

```
curl -i localhost:9000/results

```

you should see somthing like this


```

..
{
          "gamedate": "Sat 07 May 2016",
          "team": "STOKE CITY",
          "awayorhome": "H",
          "competition": "League",
          "result": "15:00",
          "score1": "",
          "score2": ""
     },
...