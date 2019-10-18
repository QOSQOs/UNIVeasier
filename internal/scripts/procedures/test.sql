DELIMITER //

CREATE PROCEDURE dbtest.GetListTest()
BEGIN
    SELECT * FROM dbtest.tabletest;
END //

CREATE PROCEDURE dbtest.GetTest(parameter_name VARCHAR(255))
BEGIN
    SELECT * FROM dbtest.tabletest WHERE name = parameter_name;
END //

DELIMITER ;