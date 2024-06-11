FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=fakepassword
ENV POSTGRES_DB=frdatabase

# Copy initialization scripts into Docker image
COPY initialize_database.sql /docker-entrypoint-initdb.d/
COPY test_data.sql /docker-entrypoint-initdb.d/
