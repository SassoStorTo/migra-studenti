SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, SC.IdS, SC.IdC, SC.CreationDate
          FROM students AS S INNER JOIN
               studentclass AS SC ON S.Id = SC.IdS
          WHERE SC.CreationDate = (SELECT MAX(CreationDate)
                                        FROM studentclass
                                        WHERE IdS = S.Id)


--------------------------------------------------------

SELECT B.Id, B.Name, B.LastName, B.DateOfBirth
FROM (SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, SC.IdS, SC.IdC, SC.CreationDate
          FROM students AS S INNER JOIN
               studentclass AS SC ON S.Id = SC.IdS
          WHERE SC.CreationDate = (SELECT MAX(CreationDate)
                                        FROM studentclass
                                        WHERE IdS = S.Id)) AS B 
WHERE B.IdC = ($1)


--------------------------------------------------------

SELECT C.Id, C.Year, C.Section, C.ScholarYearStart, M.name, COUNT(SC.Name) AS NumberStudents
FROM classes AS C INNER JOIN
          majors AS M ON C.IdM = M.Id LEFT JOIN
          allactivestudentclass AS SC ON C.Id = SC.IdC 
GROUP BY C.ScholarYearStart, C.Year, C.Section, C.Id, M.name
ORDER BY C.ScholarYearStart, C.Year, C.Section;

--------------------------------------------------------

SELECT *
FROM AUTO
WHERE modello IN ('Punto', 'Fiat');

SELECT *
FRom AUTO
WHERE cilindrata > ANY (SELECT cilindrata
                    FROM AUTO
                    WHERE modello = 'Punto');

SELECT *
FROM AUTO
WHERE modello IN (SELECT modello
                    FROM  AUTO
                    WHERE prduttore = 'fiat');
       

SELECT modello
FROM  AUTO
WHERE prduttore = 'fiat';

SELECT *
FROM AUTO
WHERE modello IN (SELECT modello
                    FROM  Proprietario
                    WHERE luogo = 'Cesena'
                    GROUP BY luogo
                    HAVING COUNT(*) > 2);

SELECT modello
FROM  Auto
WHERE modello = 'fiat'
GROUP BY modello
HAVING COUNT(*) > 2

SELECT *
FROM Ordini
Where luogo IN (SELECT luogo
          FROM Utenti
          WHere nome = 'Paolo');

SELECT *
FROM Auto
WHERE EXISTS (SELECT Modello
              FROM Modello
              WHERE Auto.Modello = Modello.Modello);


SELECT *
FROM Studenti AS S INNER JOIN
     Frequenta AS F ON S.id = F.IdStudenti INNER JOIN
     Corsi As C ON C.id = F.IdCorso;

CREATE TABLE Frequenta(
     IdStudenti INT NOT NULL,
     IdCorso INT NOT NULL,
     PRIMARY KEY (IdStudenti, IdCorso)
)


SELECT *
FROM Studenti AS S LEFT JOIN 
     VAlutazioni AS V ON S.matricola = V.matricola
WHERE S.nome = 'Paolo';