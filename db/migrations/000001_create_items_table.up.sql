CREATE TABLE IF NOT EXISTS items(
id SERIAL PRIMARY KEY,
usuario VARCHAR(100) NOT NULL,
fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
fecha_actualizacion TIMESTAMP,
estado_abierto BOOLEAN DEFAULT TRUE
);