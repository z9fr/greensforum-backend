docker run --name greensforumdb -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

  
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_TABLE=postgres
export JWT_SECRET=secret
export REFRESH_TOKEN_SECRET=secret
