🎩 Relational Database in Go (with PostgreSQL Magic)

Why This Exists?

Decided to crank up my Golang skills – and what better way than by wrangling a relational database?  
PostgreSQL + Go = chef’s kiss.  

🛠️ What You’ll Need

- PostgreSQL – Either install it the old-fashioned way (binary download) or roll with Docker like I did.  
- Golang– Obviously.  

🚀 Spinning Up PostgreSQL (Docker Style)  

Don’t feel like installing Postgres directly? I didn't either lol.  
Run this one-liner to get Postgres up and running in Docker:  
```bash
docker run \
-d \
-e POSTGRES_HOST_AUTH_METHOD=trust \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=dbname \
-p 5432:5432 \
postgres:16-alpine
```
🔹 `POSTGRES_HOST_AUTH_METHOD=trust` – No password hassle for local dev.  
🔹 `POSTGRES_USER=user` – Customize it to whatever fits your vibe.  
🔹 `POSTGRES_DB=dbname` – Name it something snazzy.  
🔹 `-p 5432:5432` – Default Postgres port, but tweak as needed.  

🔗 Connecting Go to PostgreSQL
Check out the `main.go` file for how the magic happens:  
- Connect to PostgreSQL using `database/sql` and `pgx` (PostgreSQL driver).  
- Run queries like:  
  ```sql
  SELECT bday_year FROM users WHERE name = 'shayak';
  ```
- Insert new records (commented out for now).  

✨ Key Features of the Code:
Connection Pooling Efficient resource handling.  
Query Execution Fetch rows, loop through results.  
Error Handling – If something breaks, you'll know.  

🚧 TODO:
- Add transaction handling.  
- Build a proper schema migration.  
- Test inserting new users more aggressively.  
🛡️ Disclaimer
⚠️ **This is for dev purposes** – `sslmode=disable` and `trust` are great for local development but tighten it up before going live!  

Let me know if you get stuck or have fun ideas to extend this!
