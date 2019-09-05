# tdict
A simple offline dictionary for the terminal. 
Just give it the word you want to look up as an argument. For example:
```
$ tdict circular
Noun | an advertisement (usually printed on a page or in a leaflet) intended for wide distribution
Adjective | describing a circle; moving in a circle
SYNONYMS: Broadsheet, Rotary, Handbill, Bill, Flyer
ANTONYMS: square
$ tdict wanton
Verb | behave extremely cruelly and brutally
Noun | lewd or lascivious woman
Verb | indulge in a carefree or voluptuous way of life
Verb | spend wastefully
Verb | engage in amorous play
SYNONYMS: Loose, Motiveless, Easy, Trifle, Trifle away
```

The program expects the definition files from [here](https://github.com/tusharlock10/Dictionary/blob/master/data.7z) to be in `/usr/share/tdict/`

The directory should look like this:
```
$ ls /usr/share/tdict/
DA.json  DD.json  DG.json  DJ.json  DM.json  DP.json  DS.json  DV.json  DY.json
DB.json  DE.json  DH.json  DK.json  DN.json  DQ.json  DT.json  DW.json  DZ.json
DC.json  DF.json  DI.json  DL.json  DO.json  DR.json  DU.json  DX.json
```
