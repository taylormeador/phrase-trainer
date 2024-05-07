package db

// TODO: link the user_id once user data is implemented
var userUploadsDDL string = `
	CREATE TABLE IF NOT EXISTS user_uploads (
		id SERIAL PRIMARY KEY,
		user_id INT,
		timestamp TIMESTAMP,
		file_name varchar(128),
		file_label varchar(128),
		blob_name char(36),
		UNIQUE(user_id, file_name)
	);
`
