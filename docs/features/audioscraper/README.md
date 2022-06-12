# audioscraper
- An automation to aggregate audio files from various sources.
- [Documentation](document.md)

# Schedule
- [x] Lay foundation for this project
- [x] Create supplemental visual diagrams
- [x] Break down implementation and append them to schedule
- [x] Prepare development environment
  - docker, golang
- [ ] Implementation
- [ ] Use the tool to get the beloved vocalo songs!

---

# Development
## Policy
- Write test code with 100% coverage
  - use mock lib if required

## Approaches
- Probably we could consider approach on how to develop.

### Candidates
#### Bottom up
- Develop from the ones that are less dependent.
- Combine the implemented components to build more abstracted components, as a result build the application.
- pros:
  - Easier to start working.
  - Easier to start writing tsets.
- cons:
  - Kind of hard to see the whole picture in the beginning.
    - Could lead to realizing a flaw of architecture in late phase.

#### Top down
- Develop from the top level like going down the steps.
- Leave the unimplemented part placeholder-ed.
- pros:
  - Easier to grasp the entire application.
  - Probably it is suitable for more agile, as we get to understand more about application from the beginning.
- cons:
  - Needs more planning on how to proceed to avoid confusion.

### Conclusion
- Let's go with top-down approach this time:
  - It should be more scaleble (team members).
  - It should provides more chances to brush up architecture.

## Plan
### Phase 1: Lay foundation and create skeleton of the application
- [x] Prepare developmenet environment
  - docker
    - make sure to pick docker image for golang which supports generics
  - docker-compose
- [x] create skeleton
  - follow the folder structure

### Phase 2: Implement lib/work package
- [x] Worker and WorkStation

### Phase 3: Implement binding/nicovideo package
- follow the document

### Phase 4: Implement rest of application agnostic package
- MediaDownloader
- AudioExtractor

### Phase 5: Complete application
- Configuration
- Application
- CLI
- Implementation of work
- Implement init script
