db.createUser(
  {
    user: 'mock user',
    pwd: 'mock pwd',
    roles: [
      {
	role: "readWrite",
	db: "harvest-db"
      }
    ]
  }
)
