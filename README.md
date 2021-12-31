# kaomoji-db

This project will distribute and collect well tagged kaomojis (so its easier to find the one you want).

The project consists of 2 base parts: 
  - API
  - WEB-UI
  - (coming soon) QT and gdm desktop GUI

# API Features 
  - [ ] Filtering kaomojis
    - [ ] by tags
    - [ ] by id
    - [ ] by similarity [with a threshold]
  - [ ] Sorting 
    - [ ] by popularity
    - [ ] by length
    - [ ] by tag/mood
    - [ ] by published date [date posted to the server]
  - [ ] Users
    - [ ] Listing
      - [ ] Default lists
        - [ ] Most used
        - [ ] Favourites
      - [ ] Create new lists
      - [ ] Delete lists
      - [ ] Share lists
        - [ ] With other users
        - [ ] By making a hidden list (only acessible if you have the link) 
        - [ ] By making a public list 
      - [ ] Download list (**NO user required**)
        - [ ] As .tsv (Tab Separated Values)
        - [ ] As .tcsv (Tab comma Separated Values): this is the format used bay the [splatmoji](https://github.com/cspeterson/splatmoji) project, but will download a file with extension ***.tsv*** since .tcsv is a convenction of mine
  - [ ] DB => MongoDB

# About this project
This project started being a port of [splatmoji-kaomoji-edit-tool](https://github.com/iacchus/splatmoji-kaomoji-edit-tool) to **GO** since the maintainer of the project droped support and I was looking forward to add many features to make this more, **script friendly** (to ease interaction with other software),**user friendly**, and even give it a **GUI!**

But then I had to make a full web app for school, and since I wasn't fully happy with the [splatmoji](https://github.com/cspeterson/splatmoji) project, since I really strugle reading shell scripts and it din't quite work propertly on the type mode,also the way the data is stored is quite dumb (I got pretty anoyed by this).

I will make this project **compatible** with [splatmoji](https://github.com/cspeterson/splatmoji) anyways, to make it work in the meantime, while the desktop GUI is still in progress.

All that made the perfect excuse to go forward on making an entire 
