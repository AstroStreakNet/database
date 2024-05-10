# AstroStreakNet Database Management System

The AstroStreakNet Database Management System is a backend system designed to manage various aspects of an astronomy-related platform, including user management, image storage, analytics, and authentication. It provides RESTful APIs for interacting with the database, allowing clients to perform CRUD operations on different resources.

## Features

- **User Management:** Allows the creation, retrieval, updating, and deletion of user accounts. Users can be associated with various attributes such as first name, last name, date of birth, gender, and permissions.
- **Image Management:** Provides functionality to manage astronomical images, including storing image metadata such as image path, observatory code, time of observation, exposure duration, coordinates, streak type, and user ID.
- **Analytics:** Enables the collection and retrieval of analytics data related to user interactions with the platform, including page views, image views, clicks, session data, referral sources, navigation paths, abandonment rate, repeat visits, and time between visits.
- **Authentication:** Facilitates user authentication by managing session tokens, login times, and logout times.

## Installation

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/AstroStreakNet/database.git
    ```

2. Install dependencies:

    ```bash
    cd database
    go mod tidy
    ```

3. Configure the database connection by updating the `config/config.go` file with your database credentials.

4. Run the application:

    ```bash
    go run main.go
    ```

## API Endpoints

### Users

- **GET /user/:userID**: Get user details by user ID.
- **POST /user/**: Create a new user.
- **PUT /user/:userID**: Update user details.
- **DELETE /user/:userID**: Delete a user by user ID.

### Images

- **GET /image/:imageID**: Get image details by image ID.
- **POST /image/**: Create a new image.
- **PUT /image/:imageID**: Update image details.
- **DELETE /image/:imageID**: Delete an image by image ID.

### Analytics

- **GET /analytics/:analyticsID**: Get analytics data by analytics ID.
- **POST /analytics/**: Create new analytics data.
- **PUT /analytics/:analyticsID**: Update analytics data.
- **DELETE /analytics/:analyticsID**: Delete analytics data by analytics ID.

### Authentication

- **POST /authentication/**: Create a new authentication session.
- **GET /authentication/:authID**: Get authentication details by authentication ID.
- **DELETE /authentication/:authID**: Delete an authentication session by authentication ID.
