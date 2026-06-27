Presentation Layer

Controller

Service

Repository

Database


Responsibilities:

1) Controller

Receive HTTP request
Parse JSON
Call service
Return response

No business logic.

2) Service

Contains business rules.

Example:

Create Blog

Check duplicate title


Validate business rules

Call Repository


3) Repository

Only database queries.

Example:

Create

FindByID

Update

Delete

FindAll

No HTTP code.


4) Model

Represents database tables.

5) DTO

Represents API request and response objects.

6) Middleware

Authentication

Authorization

Logging

CORS

Rate Limiting

Recovery


7) Config

Application configuration.


8) Logger

Centralized logging.

9) Response

Common API responses.

10) Validator

Request validation.