INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com","$2a$10$2tnDSaLzIgqrK7NGbglg0.HyNPB4DOiY/7CLVPnd.pAIh7GlzmrAC"),
("Usuario 2", "usuario_2", "usuario2@gmail.com","$2a$10$2tnDSaLzIgqrK7NGbglg0.HyNPB4DOiY/7CLVPnd.pAIh7GlzmrAC"),
("Usuario 3", "usuario_3", "usuario3@gmail.com","$2a$10$2tnDSaLzIgqrK7NGbglg0.HyNPB4DOiY/7CLVPnd.pAIh7GlzmrAC"),
("Usuario 4", "usuario_4", "usuario4@gmail.com","$2a$10$2tnDSaLzIgqrK7NGbglg0.HyNPB4DOiY/7CLVPnd.pAIh7GlzmrAC"),
("Usuario 5", "usuario_5", "usuario5@gmail.com","$2a$10$2tnDSaLzIgqrK7NGbglg0.HyNPB4DOiY/7CLVPnd.pAIh7GlzmrAC");


INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1,2),
(1,3),
(1,4),
(1,5),
(2,1),
(5,1);