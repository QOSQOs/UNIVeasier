DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertPhoto(
  _id INT,
  _pic BLOB,
  _description VARCHAR(200),
  _created_date DATE,
  _last_modified_date DATE,
  _album_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.photo(pic, description, created_date, last_modified_date, album_id, created_by) VALUES 
    (_pic, _description, _created_date, _last_modified_date, _album_id, _created_by);
  ELSE
    UPDATE dbqosqos.photo
    SET
      pic = _pic,
      description = _description,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      album_id = _album_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetPhotoById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.photo
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListPhoto()
BEGIN
    SELECT * FROM dbqosqos.photo;
END //

DELIMITER ;