ARG PGUSER
ARG PGPASSWORD
ARG PGDBNAME

FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=${PGUSER}
ENV POSTGRES_PASSWORD=${PGPASSWORD}
ENV POSTGRES_DB=${PGDBNAME}

# Copy initialization scripts into Docker image
COPY initialize_database.sql /docker-entrypoint-initdb.d/
COPY test_data.sql /docker-entrypoint-initdb.d/
