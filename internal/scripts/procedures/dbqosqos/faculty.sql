DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertFaculty(
  _id INT,
  _acronym VARCHAR(50),
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _logo BLOB,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _university_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.faculty(acronym, name, description, logo, is_verified, doc_verifier, created_date, last_modified_date, university_id, created_by) VALUES 
    (_acronym, _name, _description, _logo, _is_verified, _doc_verifier, _created_date, _last_modified_date, _university_id, _created_by);
  ELSE
    UPDATE dbqosqos.faculty
    SET
      acronym = _acronym,
      name = _name,
      description = _description,
      logo = _logo,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      university_id = _university_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetFacultyById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.faculty
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListFaculty()
BEGIN
    SELECT * FROM dbqosqos.faculty;
END //

DELIMITER ;