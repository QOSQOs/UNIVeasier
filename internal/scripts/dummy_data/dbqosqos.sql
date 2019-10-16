INSERT INTO dbqosqos.person(code, gender, first_name, last_name, phone, email, avatar, type, home_city, current_city, ethnic, nationality, birth_date, admission_year, period, is_verified, doc_verifier, created_date, last_modified_date) VALUES
('192619', 0, 'Pepito', 'Quispe Quispe', '984867072', 'pepito@gmail.com', NULL, 0, 'Sicuani', 'Cusco', NULL, 'peruvian', '1995-02-22', NULL, NULL, 0, NULL, now(), now()),
('192618', 1, 'Pepito1', 'Quispe Quispe', '984867073', 'pepito1@gmail.com', NULL, 0, 'Urcos', 'Cusco', NULL, 'peruvian', '1995-01-22', NULL, NULL, 1, NULL, now(), now()),
('192617', 0, 'Pepito2', 'Quispe Quispe', '984867074', 'pepito2@gmail.com', NULL, 1, 'Tinta', 'Cusco', NULL, 'peruvian', '1994-02-20', 2011, 1, 0, NULL, now(), now()),
('192616', 1, 'Pepito3', 'Quispe Quispe', '984867075', 'pepito3@gmail.com', NULL, 1, 'Urubamba', 'Cusco', NULL, 'peruvian', '1995-05-12', 2011, 2, 2, NULL, now(), now()),
('192615', 0, 'Pepito4', 'Quispe Quispe', '984867076', 'pepito4@gmail.com', NULL, 2, 'Sicuani', 'Cusco', NULL, 'peruvian', '1995-01-11', 2004, 1, 1, NULL, now(), now()),
('192614', 1, 'Pepito5', 'Quispe Quispe', '984867077', 'pepito5@gmail.com', NULL, 2, 'Calca', 'Cusco', NULL, 'peruvian', '1995-05-19', 2003, 2, 0, NULL, now(), now()),
('192614', 1, 'Professor', 'Professor', '984867067', 'professor@gmail.com', NULL, 3, 'Cusco', 'Cusco', NULL, 'peruvian', '1990-05-19', 2000, 1, 1, NULL, now(), now());

SELECT * FROM dbqosqos.person;

INSERT INTO dbqosqos.type_university(name, description, is_verified, doc_verifier, created_date, last_modified_date, created_by) VALUES 
('public', 'this is a public university', 1, NULL, now(), now(), 4),
('private', 'this is a private university', 0, NULL, now(), now(), 6);

SELECT * FROM dbqosqos.type_university;

INSERT INTO dbqosqos.university(local_rank, global_rank, acronym, name, region, description, latitude, longitude, logo, is_verified, doc_verifier, reference, type_university_id, album_id, created_date, last_modified_date, created_by) VALUES 
(1, 10, 'UNSAAC', 'San Antonio Abap of Cusco', 1, 'this is a description for UNSAAC', 13.5219, 71.9583, NULL, 1, NULL, 'unsaac.edu.com', 1, NULL, now(), now(), 4),
(2, 100, 'UAC', 'Andina of Cusco', 1, 'this a description for UAC', 13.5219, 71.9583, NULL, 0, NULL, 'uac.edu.com', 2, NULL, now(), now(), 4);

SELECT * FROM dbqosqos.university;

INSERT INTO dbqosqos.faculty(acronym, name, description, logo, is_verified, doc_verifier, created_date, last_modified_date, university_id, created_by) VALUES
('FIIES', 'Facultad de Ing. Informatica, Electronica, etc.', 'this is a faculty of UNSAAC', NULL, 1, NULL, now(), now(), 1, 4),
('FICU', 'Facultad de Ing. Civil y Urbanismo', 'this is a faculty of UNSAAC', NULL, 0, NULL, now(), now(), 1, 6);

SELECT * FROM dbqosqos.faculty;

INSERT INTO dbqosqos.career(acronym, name, description, logo, is_verified, doc_verifier, reference, created_date, last_modified_date, faculty_id, album_id, created_by) VALUES
('IIDS', 'Ing. Informatica y de Systemas', 'this is a career of FIIES', NULL, 1, NULL, 'info.edu.com', now(), now(), 1, NULL, 4),
('IC', 'Ing. Civil', 'this is a career of FICU', NULL, 0, NULL, 'civil.edu.com', now(), now(), 2, NULL, 6);

SELECT * FROM dbqosqos.career;

