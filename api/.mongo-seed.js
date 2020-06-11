async function init() {
  rs.initiate();
  rs.status();

  await new Promise(r => setTimeout(r, 2000));

  db.createUser(
    {
      user: "mock-user",
      pwd: "mock-pwd",
      roles: [
        {
  	      role: "readWrite",
  	      db: "mock-db"
        }
      ]
    }
  );
}

init();
