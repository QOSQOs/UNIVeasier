DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertAward(
  _id INT,
  _year INT,
  _period INT,
  _description VARCHAR(500),
  _event VARCHAR(100),
  _place INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _career_id INT,
  _album_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.award(year, period, description, event, place, is_verified, doc_verifier, created_date, last_modified_date, career_id, album_id, created_by) VALUES 
    (_year, _period, _description, _event, _place, _is_verified, _doc_verifier, _created_date, _last_modified_date, _career_id, _album_id, _created_by);
  ELSE
    UPDATE dbqosqos.award
    SET
      year = _year,
      period = _period,
      description = _description,
      event = _event,
      place = _place,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      career_id = _career_id,
      album_id = _album_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetAwardById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.award
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListAward()
BEGIN
    SELECT * FROM dbqosqos.award;
END //

DELIMITER ;