INSERT INTO dbqosqos.career_person(career_id, person_id, is_verified, doc_verifier, created_date, last_modified_date) VALUES
(1, 3, 1, NULL, now(), now()),
(1, 4, 0, NULL, now(), now()),
(1, 5, 1, NULL, now(), now()),
(1, 6, 0, NULL, now(), now());

SELECT * FROM dbqosqos.career_person;

INSERT INTO dbqosqos.award(year, period, description, event, place, is_verified, doc_verifier, created_date, last_modified_date, career_id, album_id, created_by) VALUES
(2010, 1, 'basketball championship', 'basketball championship', 1, 0, NULL, now(), now(), 1, NULL, 4),
(2010, 1, 'chess championship', 'chess championship', 2, 1, NULL, now(), now(), 1, NULL, 6),
(2010, 1, 'football championship', 'football championship', 3, 0, NULL, now(), now(), 1, NULL, 4),
(2010, 1, 'volleyball championship', 'volleyball championship', 1, 1, NULL, now(), now(), 1, NULL, 6),
(2010, 1, 'swimming championship', 'swimming championship', 2, 0, NULL, now(), now(), 1, NULL, 6);

SELECT * FROM dbqosqos.award;

INSERT INTO dbqosqos.type_course(acronym, name, description, is_verified, doc_verifier, created_date, last_modified_date, required_credits, career_id, created_by) VALUES
('EE', 'electivo especifico', 'this is a type of course', 0, NULL, now(), now(), 10, 1, 4),
('OE', 'obligatorio especifico', 'this is a type of course', 0, NULL, now(), now(), 100, 1, 4),
('EG', 'elective general', 'this is a type of course', 0, NULL, now(), now(), 20, 1, 4),
('EGT', 'elective general T.', 'this is a type of course', 0, NULL, now(), now(), 30, 1, 4),
('OEFE', 'obligatorio EFE.', 'this is a type of course', 0, NULL, now(), now(), 10, 1, 4),
('OEES', 'obligatorio OEES.', 'this is a type of course', 0, NULL, now(), now(), 10, 1, 4),
('AC', 'actividades', 'this is a type of course', 0, NULL, now(), now(), 3, 1, 4),
('SEM', 'seminario', 'this is a type of course', 0, NULL, now(), now(), 3, 1, 4),
('AEX', 'actividades EX.', 'this is a type of course', 0, NULL, now(), now(), 3, 1, 4),
('EEFE', 'electivo EFE.', 'this is a type of course', 0, NULL, now(), now(), 5, 1, 4);

SELECT * FROM dbqosqos.type_course;

