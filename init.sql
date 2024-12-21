-- Create a 'root' user
CREATE DATABASE mydb owner root
CREATE ROLE root WITH LOGIN PASSWORD 'root';

-- Grant superuser privileges to 'root'
ALTER ROLE root WITH SUPERUSER;

-- Create additional database if needed
CREATE DATABASE additionaldb OWNER root;

-- Grant privileges to 'user'
GRANT ALL PRIVILEGES ON DATABASE mydb TO user;
