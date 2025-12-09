CREATE TABLE IF NOT EXISTS state (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL,
    acronym text NOT NULL,
    country_id bigint NOT NULL REFERENCES country(id) ON DELETE CASCADE
);

INSERT INTO state VALUES (1, 'Acre', 'AC', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (2, 'Alagoas', 'AL', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (3, 'Amazonas', 'AM', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (4, 'Amapá', 'AP', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (5, 'Bahia', 'BA', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (6, 'Ceará', 'CE', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (7, 'Distrito Federal', 'DF', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (8, 'Espírito Santo', 'ES', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (9, 'Goiás', 'GO', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (10, 'Maranhão', 'MA', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (11, 'Minas Gerais', 'MG', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (12, 'Mato Grosso do Sul', 'MS', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (13, 'Mato Grosso', 'MT', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (14, 'Pará', 'PA', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (15, 'Paraíba', 'PB', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (16, 'Pernambuco', 'PE', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (17, 'Piauí', 'PI', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (18, 'Paraná', 'PR', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (19, 'Rio de Janeiro', 'RJ', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (20, 'Rio Grande do Norte', 'RN', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (21, 'Rondônia', 'RO', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (22, 'Roraima', 'RR', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (23, 'Rio Grande do Sul', 'RS', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (24, 'Santa Catarina', 'SC', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (25, 'Sergipe', 'SE', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (26, 'São Paulo', 'SP', 82) ON CONFLICT (id) DO NOTHING;
INSERT INTO state VALUES (27, 'Tocantins', 'TO', 82) ON CONFLICT (id) DO NOTHING;

-- Ajusta a sequência para o maior ID existente para evitar erros em novos inserts
SELECT setval('state_id_seq', (SELECT MAX(id) FROM state));
