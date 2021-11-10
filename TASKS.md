# Tasks

## Backlog
- [ ] consider architecture of implementation and create visual diagram
  - [ ] er diagram
    - description, interface, member fields
    - consider the input
      - shouldn't the input from configuration file be key-value object instead of mere url?
        - for extensibility
  - [ ] sequence diagram
    - flow of executing audio scraper
- [ ] research if extracting audio for personal use conflicts with dwango's terms of use
- [ ] Add initialization script and gitignore
  - should create config file for audioscraper
    - should have template in somewhere, and just copy it
  - this config file should be ignored by git
    - so that users can configure them without worrying about git

## Done
- [x] create visual diagram to describe program architecture
  - overview of program, relationships among them
- [x] lay foundation for audio scraper
  - [x] fill out specification
  - [x] fill out requirements
  - [x] fill out draft implementation
  - [x] create draft schedule
