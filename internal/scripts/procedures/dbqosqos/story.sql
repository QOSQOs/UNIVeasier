DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertStory(
  _id INT,
  _title VARCHAR(100),
  _body VARCHAR(500),
  _recommended INT,
  _unrecommended INT,
  _views INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _allow_comments TINYINT(1),
  _layer INT,
  _created_date DATE,
  _last_modified_date DATE,
  _created_by INT,
  _career_id INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.story(title, body, recommended, unrecommended, views, is_verified, doc_verifier, allow_comments, layer, created_date, last_modified_date, created_by, career_id) VALUES 
    (_title, _body, _recommended, _unrecommended, _views, _is_verified, _doc_verifier, _allow_comments, _layer, _created_date, _last_modified_date, _created_by, _career_id);
  ELSE
    UPDATE dbqosqos.story
    SET
      title = _title,
      body = _body,
      recommended = _recommended,
      unrecommended = _unrecommended,
      views = _views,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      allow_comments = _allow_comments,
      layer = _layer,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      created_by = _created_by,
      career_id = _career_id
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetStoryById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.story
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListStory()
BEGIN
    SELECT * FROM dbqosqos.story;
END //

DELIMITER ;