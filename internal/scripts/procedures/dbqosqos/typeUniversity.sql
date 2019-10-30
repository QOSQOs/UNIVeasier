DELIMITER //

/*
  Insert a new record for the TypeUniversity Table if the _id parameter is null,
  otherwise, the record already exist and we only need to update it.
*/
CREATE PROCEDURE dbqosqos.InsertTypeUniversity(
  _id INT,
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.type_university(id, name, description, is_verified, doc_verifier, created_date, last_modified_date, created_by) VALUES
    (_id, _name, _description, _is_verified, _doc_verifier, _created_date, _last_modified_date, _created_by);
  ELSE
    UPDATE dbqosqos.type_university
    SET
      name = _name,
      description = _description,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

/*
  Get a TypeUniversity record by its _id
*/
CREATE PROCEDURE dbqosqos.GetTypeUniversityById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.type_university 
    WHERE id = _id;
END //

/*
  Get a TypeUniversity record by its _id
*/
CREATE PROCEDURE dbqosqos.GetListTypeUniversity()
BEGIN
    SELECT * FROM dbqosqos.type_university;
END //

DELIMITER ;