DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertCareer(
  _id INT,
  _acronym VARCHAR(50),
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _logo BLOB,
  _is_verified INT,
  _doc_verifier BLOB,
  _reference VARCHAR(200),
  _created_date DATE,
  _last_modified_date DATE,
  _faculty_id INT,
  _album_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.career(id, acronym, name, description, logo, is_verified, doc_verifier, reference, created_date, last_modified_date, faculty_id, album_id, created_by) VALUES 
    (_id, _acronym, _name, _description, _logo, _is_verified, _doc_verifier, _reference, _created_date, _last_modified_date, _faculty_id, _album_id, _created_by);
  ELSE
    UPDATE dbqosqos.career
    SET
      acronym = _acronym,
      name = _name,
      description = _description,
      logo = _logo,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      reference = _reference,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      faculty_id = _faculty_id,
      album_id = _album_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetCareerById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.career
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListCareer()
BEGIN
    SELECT * FROM dbqosqos.career;
END //

DELIMITER ;