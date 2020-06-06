db.createUser(
  {
    user: "mock-user",
    pwd: "mock-pwd",
    roles: [
      {
	      role: "root",
	      db: "admin"
      },
      {
	      role: "readWrite",
	      db: "mock-db"
      }
    ]
  }
)
