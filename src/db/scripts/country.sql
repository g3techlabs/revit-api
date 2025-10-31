--
-- PostgreSQL database dump
--

\restrict jB1ZgHz375D4hCFiIeOWXIH10gB4dFMB4tN77rmKzLAhnUBmMaDrOOwgJjTcPpY

-- Dumped from database version 18.0 (Debian 18.0-1.pgdg13+3)
-- Dumped by pg_dump version 18.0 (Debian 18.0-1.pgdg13+3)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: country; Type: TABLE; Schema: public; Owner: bruce
--

CREATE TABLE public.country (
    id bigint NOT NULL,
    name text NOT NULL,
    acronym text NOT NULL
);


ALTER TABLE public.country OWNER TO bruce;

--
-- Name: country_id_seq; Type: SEQUENCE; Schema: public; Owner: bruce
--

CREATE SEQUENCE public.country_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.country_id_seq OWNER TO bruce;

--
-- Name: country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bruce
--

ALTER SEQUENCE public.country_id_seq OWNED BY public.country.id;


--
-- Name: country id; Type: DEFAULT; Schema: public; Owner: bruce
--

ALTER TABLE ONLY public.country ALTER COLUMN id SET DEFAULT nextval('public.country_id_seq'::regclass);


--
-- Data for Name: country; Type: TABLE DATA; Schema: public; Owner: bruce
--

