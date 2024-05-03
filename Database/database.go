package db_handler

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "socialNetwork.db")
	if err != nil {
		return nil, err
	}

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
        uuid TEXT NOT NULL PRIMARY KEY,
        username TEXT NOT NULL,
		first_name TEXT,
		last_name TEXT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		date_of_birth TEXT,
		about_me TEXT,
		have_image BOOL
    );`

	_, err = db.Exec(createUserTable)
	if err != nil {
		fmt.Println("Error creating user table")
		return nil, err
	}

	createFollowerTable := `CREATE TABLE IF NOT EXISTS user_follower (
		user_uuid TEXT NOT NULL,
		followed_by TEXT TEXT NOT NULL,
		is_accepted BOOL
	);`

	_, err = db.Exec(createFollowerTable)
	if err != nil {
		fmt.Println("Error creating follower table")
		return nil, err
	}

	createPostTable := `CREATE TABLE IF NOT EXISTS posts (
		uuid TEXT NOT NULL PRIMARY KEY,
		user_uuid TEXT NOT NULL,
		content TEXT NOT NULL,
		date TEXT NOT NULL,
		type TEXT,
		group_uuid TEXT,
		have_image BOOL
	);`

	_, err = db.Exec(createPostTable)
	if err != nil {
		fmt.Println("Error creating post table")
		return nil, err
	}

	createPrivatePostTable := `CREATE TABLE IF NOT EXISTS private_post(
		post_uuid TEXT NOT NULL,
		users_uuid TEXT NOT NULL
	);`

	_, err = db.Exec(createPrivatePostTable)
	if err != nil {
		fmt.Println("Error creating post table")
		return nil, err
	}

	createUserInGroupTable := `CREATE TABLE IF NOT EXISTS user_in_group(
		user_uuid TEXT NOT NULL,
		group_uuid TEXT NOT NULL,
		status TEXT
	);`

	_, err = db.Exec(createUserInGroupTable)
	if err != nil {
		fmt.Println("Error creating UserInGroup table")
		return nil, err
	}

	createUserEventGroupTable := `CREATE TABLE IF NOT EXISTS user_event_group(
		event_uuid TEXT NOT NULL,
		user_uuid TEXT NOT NULL,
		status TEXT
	);`

	_, err = db.Exec(createUserEventGroupTable)
	if err != nil {
		fmt.Println("Error creating UserEventGroup table")
		return nil, err
	}

	createNotificationsTable := `CREATE TABLE IF NOT EXISTS notifications(
		uuid TEXT NOT NULL PRIMARY KEY,
		user_receive_uuid NOT NULL,
		user_send_uuid NOT NULL,
		group_send_uuid,
		type TEXT,
		is_read BOOL NOT NULL,
		created_at TEXT NOT NULL,
		related_id TEXT
	);`

	_, err = db.Exec(createNotificationsTable)
	if err != nil {
		fmt.Println("Error creating Notification table")
		return nil, err
	}

	createGroupEventTable := `CREATE TABLE IF NOT EXISTS group_event(
		uuid TEXT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		group_uuid TEXT NOT NULL,
		description TEXT NOT NULL,
		date_start TEXT NOT NULL,
		date_end TEXT NOT NULL
	);`

	_, err = db.Exec(createGroupEventTable)
	if err != nil {
		fmt.Println("Error creating GroupEvent table")
		return nil, err
	}

	createGroupTable := `CREATE TABLE IF NOT EXISTS groups(
		uuid TEXT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		have_image BOOL,
		creation_date TEXT NOT NULL,
		created_by_user_uuid TEXT NOT NULL
	);`

	_, err = db.Exec(createGroupTable)
	if err != nil {
		fmt.Println("Error creating Group table")
		return nil, err
	}

	createCommentsTable := `CREATE TABLE IF NOT EXISTS comments(
		uuid TEXT NOT NULL PRIMARY KEY,
		user_uuid TEXT NOT NULL,
		post_uuid TEXT NOT NULL,
		content TEXT NOT NULL,
		date TEXT NOT NULL
	);`

	_, err = db.Exec(createCommentsTable)
	if err != nil {
		fmt.Println("Error creating Comments table")
		return nil, err
	}

	createGroupChatTable := `CREATE TABLE IF NOT EXISTS group_chat(
		user_uuid TEXT NOT NULL,
		group_uuid TEXT NOT NULL,
		date TEXT NOT NULL,
		content TEXT NOT NULL
	);`

	_, err = db.Exec(createGroupChatTable)
	if err != nil {
		fmt.Println("Error creating GroupChat table")
		return nil, err
	}

	createUserChatTable := `CREATE TABLE IF NOT EXISTS user_chat(
		sender_uuid TEXT NOT NULL,
		receiver_uuid TEXT NOT NULL,
		date TEXT NOT NULL,
		content TEXT NOT NULL
	);`

	_, err = db.Exec(createUserChatTable)
	if err != nil {
		fmt.Println("Error creating UserChat table")
		return nil, err
	}

	return db, nil
}
