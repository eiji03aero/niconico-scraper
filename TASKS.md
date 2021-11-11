# Tasks

## Backlog
- [ ] consider architecture of implementation and create visual diagram
  - [x] er diagram
    - description, interface, member fields
    - draft with text is done, just proceed with creating diagram
  - [ ] sequence diagram
    - flow of executing audio scraper
    - try mermaid
      - https://mermaid-js.github.io/mermaid/#/
- [ ] research if extracting audio for personal use conflicts with dwango's terms of use

- [ ] Add initialization script and gitignore
  - should create config file for audioscraper
    - should have template in somewhere, and just copy it
  - this config file should be ignored by git
    - so that users can configure them without worrying about git

- [ ] experiment performance with Worker
  - probably should not experiment it with niconico douga
    - could lead to serious trouble
  - So experiment with dummy data instead

## Done
- [x] create visual diagram to describe program architecture
  - overview of program, relationships among them
- [x] lay foundation for audio scraper
  - [x] fill out specification
  - [x] fill out requirements
  - [x] fill out draft implementation
  - [x] create draft schedule
