# splatmoji-kaomoji-db-edior

# File format *Standarts used**
> by Standards i do NOT mean industry standards, this explains standarized formats for db files compatible with the project
- ## [ .tsv ] :
  ### **Tab Separated Values**
  This formar is although very weird is a stardard format non the less.

- ## [ .csv ] :
  ### **Comma Separated Values**
  This format is probably the most used for storing data in a table form, this IS an industry standard.
  

- ## [ .tcsv ] :
  ### **Tab Comma Separated Values**
  > this format is the one used by the splatmoji project [I made up the name tcsv], stored in a .tsv file, .tcsv file extensions will also be acepted by this project. (wou will understand why of this format in a minute:)
  
  This format happens to be a mixed format between a tsv and csv. this format is esentially a tsv format with only 2 column, where the second holds a set of values **encapsulated** in a csv format. 
  example:
  >tabs will be changed by " --> " for the example so they are vissible
  ```csv
  kaomoji --> key,key2,key3
  (ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ --> excited,spark
  ``` 
  this weirdo of a format is used due to the nature of tabs and kaomojis, kaomojis need to be static, so that no mater where you paste them, the look the same, where tabs are inheritly dynamic, and because of that you will not find any tab inside any kaomoji, and so, tsv format is perfect for storing kaomojis along side descriptive data for it.

  **But why encapsulating the keywords in a single tsv column using csv?**
  As far as I know, theres no specific reason, and in my opinion, pure tsv format sould be used.


# Coming soon features:
- GUI
- Merge db files
- push db file to your current db (related to merge db files)
- check db file syntax 
- Import db (checks syntax and replaces current db file)
# About this project
this project is in part a port of [splatmoji-kaomoji-edit-tool](https://github.com/iacchus/splatmoji-kaomoji-edit-tool) to **GO** but as the maintainer of the project in lookin forward to add many features to make this more, **script friendly** (so its easier to automate usage and ease interaction with other software),**user friendly**, and even give it a **GUI!**
