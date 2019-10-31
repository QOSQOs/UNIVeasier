DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertUniversity(
  _id INT,
  _local_rank INT,
  _global_rank INT,
  _acronym VARCHAR(50),
  _name VARCHAR(200),
  _region INT,
  _description VARCHAR(1000),
  _latitude FLOAT,
  _longitude FLOAT,
  _logo BLOB,
  _is_verified INT,
  _doc_verifier BLOB,
  _reference VARCHAR(200),
  _type_university_id INT,
  _album_id INT,
  _created_date DATE,
  _last_modified_date DATE,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.university(id, local_rank, global_rank, acronym, name, region, description, latitude, longitude, logo, is_verified, doc_verifier, reference, type_university_id, album_id, created_date, last_modified_date, created_by) VALUES 
    (_id, _local_rank, _global_rank, _acronym, _name, _region, _description, _latitude, _longitude, _logo, _is_verified, _doc_verifier, _reference, _type_university_id, _album_id, _created_date, _last_modified_date, _created_by);
  ELSE
    UPDATE dbqosqos.university
    SET
      local_rank = _local_rank,
      global_rank = _global_rank,
      acronym = _acronym,
      name = _name,
      region = _region,
      description = _description,
      latitude = _latitude,
      longitude = _longitude,
      logo = _logo,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      reference = _reference,
      type_university_id = _type_university_id,
      album_id = _album_id,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetUniversityById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.university
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListUniversity()
BEGIN
    SELECT * FROM dbqosqos.university;
END //

DELIMITER ;