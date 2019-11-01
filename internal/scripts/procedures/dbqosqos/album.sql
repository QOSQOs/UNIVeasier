DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertAlbum(
  _id INT,
  _name VARCHAR(200),
  _description VARCHAR(1000),
  _created_date DATE,
  _last_modified_date DATE,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.album(name, description, created_date, last_modified_date, created_by) VALUES 
    (_name, _description, _created_date, _last_modified_date, _created_by);
  ELSE
    UPDATE dbqosqos.album
    SET
      name = _name,
      description = _description,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetAlbumById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.album
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListAlbum()
BEGIN
    SELECT * FROM dbqosqos.album;
END //

DELIMITER ;