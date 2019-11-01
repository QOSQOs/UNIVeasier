DELIMITER // 

CREATE PROCEDURE dbqosqos.InsertSchedule(
  _id INT,
  _start DATE,
  _end DATE,
  _day INT,
  _room VARCHAR(50),
  _type_lesson INT,
  _is_verified INT,
  _doc_verifier BLOB,
  _created_date DATE,
  _last_modified_date DATE,
  _course_id INT,
  _created_by INT)
BEGIN
  IF _id IS NULL THEN
    INSERT INTO dbqosqos.schedule(start, end, day, room, type_lesson, is_verified, doc_verifier, created_date, last_modified_date, course_id, created_by) VALUES 
    (_start, _end, _day, _room, _type_lesson, _is_verified, _doc_verifier, _created_date, _last_modified_date, _course_id, _created_by);
  ELSE
    UPDATE dbqosqos.schedule
    SET
      start = _start,
      end = _end,
      day = _day,
      room = _room,
      type_lesson = _type_lesson,
      is_verified = _is_verified,
      doc_verifier = _doc_verifier,
      created_date = _created_date,
      last_modified_date = _last_modified_date,
      course_id = _course_id,
      created_by = _created_by
    WHERE
      id = _id;
  END IF;
END //

CREATE PROCEDURE dbqosqos.GetScheduleById(_id INT)
BEGIN
    SELECT * FROM dbqosqos.schedule
    WHERE id = _id;
END //

CREATE PROCEDURE dbqosqos.GetListSchedule()
BEGIN
    SELECT * FROM dbqosqos.schedule;
END //

DELIMITER ;