DELIMITER //

CREATE PROCEDURE dbtest.GetListTest()
BEGIN
    SELECT * FROM dbtest.tabletest;
END //

CREATE PROCEDURE dbtest.GetTest(parameter_name VARCHAR(255))
BEGIN
    SELECT * FROM dbtest.tabletest WHERE name = parameter_name;
END //

CREATE PROCEDURE AddTest(_name VARCHAR(255), _age INT, _email VARCHAR(50))
BEGIN
    INSERT INTO dbtest.tabletest(name, age, email) VALUES (_name, _age, _email);
END //

DELIMITER ;