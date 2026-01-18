# Phoenix 

Phoenix is a platform designed for learning and practicing data structures and algorithms, offering a robust environment for competitive programming. Built to handle untrusted user code, Phoenix features a secure and scalable code execution sandbox. Leveraging Linux namespaces, cgroups, and the Isolate sandbox, the platform can safely compile and run user-submitted solutions against custom test cases concurrently. This sandboxing mechanism protects the host system from potential risks associated with executing arbitrary code. Users can submit their solutions, which are then executed within the sandboxed environment. The platform provides immediate feedback on both the correctness and runtime performance of the submission, similar to popular platforms like Codeforces or LeetCode.

## Getting Started

### Dependencies
* PostgreSQL
* Linux 
* Isolate
* GCC
* Go 

### 1. Setting up isolate 
```
sudo apt-get update
sudo apt-get install libcap-dev asciidoc-base
cd ~/Downloads
git clone https://github.com/phoenix-algo/isolate.git
cd isolate
make
sudo make install
```

### 2. Setting up GCC
```
sudo apt-get update
sudo apt install build-essential
```

### 3. Clone the github repository

### 4. Configure up the enviroment variables 
1. create the .env file in the root of the project
2. set the following environment variables in your .env file
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_NAME=project
DB_PASSWORD=?

JWT_SECRET=?
COOKIE_LIFETIME=7

SERVER_HOST=localhost
SERVER_PORT=8080
```

### 5. Running the server
```
go mod download
go run .
```
## License
This project is licensed under the MIT License - see the license.md file for details