INSERT INTO dbqosqos.course(code, name, description, credits, batch, year, period, semester, total_hours, is_verified, doc_verifier, created_date, last_modified_date, type_course_id, career_id, professor, created_by) VALUES
('IF607', 'BIOINFORMATICA', 'this is a course', 4, 1, 2019, 1, 0, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF404', 'LENGUAJE ENSAMBLADOR', 'this is a course', 3, 1, 2019, 1, 0, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('ME165', 'MATEMATICA BASICA II', 'this is a course', 4, 1, 2019, 1, 0, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF606', 'PROCESAMIENTO DE LENGUAJE NATURAL', 'this is a course', 4, 1, 2019, 1, 0, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF214', 'TEORIA GENERAL DE SISTEMAS', 'this is a course', 4, 1, 2019, 1, 0, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('DE901', 'CONSTITUCION POLITICA Y DERECHOS HUMANOS', 'this is a course', 3, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('DE901', 'CONSTITUCION POLITICA Y DERECHOS HUMANOS', 'this is a course', 3, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ED901', 'ESTRATEGIAS DE APRENDIZAJE AUTONOMO', 'this is a course', 4, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ED901', 'ESTRATEGIAS DE APRENDIZAJE AUTONOMO', 'this is a course', 4, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('FP901', 'FILOSOFIA Y ETICA', 'this is a course', 3, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('FP901', 'FILOSOFIA Y ETICA', 'this is a course', 3, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ME901', 'MATEMATICA I', 'this is a course', 4, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ME901', 'MATEMATICA I', 'this is a course', 4, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('LC901', 'REDACCION DE TEXTOS', 'this is a course', 4, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('LC901', 'REDACCION DE TEXTOS', 'this is a course', 4, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('AS901', 'SOCIEDAD Y CULTURA', 'this is a course', 3, 1, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('AS901', 'SOCIEDAD Y CULTURA', 'this is a course', 3, 2, 2019, 1, 1, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ME903', 'CÁLCULO I', 'this is a course', 4, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 4, 1, 7, 4),
('ME903', 'CÁLCULO I', 'this is a course', 4, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 4, 1, 7, 4),
('FI902', 'FÍSICA I', 'this is a course', 4, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 4, 1, 7, 4),
('FI902', 'FÍSICA I', 'this is a course', 4, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 4, 1, 7, 4),
('IF468', 'FUNDAMENTOS DE LA PROGRAMACIÓN', 'this is a course', 4, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 5, 1, 7, 4),
('IF468', 'FUNDAMENTOS DE LA PROGRAMACIÓN', 'this is a course', 4, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 5, 1, 7, 4),
('FP902', 'LIDERAZGO Y HABILIDADES SOCIALES', 'this is a course', 3, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('FP902', 'LIDERAZGO Y HABILIDADES SOCIALES', 'this is a course', 3, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('ME307', 'MATEMÁTICAS DISCRETAS I', 'this is a course', 4, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 5, 1, 7, 4),
('ME307', 'MATEMÁTICAS DISCRETAS I', 'this is a course', 4, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 5, 1, 7, 4),
('IF902', 'TECNOLOGIAS DE LA INFORMACIÓN Y LA COMUNICACIÓN', 'this is a course', 3, 1, 2019, 1, 2, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('IF902', 'TECNOLOGIAS DE LA INFORMACIÓN Y LA COMUNICACIÓN', 'this is a course', 3, 2, 2019, 1, 2, 35, 0, NULL, now(), now(), 3, 1, 7, 4),
('IF450', 'ABSTRACCIÓN DE DATOS Y OBJETOS', 'this is a course', 4, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF450', 'ABSTRACCIÓN DE DATOS Y OBJETOS', 'this is a course', 4, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME351', 'ALGEBRA LINEAL', 'this is a course', 4, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME351', 'ALGEBRA LINEAL', 'this is a course', 4, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME350', 'CALCULO II', 'this is a course', 4, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME350', 'CALCULO II', 'this is a course', 4, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('FI370', 'FISICA APLICADA', 'this is a course', 4, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('FI370', 'FISICA APLICADA', 'this is a course', 4, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME352', 'PROBABILIDADES Y ESTADÍSTICA', 'this is a course', 4, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME352', 'PROBABILIDADES Y ESTADÍSTICA', 'this is a course', 4, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF451', 'PROGRAMACION I', 'this is a course', 2, 1, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF451', 'PROGRAMACION I', 'this is a course', 2, 2, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF451', 'PROGRAMACION I', 'this is a course', 2, 3, 2019, 1, 3, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF480', 'ADMINISTRACIÓN DE TECNOLOGIAS DE INFORMACIÓN', 'this is a course', 3, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF452', 'ALGORITMOS Y ESTRUCTURAS DE DATOS', 'this is a course', 4, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF452', 'ALGORITMOS Y ESTRUCTURAS DE DATOS', 'this is a course', 4, 2, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('EL371', 'ELECTRÓNICA Y DISEÑO DIGITAL', 'this is a course', 3, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('EL371', 'ELECTRÓNICA Y DISEÑO DIGITAL', 'this is a course', 3, 2, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME354', 'INVESTIGACIÓN OPERATIVA', 'this is a course', 4, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME355', 'MATEMÁTICA DISCRETAS II', 'this is a course', 4, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME355', 'MATEMÁTICA DISCRETAS II', 'this is a course', 4, 2, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF060', 'MÚSICA', 'this is a course', 2, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 9, 1, 7, 4),
('IF453', 'PROGRAMACIÓN II', 'this is a course', 2, 1, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF453', 'PROGRAMACIÓN II', 'this is a course', 2, 2, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF453', 'PROGRAMACIÓN II', 'this is a course', 2, 3, 2019, 1, 4, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF610', 'ANALISIS Y DISEÑO DE SISTEMAS DE INFORMACIÓN', 'this is a course', 4, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME356', 'ECUACIONES DIFERENCIALES', 'this is a course', 4, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('ME356', 'ECUACIONES DIFERENCIALES', 'this is a course', 4, 2, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF481', 'INGENIERIA ECONOMICA', 'this is a course', 3, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF650', 'MODELOS PROBABILISTICOS', 'this is a course', 4, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF550', 'ORGANIZACIÓN Y ARQUITECTURA DEL COMPUTADOR', 'this is a course', 4, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF454', 'TEORIA DE LA COMPUTACION', 'this is a course', 3, 1, 2019, 1, 5, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF455', 'ALGORITMOS PARALELOS Y DISTRIBUIDOS', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF458', 'COMPUTACIÓN GRÁFICA I', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 10, 1, 7, 4),
('IF612', 'FUNDAMENTOS Y DISEÑO DE BASES DE DATOS', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF413', 'LABORATORIO V', 'this is a course', 2, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF611', 'METODOLOGÍA DE DESARROLLO DE SOFTWARE', 'this is a course', 3, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF457', 'METODOS NUMERICOS', 'this is a course', 3, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF461', 'METODOS NUMERICOS', 'this is a course', 3, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF502', 'MICROPROCESADORES', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF301', 'SISTEMAS DE BASES DE DATOS I', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF202', 'SISTEMAS DE INFORMACION II', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF551', 'SISTEMAS OPERATIVOS', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 6, 1, 7, 4),
('IF533', 'SISTEMAS OPERATIVOS I', 'this is a course', 4, 1, 2019, 1, 6, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF420', 'ANALISIS Y DISEÑO DE ALGORITMOS', 'this is a course', 4, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF406', 'COMPUTACION GRAFICA I', 'this is a course', 4, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('CO152', 'CONTABILIDAD GENERAL I', 'this is a course', 3, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('ME766', 'INVESTIGACION DE OPERACIONES I', 'this is a course', 4, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF414', 'LABORATORIO VI', 'this is a course', 2, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF414', 'LABORATORIO VI', 'this is a course', 2, 2, 2019, 1, 7, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF302', 'SISTEMAS DE BASES DE DATOS II', 'this is a course', 4, 1, 2019, 1, 7, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF053', 'ACTIVIDADES DE PRODUCCION DE BIENES Y PRESTACION DE SERVICIOS', 'this is a course', 1, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 7, 1, 7, 4),
('IF203', 'ADMINISTRACION DE CENTROS DE COMPUTO', 'this is a course', 4, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF601', 'INTELIGENCIA ARTIFICIAL', 'this is a course', 4, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF471', 'MODELOS Y SIMULACION', 'this is a course', 4, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF513', 'PROGRAMACION CONCURRENTE Y PARALELA', 'this is a course', 4, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF505', 'REDES Y TELEPROCESO I', 'this is a course', 4, 1, 2019, 1, 8, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF425', 'CONSTRUCCION DE COMPILADORES', 'this is a course', 4, 1, 2019, 1, 9, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('EC702', 'COSTOS Y PRESUPUESTOS', 'this is a course', 3, 1, 2019, 1, 9, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF209', 'INGENIERIA DE SOFTWARE', 'this is a course', 4, 1, 2019, 1, 9, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF506', 'REDES Y TELEPROCESO II', 'this is a course', 4, 1, 2019, 1, 9, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF213', 'SISTEMAS DE INFORMACION GERENCIAL', 'this is a course', 3, 1, 2019, 1, 9, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF210', 'PLANEAMIENTO ESTRATEGICO', 'this is a course', 4, 1, 2019, 1, 10, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF603', 'ROBOTICA Y PROCESAMIENTO DE SEÑAL', 'this is a course', 4, 1, 2019, 1, 10, 35, 0, NULL, now(), now(), 2, 1, 7, 4),
('IF212', 'SEGURIDAD CONTROL Y AUDITORIA DE SISTEMAS DE INFORMACION', 'this is a course', 4, 1, 2019, 1, 10, 35, 0, NULL, now(), now(), 1, 1, 7, 4),
('IF001', 'SEMINARIO EN INFORMATICA', 'this is a course', 3, 1, 2019, 1, 10, 35, 0, NULL, now(), now(), 8, 1, 7, 4);

SELECT * FROM dbqosqos.course;

INSERT INTO dbqosqos.courses_connection(year, period, threshold_credits, child_course, parent_course, created_date, last_modified_date, career_id, created_by) VALUES
(2019, 1, NULL, 'IF422', 'IF404', now(), now(), 1, 4),
(2019, 1, NULL, 'IF410', 'IF404', now(), now(), 1, 4),
(2019, 1, NULL, 'IF601', 'IF606', now(), now(), 1, 4),
(2019, 1, 70, NULL, 'IF214', now(), now(), 1, 4),
(2019, 1, NULL, 'ME901', 'ME903', now(), now(), 1, 4),	
(2019, 1, 15, NULL, 'FI902', now(), now(), 1, 4),
(2019, 1, 15, NULL, 'FP902', now(), now(), 1, 4),
(2019, 1, 15, NULL, 'IF902', now(), now(), 1, 4),
(2019, 1, NULL, 'IF468', 'IF450', now(), now(), 1, 4),
(2019, 1, NULL, 'ME307', 'ME351', now(), now(), 1, 4),
(2019, 1, NULL, 'ME903', 'ME350', now(), now(), 1, 4),
(2019, 1, NULL, 'FI902', 'FI370', now(), now(), 1, 4),
(2019, 1, NULL, 'ME903', 'ME352', now(), now(), 1, 4),
(2019, 1, NULL, 'IF468', 'IF451', now(), now(), 1, 4),
(2019, 1, NULL, 'ME352', 'IF480', now(), now(), 1, 4),
(2019, 1, NULL, 'IF450', 'IF452', now(), now(), 1, 4),
(2019, 1, NULL, 'FI370', 'EL371', now(), now(), 1, 4),
(2019, 1, NULL, 'ME351', 'ME354', now(), now(), 1, 4),
(2019, 1, NULL, 'ME307', 'ME355', now(), now(), 1, 4),
(2019, 1, 60, NULL, 'IF060', now(), now(), 1, 4),
(2019, 1, NULL, 'IF451', 'IF453', now(), now(), 1, 4),
(2019, 1, NULL, 'IF450', 'IF610', now(), now(), 1, 4),
(2019, 1, NULL, 'IF480', 'IF610', now(), now(), 1, 4),
(2019, 1, NULL, 'ME350', 'ME356', now(), now(), 1, 4),
(2019, 1, NULL, 'IF480', 'IF481', now(), now(), 1, 4),
(2019, 1, NULL, 'ME352', 'IF650', now(), now(), 1, 4),
(2019, 1, NULL, 'EL371', 'IF550', now(), now(), 1, 4),
(2019, 1, NULL, 'IF452', 'IF454', now(), now(), 1, 4),
(2019, 1, NULL, 'ME355', 'IF454', now(), now(), 1, 4),
(2019, 1, NULL, 'IF452', 'IF455', now(), now(), 1, 4),
(2019, 1, NULL, 'IF452', 'IF458', now(), now(), 1, 4),
(2019, 1, NULL, 'ME351', 'IF458', now(), now(), 1, 4),
(2019, 1, NULL, 'IF610', 'IF612', now(), now(), 1, 4),
(2019, 1, NULL, 'IF412', 'IF413', now(), now(), 1, 4),
(2019, 1, NULL, 'IF610', 'IF611', now(), now(), 1, 4),
(2019, 1, NULL, 'ME356', 'IF457', now(), now(), 1, 4),
(2019, 1, NULL, 'ME253', 'IF461', now(), now(), 1, 4),
(2019, 1, NULL, 'IF422', 'IF461', now(), now(), 1, 4),
(2019, 1, NULL, 'IF501', 'IF502', now(), now(), 1, 4),
(2019, 1, NULL, 'IF231', 'IF301', now(), now(), 1, 4),
(2019, 1, NULL, 'ME253', 'IF301', now(), now(), 1, 4),
(2019, 1, NULL, 'IF231 ', 'IF202', now(), now(), 1, 4),
(2019, 1, NULL, 'IF201', 'IF202', now(), now(), 1, 4),
(2019, 1, NULL, 'IF550', 'IF551', now(), now(), 1, 4),
(2019, 1, NULL, 'IF403', 'IF533', now(), now(), 1, 4),
(2019, 1, NULL, 'IF411', 'IF533', now(), now(), 1, 4),
(2019, 1, NULL, 'IF403', 'IF420', now(), now(), 1, 4),
(2019, 1, NULL, 'IF403', 'IF406', now(), now(), 1, 4),
(2019, 1, NULL, 'IF402', 'CO152', now(), now(), 1, 4),
(2019, 1, NULL, 'IF422', 'CO152', now(), now(), 1, 4),
(2019, 1, NULL, 'ME360', 'ME766', now(), now(), 1, 4),
(2019, 1, NULL, 'IF301', 'IF414', now(), now(), 1, 4),
(2019, 1, NULL, 'IF301', 'IF302', now(), now(), 1, 4),
(2019, 1, 130, NULL, 'IF053', now(), now(), 1, 4),
(2019, 1, 130, NULL, 'IF203', now(), now(), 1, 4),
(2019, 1, NULL, 'IF301', 'IF601', now(), now(), 1, 4),
(2019, 1, NULL, 'ME660', 'IF471', now(), now(), 1, 4),
(2019, 1, NULL, 'IF501', 'IF513', now(), now(), 1, 4),
(2019, 1, NULL, 'IF403', 'IF513', now(), now(), 1, 4),
(2019, 1, NULL, 'IF533', 'IF505', now(), now(), 1, 4),
(2019, 1, 150, NULL, 'IF425', now(), now(), 1, 4),
(2019, 1, 100, NULL, 'EC702', now(), now(), 1, 4),
(2019, 1, NULL, 'IF202', 'IF209', now(), now(), 1, 4),
(2019, 1, NULL, 'IF505', 'IF506', now(), now(), 1, 4),
(2019, 1, NULL, 'IF202', 'IF213', now(), now(), 1, 4),
(2019, 1, NULL, 'IF209', 'IF210', now(), now(), 1, 4),
(2019, 1, NULL, 'IF601', 'IF603', now(), now(), 1, 4),
(2019, 1, NULL, 'IF213', 'IF212', now(), now(), 1, 4),
(2019, 1, 180, NULL, 'IF001', now(), now(), 1, 4);

SELECT * FROM dbqosqos.courses_connection;

INSERT INTO dbqosqos.interest(tag, is_skill, created_date, last_modified_date, created_by) VALUES
('machine learning', 1, now(), now(), 4),
('big data', 0, now(), now(), 4),
('robotic', 0, now(), now(), 4),
('machine learning', 0, now(), now(), 5),
('software', 1, now(), now(), 6),
('data visualization', 0, now(), now(), 6);

SELECT * FROM dbqosqos.interest;

INSERT INTO dbqosqos.story(title, body, recommended, unrecommended, views, is_verified, doc_verifier, allow_comments, layer, created_date, last_modified_date, created_by, career_id) VALUES
('story 1', 'this is the body of the story 1', 2, 2, 4, 0, NULL, 1, 2, now(), now(), 4, 1),
('story 2', 'this is the body of the story 2', 0, 4, 4, 0, NULL, 1, 3, now(), now(), 6, 1),
('story 3', 'this is the body of the story 3', 4, 0, 4, 1, NULL, 1, 2, now(), now(), 4, 1),
('story 4', 'this is the body of the story 4', 0, 0, 0, 1, NULL, 0, 3, now(), now(), 6, 1);

SELECT * FROM dbqosqos.story;

INSERT INTO dbqosqos.comment(content, likes, dislikes, created_date, last_modified_date, story_id, created_by) VALUES
('this is the comment 1 for the story 1', 1, 2, now(), now(), 1, 5),
('this is the comment 2 for the story 1', 1, 2, now(), now(), 1, 6),
('this is the comment 1 for the story 3', 6, 1, now(), now(), 1, 4);

SELECT * FROM dbqosqos.comment;

INSERT INTO dbqosqos.tag_container(tag, agree, disagree, reference, type, layer, created_date, last_modified_date, is_verified, doc_verifier, career_id, created_by) VALUES
('logic reasoning', 4, 1, 'logic-reaso.com.pe', 1, 2, now(), now(), 0, NULL, 1, 6),
('math', 4, 0, 'math.com.pe', 1, 2, now(), now(), 0, NULL, 1, 4),
('machine learning', 4, 1, 'machine-learning.com.pe', 2, 2, now(), now(), 0, NULL, 1, 6),
('web developer', 2, 2, 'web-developer.com.pe', 2, 2, now(), now(), 0, NULL, 1, 4),
('artificial intelligence', 4, 1, 'arti-intell.com.pe', 2, 3, now(), now(), 0, NULL, 1, 6),
('sofware engineer', 3, 1, 'software-engineer.com.pe', 2, 3, now(), now(), 0, NULL, 1, 5),
('facebook', 4, 0, 'web-developer.com.pe', 3, 3, now(), now(), 1, NULL, 1, 6),
('amazon', 4, 0, 'arti-intell.com.pe', 3, 3, now(), now(), 0, NULL, 1, 6),
('SAP', 2, 1, 'software-engineer.com.pe', 3, 3, now(), now(), 1, NULL, 1, 6);

SELECT * FROM dbqosqos.comment;
