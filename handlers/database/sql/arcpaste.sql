-- name: create-paste
INSERT INTO paste
VALUES (id = ?,
        user_id = ?,
        raw = ?,
        language = ?,
        expire = ?,
        password = ?,
        users = ?)
