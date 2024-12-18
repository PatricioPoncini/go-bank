# Go Bank 🏛️💰

A simple and secure banking application built with Go.

---

## 🚀 Getting Started

Follow these steps to clone, install, and run the project using Docker.

---

### 🛠 Prerequisites

- [Docker](https://www.docker.com/) installed on your machine.
- [Git](https://git-scm.com/) installed.

---

### 📥 Cloning the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/PatricioPoncini/go-bank.git
cd go-bank
```

### 🐳 Setting Up the Docker Container
Start a PostgreSQL container:
```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

### 📄 Creating the .env File
Create a .env file in the project root with the following environment variables:
```bash
JWT_SECRET=my_super_secret
```

### 🔧 Running the Application
To start the application, simply execute the following command:
```bash
make run
```