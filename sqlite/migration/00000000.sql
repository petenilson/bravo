CREATE TABLE efforts (
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	strava_id     INTEGER NOT NULL UNIQUE,
	event_id      INTEGER NOT NULL REFERENCES events (id) ON DELETE CASCADE,
	segment_id    INTEGER NOT NULL REFERENCES segments (id) ON DELETE CASCADE,
	athlete_id    INTEGER NOT NULL REFERENCES athletes (id) ON DELETE CASCADE,
	start_date    TEXT NOT NULL,
	created_at    TEXT NOT NULL,
);

CREATE TABLE events (
	id                INTEGER PRIMARY KEY AUTOINCREMENT,
	name              TEXT NOT NULL,
	segment_id        INTEGER NOT NULL REFERENCES segments (id) ON DELETE CASCADE,
	creator_id        INTEGER NOT NULL REFERENCES athletes (id) ON DELETE CASCADE,
  invite_code       TEXT UNIQUE NOT NULL,
	is_active         INTEGER NOT NULL CHECK (Active IN (0, 1))
	participants      INTEGER NOT NULL,
	max_participants  INTEGER NOT NULL,
	start_time        TEXT NOT NULL,
  end_time          TEXT NOT NULL,
	created_at        TEXT NOT NULL,
	updated_at        TEXT NOT NULL,
);

CREATE TABLE event_memberships (
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id   INTEGER NOT NULL REFERENCES dials (id) ON DELETE CASCADE,
	athlete_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,

	UNIQUE(event_id, athlete_id)
);

CREATE TABLE segments (
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	strava_id     INTEGER NOT NULL UNIQUE,
	name          TEXT NOT NULL,
	city          TEXT NOT NULL,
	country       TEXT NOT NULL,
	updated_at    TEXT NOT NULL,
);

CREATE TABLE athletes (
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	strava_id     INTEGER NOT NULL UNIQUE,
	first_name    TEXT NOT NULL,
	last_name     TEXT NOT NULL,
	created_at    TEXT NOT NULL,
	updated_at    TEXT NOT NULL,
);

CREATE TABLE auths (
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id       INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
	access_token  TEXT NOT NULL,
	refresh_token TEXT NOT NULL,
	expiry        TEXT,
	created_at    TEXT NOT NULL,
	updated_at    TEXT NOT NULL,
);
