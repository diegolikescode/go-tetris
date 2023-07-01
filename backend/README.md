# THE BACKEND

## REQUIREMENTS

- [x] An endpoint that allows the frontend to load the CSV file.
- [x] An endpoint that allows the frontend to search through the loaded CSV data.
- [x] The search endpoint should accept query parameters for search terms and filters, and should return the matching results.
- [ ] Appropriate error handling for invalid search queries or other errors.

## INSTRUCTIONS

- [x] The backend should be implemented as a RESTful API using Node. (Try not to use an opinionated framework such as Adonis or Nest).
- [x] [POST /api/files] An endpoint that accepts a CSV file upload from the frontend and stores the data in a database or a data structure.
- [x] [GET /api/users] Should include an endpoint that allows the frontend to search through the loaded CSV data.
- [x] The search endpoint should accept a ?q= query parameter for search terms and should search through EVERY column of the CSV
- [ ] The backend should include appropriate error handling for invalid requests or other errors.
