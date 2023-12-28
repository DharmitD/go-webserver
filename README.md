# Go Web Application with User Authentication

This is a simple Go web application that demonstrates user authentication, form submissions, and a user dashboard. Users can log in, submit messages, and view their submissions in a dashboard.

## Setup and Usage

### Prerequisites

- [Go](https://golang.org/dl/): Ensure that you have Go installed on your system.

### 1. Clone the Repository

```bash
git clone https://github.com/DharmitD/go-webserver.git
cd go-webserver
```

### 2. Create the Directory Structure

Create the necessary directory structure for your project:

```markdown
go-webserver/
  ├── main.go
  ├── templates/
  │   ├── home.html
  │   ├── form.html
  │   ├── dashboard.html
  │   ├── login.html
  ├── static/
  │   ├── style.css
```

### 3. Customize HTML and CSS

Customize the HTML templates (`home.html`, `form.html`, `dashboard.html`, `login.html`) and the CSS file (`style.css`) to match your project's design and functionality.

### 4. Run the Go Server

Start the Go web server:

```bash
go run main.go
```

### 5. Access the Application

You should see a message indicating that the server is listening on port 8080. Open a web browser and navigate to http://localhost:8080 to access the home page of the web application.

## Features

- **User Authentication**: Users can log in using predefined credentials.
- **Form Submission**: Users can submit messages through a form.
- **User Dashboard**: Authenticated users can view their submissions in a dashboard.

## Further Enhancements/Additions

Here are some further enhancements or additions that can be made to this project:

- **User Registration:** Implement user registration functionality to allow users to sign up for new accounts.
- **Password Hashing:** Improve security by hashing and securely storing user passwords.
- **Database Integration:** Replace in-memory storage with a database (e.g., PostgreSQL, MySQL) for persistent data storage.
- **User Profiles:** Allow users to update their profiles, including their names, email addresses, and profile pictures.
- **Pagination:** Add pagination to the dashboard if there are a large number of submissions.
- **Notifications:** Implement a notification system to alert users of new submissions or updates.
- **Error Handling:** Enhance error handling and provide clear error messages to users.
- **Security:** Implement security measures like CSRF protection and session management best practices.
- **Testing:** Write unit tests and integration tests to ensure the reliability of the application.
- **Deployment:** Deploy the application to a production server using a web server or cloud platform.