INSERT INTO public.country VALUES (1, 'Togo', 'TG');
INSERT INTO public.country VALUES (2, 'Mauritânia', 'MR');
INSERT INTO public.country VALUES (3, 'Israel', 'IL');
INSERT INTO public.country VALUES (4, 'Zimbabue', 'ZW');
INSERT INTO public.country VALUES (5, 'Bahamas, Ilhas', 'BS');
INSERT INTO public.country VALUES (6, 'Catar', 'QA');
INSERT INTO public.country VALUES (7, 'Ira, Republica Islâmica do', 'IR');
INSERT INTO public.country VALUES (8, 'África do Sul', 'ZA');
INSERT INTO public.country VALUES (9, 'Sudão', 'SD');
INSERT INTO public.country VALUES (10, 'Bermudas', 'BM');
INSERT INTO public.country VALUES (11, 'Líbia', 'LY');
INSERT INTO public.country VALUES (12, 'Belarus, Republica da', 'BY');
INSERT INTO public.country VALUES (13, 'Alemanha', 'DE');
INSERT INTO public.country VALUES (14, 'Gana', 'GH');
INSERT INTO public.country VALUES (15, 'Itália', 'IT');
INSERT INTO public.country VALUES (16, 'Quirguiz, Republica', 'KG');
INSERT INTO public.country VALUES (17, 'Curaçao', 'CW');
INSERT INTO public.country VALUES (18, 'Svalbard e Jan Mayen', 'SJ');
INSERT INTO public.country VALUES (19, 'Republika Srbija', 'RS');
INSERT INTO public.country VALUES (20, 'Arábia Saudita', 'SA');
INSERT INTO public.country VALUES (21, 'Gibraltar', 'GI');
INSERT INTO public.country VALUES (22, 'Emirados Árabes Unidos', 'AE');
INSERT INTO public.country VALUES (23, 'Wallis e Futuna, Ilhas', 'WF');
INSERT INTO public.country VALUES (24, 'Eswatini', 'SZ');
INSERT INTO public.country VALUES (25, 'Suriname', 'SR');
INSERT INTO public.country VALUES (26, 'Eslovênia, Republica da', 'SI');
INSERT INTO public.country VALUES (27, 'Albânia, Republica da', 'AL');
INSERT INTO public.country VALUES (28, 'Feroe, Ilhas', 'FO');
INSERT INTO public.country VALUES (29, 'Nepal', 'NP');
INSERT INTO public.country VALUES (30, 'Bulgária, Republica da', 'BG');
INSERT INTO public.country VALUES (31, 'Micronesia', 'FM');
INSERT INTO public.country VALUES (32, 'São Vicente e Granadinas', 'VC');
INSERT INTO public.country VALUES (33, 'Papua Nova Guine', 'PG');
INSERT INTO public.country VALUES (34, 'Mali', 'ML');
INSERT INTO public.country VALUES (35, 'Guine-Bissau', 'GW');
INSERT INTO public.country VALUES (36, 'Serra Leoa', 'SL');
INSERT INTO public.country VALUES (37, 'Marrocos', 'MA');
INSERT INTO public.country VALUES (38, 'Trinidad e Tobago', 'TT');
INSERT INTO public.country VALUES (39, 'Território Britânico do Oceano Indico', 'IO');
INSERT INTO public.country VALUES (40, 'Tunísia', 'TN');
INSERT INTO public.country VALUES (41, 'Dominica, Ilha', 'DM');
INSERT INTO public.country VALUES (42, 'Austrália', 'AU');
INSERT INTO public.country VALUES (43, 'Jordânia', 'JO');
INSERT INTO public.country VALUES (44, 'Nigéria', 'NG');
INSERT INTO public.country VALUES (45, 'Estados Unidos', 'US');
INSERT INTO public.country VALUES (46, 'Madagascar', 'MG');
INSERT INTO public.country VALUES (47, 'Butão', 'BT');
INSERT INTO public.country VALUES (48, 'São Martinho, Ilha de (Parte Francesa)', 'SM');
INSERT INTO public.country VALUES (49, 'Groenlândia', 'GL');
INSERT INTO public.country VALUES (50, 'Croácia (Republica da)', 'HR');
INSERT INTO public.country VALUES (51, 'Honduras', 'HN');
INSERT INTO public.country VALUES (52, 'Equador', 'EC');
INSERT INTO public.country VALUES (53, 'Malavi', 'MW');
INSERT INTO public.country VALUES (54, 'Guernsey, Ilha do Canal (Inclui Alderney e Sark)', 'GG');
INSERT INTO public.country VALUES (55, 'Sudao do Sul', 'SS');
INSERT INTO public.country VALUES (56, 'Marshall, Ilhas', 'MH');
INSERT INTO public.country VALUES (57, 'Hong Kong', 'HK');
INSERT INTO public.country VALUES (58, 'Vaticano, Estado da Cidade do', 'VA');
INSERT INTO public.country VALUES (59, 'Kiribati', 'KI');
INSERT INTO public.country VALUES (60, 'Guine', 'GN');
INSERT INTO public.country VALUES (61, 'Colômbia', 'CO');
INSERT INTO public.country VALUES (62, 'Martinica', 'MQ');
INSERT INTO public.country VALUES (63, 'Tcheca, Republica', 'CZ');
INSERT INTO public.country VALUES (64, 'Mayotte (Ilhas Francesas)', 'YT');
INSERT INTO public.country VALUES (65, 'Índia', 'IN');
INSERT INTO public.country VALUES (66, 'Coreia, Republica Popular Democrática da', 'KP');
INSERT INTO public.country VALUES (67, 'Irlanda', 'IE');
INSERT INTO public.country VALUES (68, 'Republica Dominicana', 'DO');
INSERT INTO public.country VALUES (69, 'Líbano', 'LB');
INSERT INTO public.country VALUES (70, 'Noruega', 'NO');
INSERT INTO public.country VALUES (71, 'Rússia, Federação da', 'RU');
INSERT INTO public.country VALUES (72, 'Tuvalu', 'TV');
INSERT INTO public.country VALUES (73, 'Países Baixos (Holanda)', 'NL');
INSERT INTO public.country VALUES (74, 'Virgens, Ilhas (E.U.A.)', 'VI');
INSERT INTO public.country VALUES (75, 'Moçambique', 'MZ');
INSERT INTO public.country VALUES (76, 'Uruguai', 'UY');
INSERT INTO public.country VALUES (77, 'Polinésia Francesa', 'PF');
INSERT INTO public.country VALUES (78, 'Guiana francesa', 'GF');
INSERT INTO public.country VALUES (79, 'Áustria', 'AT');
INSERT INTO public.country VALUES (80, 'Malásia', 'MY');
INSERT INTO public.country VALUES (81, 'Gambia', 'GM');
INSERT INTO public.country VALUES (82, 'Brasil', 'BR');
INSERT INTO public.country VALUES (83, 'Cook, Ilhas', 'CK');
INSERT INTO public.country VALUES (84, 'Malta', 'MT');
INSERT INTO public.country VALUES (85, 'Tanzânia, Republica Unida da', 'TZ');
INSERT INTO public.country VALUES (86, 'Norfolk, Ilha', 'NF');
INSERT INTO public.country VALUES (87, 'Congo', 'CG');
INSERT INTO public.country VALUES (88, 'Nicarágua', 'NI');
INSERT INTO public.country VALUES (89, 'Panamá', 'PA');
INSERT INTO public.country VALUES (90, 'Peru', 'PE');
INSERT INTO public.country VALUES (91, 'Palau', 'PW');
INSERT INTO public.country VALUES (92, 'Camarões', 'CM');
INSERT INTO public.country VALUES (93, 'Espanha', 'ES');
INSERT INTO public.country VALUES (94, 'Porto Rico', 'PR');
INSERT INTO public.country VALUES (95, 'Turquia', 'TR');
INSERT INTO public.country VALUES (96, 'Indonésia', 'ID');
INSERT INTO public.country VALUES (97, 'Man, Ilha de', 'IM');
INSERT INTO public.country VALUES (98, 'Laos, Republica Popular Democrática do', 'LA');
INSERT INTO public.country VALUES (99, 'Niue, Ilha', 'NU');
INSERT INTO public.country VALUES (100, 'Eslovaca, Republica', 'SK');
INSERT INTO public.country VALUES (101, 'Bouvet, Ilha', 'BV');
INSERT INTO public.country VALUES (102, 'Salomão, Ilhas', 'SB');
INSERT INTO public.country VALUES (103, 'Bósnia-herzegovina (Republica da)', 'BA');
INSERT INTO public.country VALUES (104, 'Suíça', 'CH');
INSERT INTO public.country VALUES (105, 'Bangladesh', 'BD');
INSERT INTO public.country VALUES (106, 'Mianmar (Birmânia)', 'MM');
INSERT INTO public.country VALUES (107, 'Chile', 'CL');
INSERT INTO public.country VALUES (108, 'Barbados', 'BB');
INSERT INTO public.country VALUES (109, 'Síria, Republica Árabe da', 'SY');
INSERT INTO public.country VALUES (110, 'Liechtenstein', 'LI');
INSERT INTO public.country VALUES (111, 'Iraque', 'IQ');
INSERT INTO public.country VALUES (112, 'Jamaica', 'JM');
INSERT INTO public.country VALUES (113, 'Uzbequistão, Republica do', 'UZ');
INSERT INTO public.country VALUES (114, 'Burkina Faso', 'BF');
INSERT INTO public.country VALUES (115, 'Bolívia', 'BO');
INSERT INTO public.country VALUES (116, 'Virgens, Ilhas (Britânicas)', 'VG');
INSERT INTO public.country VALUES (117, 'Ucrânia', 'UA');
INSERT INTO public.country VALUES (118, 'São Cristovão e Neves, Ilhas', 'KN');
INSERT INTO public.country VALUES (119, 'Níger', 'NE');
INSERT INTO public.country VALUES (120, 'Etiópia', 'ET');
INSERT INTO public.country VALUES (121, 'Coreia, Republica da', 'KR');
INSERT INTO public.country VALUES (122, 'Nauru', 'NR');
INSERT INTO public.country VALUES (123, 'Costa do Marfim', 'CI');
INSERT INTO public.country VALUES (124, 'El Salvador', 'SV');
INSERT INTO public.country VALUES (125, 'Filipinas', 'PH');
INSERT INTO public.country VALUES (126, 'Senegal', 'SN');
INSERT INTO public.country VALUES (127, 'Venezuela', 'VE');
INSERT INTO public.country VALUES (128, 'Bonaire', 'AN');
INSERT INTO public.country VALUES (129, 'Costa Rica', 'CR');
INSERT INTO public.country VALUES (130, 'Egito', 'EG');
INSERT INTO public.country VALUES (131, 'Belize', 'BZ');
INSERT INTO public.country VALUES (132, 'Gabão', 'GA');
INSERT INTO public.country VALUES (133, 'Bélgica', 'BE');
INSERT INTO public.country VALUES (134, 'Islândia', 'IS');
INSERT INTO public.country VALUES (135, 'Djibuti', 'DJ');
INSERT INTO public.country VALUES (136, 'Guine-Equatorial', 'GQ');
INSERT INTO public.country VALUES (137, 'Aruba', 'AW');
INSERT INTO public.country VALUES (138, 'Montenegro', 'ME');
INSERT INTO public.country VALUES (139, 'Timor Leste', 'TL');
INSERT INTO public.country VALUES (140, 'Bahrein, Ilhas', 'BH');
INSERT INTO public.country VALUES (141, 'Antigua e Barbuda', 'AG');
INSERT INTO public.country VALUES (142, 'Angola', 'AO');
INSERT INTO public.country VALUES (143, 'Sri Lanka', 'LK');
INSERT INTO public.country VALUES (144, 'São Tome e Príncipe, Ilhas', 'ST');
INSERT INTO public.country VALUES (145, 'Zâmbia', 'ZM');
INSERT INTO public.country VALUES (146, 'Tadjiquistao, Republica do', 'TJ');
INSERT INTO public.country VALUES (147, 'Lituânia, Republica da', 'LT');
INSERT INTO public.country VALUES (148, 'Brunei', 'BN');
INSERT INTO public.country VALUES (149, 'Dinamarca', 'DK');
INSERT INTO public.country VALUES (150, 'Vietnã', 'VN');
INSERT INTO public.country VALUES (151, 'Iémen', 'YE');
INSERT INTO public.country VALUES (152, 'Letônia, Republica da', 'LV');
INSERT INTO public.country VALUES (153, 'Polônia, Republica da', 'PL');
INSERT INTO public.country VALUES (154, 'Tonga', 'TO');
INSERT INTO public.country VALUES (155, 'Pitcairn, Ilha', 'PN');
INSERT INTO public.country VALUES (156, 'San Marino', 'SM');
INSERT INTO public.country VALUES (157, 'Macau', 'MO');
INSERT INTO public.country VALUES (158, 'Chipre', 'CY');
INSERT INTO public.country VALUES (159, 'Guadalupe', 'GP');
INSERT INTO public.country VALUES (160, 'Republica Centro-Africana', 'CF');
INSERT INTO public.country VALUES (161, 'Estônia, Republica da', 'EE');
INSERT INTO public.country VALUES (162, 'Luxemburgo', 'LU');
INSERT INTO public.country VALUES (163, 'Seychelles', 'SC');
INSERT INTO public.country VALUES (164, 'Canadá', 'CA');
INSERT INTO public.country VALUES (165, 'Romênia', 'RO');
INSERT INTO public.country VALUES (166, 'Comores, Ilhas', 'KM');
INSERT INTO public.country VALUES (167, 'Andorra', 'AD');
INSERT INTO public.country VALUES (168, 'Turcomenistão, Republica do', 'TM');
INSERT INTO public.country VALUES (169, 'Georgia, Republica da', 'GE');
INSERT INTO public.country VALUES (170, 'Botsuana', 'BW');
INSERT INTO public.country VALUES (171, 'Samoa', 'WS');
INSERT INTO public.country VALUES (172, 'Jersey, Ilha do Canal', 'JE');
INSERT INTO public.country VALUES (173, 'Armênia, Republica da', 'AM');
INSERT INTO public.country VALUES (174, 'São Martinho, Ilha de (Parte Holandesa)', 'SM');
INSERT INTO public.country VALUES (175, 'Ruanda', 'RW');
INSERT INTO public.country VALUES (176, 'Ilha Heard e Ilhas McDonald', 'HM');
INSERT INTO public.country VALUES (177, 'Formosa (Taiwan)', 'TW');
INSERT INTO public.country VALUES (178, 'Eritreia', 'ER');
INSERT INTO public.country VALUES (179, 'Argélia', 'DZ');
INSERT INTO public.country VALUES (180, 'Granada', 'GD');
INSERT INTO public.country VALUES (181, 'Samoa Americana', 'AS');
INSERT INTO public.country VALUES (182, 'México', 'MX');
INSERT INTO public.country VALUES (183, 'Afeganistão', 'AF');
INSERT INTO public.country VALUES (184, 'Mauricio', 'MU');
INSERT INTO public.country VALUES (185, 'Ilhas Caimã', 'KY');
INSERT INTO public.country VALUES (186, 'Antartica', 'AQ');
INSERT INTO public.country VALUES (187, 'Suécia', 'SE');
INSERT INTO public.country VALUES (188, 'Toquelau, Ilhas', 'TK');
INSERT INTO public.country VALUES (189, 'Libéria', 'LR');
INSERT INTO public.country VALUES (190, 'Nova Caledonia', 'NC');
INSERT INTO public.country VALUES (191, 'Argentina', 'AR');
INSERT INTO public.country VALUES (192, 'Santa Helena', 'SH');
INSERT INTO public.country VALUES (193, 'Kuwait', 'KW');
INSERT INTO public.country VALUES (194, 'Finlândia', 'FI');
INSERT INTO public.country VALUES (195, 'Ilha Herad e Ilhas Macdonald', 'AU');
INSERT INTO public.country VALUES (196, 'Ilhas Menores Distantes dos Estados Unidos', 'UM');
INSERT INTO public.country VALUES (197, 'Namíbia', 'NA');
INSERT INTO public.country VALUES (198, 'Somalia', 'SO');
INSERT INTO public.country VALUES (199, 'Oma', 'OM');
INSERT INTO public.country VALUES (200, 'Guam', 'GU');
INSERT INTO public.country VALUES (201, 'Azerbaijão, Republica do', 'AZ');
INSERT INTO public.country VALUES (202, 'Tailândia', 'TH');
INSERT INTO public.country VALUES (203, 'Benin', 'BJ');
INSERT INTO public.country VALUES (204, 'Cazaquistão, Republica do', 'KZ');
INSERT INTO public.country VALUES (205, 'Turcas e Caicos, Ilhas', 'TC');
INSERT INTO public.country VALUES (206, 'China, Republica Popular', 'CN');
INSERT INTO public.country VALUES (207, 'Coletividade de São Bartolomeu', 'BL');
INSERT INTO public.country VALUES (208, 'Saara Ocidental', 'EH');
INSERT INTO public.country VALUES (209, 'Cingapura', 'SG');
INSERT INTO public.country VALUES (210, 'São Bartolomeu', 'FR');
INSERT INTO public.country VALUES (211, 'Maldivas', 'MV');
INSERT INTO public.country VALUES (212, 'Haiti', 'HT');
INSERT INTO public.country VALUES (213, 'Mongólia', 'MN');
INSERT INTO public.country VALUES (214, 'Moldávia, Republica da', 'MD');
INSERT INTO public.country VALUES (215, 'Uganda', 'UG');
INSERT INTO public.country VALUES (216, 'Hungria, Republica da', 'HU');
INSERT INTO public.country VALUES (217, 'Terras Austrais e Antárcticas Francesas', 'TF');
INSERT INTO public.country VALUES (218, 'Marianas do Norte', 'MP');
INSERT INTO public.country VALUES (219, 'Iugoslávia, República Fed. da', 'YU');
INSERT INTO public.country VALUES (220, 'Lesoto', 'LS');
INSERT INTO public.country VALUES (221, 'Palestina', 'PS');
INSERT INTO public.country VALUES (222, 'Chade', 'TD');
INSERT INTO public.country VALUES (223, 'Reino Unido', 'GB');
INSERT INTO public.country VALUES (224, 'São Pedro e Miquelon', 'PM');
INSERT INTO public.country VALUES (225, 'Guatemala', 'GT');
INSERT INTO public.country VALUES (226, 'Cocos (Keeling), Ilhas', 'CC');
INSERT INTO public.country VALUES (227, 'Anguila', 'AI');
INSERT INTO public.country VALUES (228, 'Falkland (Ilhas Malvinas)', 'FK');
INSERT INTO public.country VALUES (229, 'Japão', 'JP');
INSERT INTO public.country VALUES (230, 'Quênia', 'KE');
INSERT INTO public.country VALUES (231, 'Portugal', 'PT');
INSERT INTO public.country VALUES (232, 'Burundi', 'BI');
INSERT INTO public.country VALUES (233, 'Ilhas Geórgia do Sul e Sandwich do Sul', 'GS');
INSERT INTO public.country VALUES (234, 'Antártida', 'AQ');
INSERT INTO public.country VALUES (235, 'Macedônia do Norte', 'MK');
INSERT INTO public.country VALUES (236, 'Fiji', 'FJ');
INSERT INTO public.country VALUES (237, 'Guiana', 'GY');
INSERT INTO public.country VALUES (238, 'Vanuatu', 'VU');
INSERT INTO public.country VALUES (239, 'Camboja', 'KH');
INSERT INTO public.country VALUES (240, 'Cabo Verde, Republica de', 'CV');
INSERT INTO public.country VALUES (241, 'Congo, Republica Democrática do', 'CD');
INSERT INTO public.country VALUES (242, 'Reunião, Ilha', 'RE');
INSERT INTO public.country VALUES (243, 'Montserrat, Ilhas', 'MS');
INSERT INTO public.country VALUES (244, 'Franca', 'FR');
INSERT INTO public.country VALUES (245, 'Cuba', 'CU');
INSERT INTO public.country VALUES (246, 'Mônaco', 'MC');
INSERT INTO public.country VALUES (247, 'Paraguai', 'PY');
INSERT INTO public.country VALUES (248, 'Terras Austrais e Antárticas Francesas', 'TF');
INSERT INTO public.country VALUES (249, 'Grécia', 'GR');
INSERT INTO public.country VALUES (250, 'Nova Zelândia', 'NZ');
INSERT INTO public.country VALUES (251, 'Santa Lucia', 'LC');
INSERT INTO public.country VALUES (252, 'Christmas, Ilha (Navidad)', 'CX');
INSERT INTO public.country VALUES (253, 'Aland, Ilhas', 'AX');
INSERT INTO public.country VALUES (254, 'Paquistão', 'PK');


--
-- Name: country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bruce
--

SELECT pg_catalog.setval('public.country_id_seq', 254, true);


--
-- Name: country country_pkey; Type: CONSTRAINT; Schema: public; Owner: bruce
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT country_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\unrestrict jB1ZgHz375D4hCFiIeOWXIH10gB4dFMB4tN77rmKzLAhnUBmMaDrOOwgJjTcPpY

