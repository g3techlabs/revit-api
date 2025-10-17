--
-- PostgreSQL database dump
--

\restrict BnruZUix6TCkQyaPeTgFKVcxSSHQwDxYSpKLJ2VgruSTq1U48AVCvsrfy5dheze

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
-- Name: state; Type: TABLE; Schema: public; Owner: bruce
--

CREATE TABLE public.state (
    id bigint NOT NULL,
    name text NOT NULL,
    acronym text NOT NULL,
    country_id bigint NOT NULL
);


ALTER TABLE public.state OWNER TO bruce;

--
-- Name: state_id_seq; Type: SEQUENCE; Schema: public; Owner: bruce
--

CREATE SEQUENCE public.state_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.state_id_seq OWNER TO bruce;

--
-- Name: state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bruce
--

ALTER SEQUENCE public.state_id_seq OWNED BY public.state.id;


--
-- Name: state id; Type: DEFAULT; Schema: public; Owner: bruce
--

ALTER TABLE ONLY public.state ALTER COLUMN id SET DEFAULT nextval('public.state_id_seq'::regclass);


--
-- Data for Name: state; Type: TABLE DATA; Schema: public; Owner: bruce
--

INSERT INTO public.state VALUES (1, 'Acre', 'AC', 82);
INSERT INTO public.state VALUES (2, 'Alagoas', 'AL', 82);
INSERT INTO public.state VALUES (3, 'Amazonas', 'AM', 82);
INSERT INTO public.state VALUES (4, 'Amapá', 'AP', 82);
INSERT INTO public.state VALUES (5, 'Bahia', 'BA', 82);
INSERT INTO public.state VALUES (6, 'Ceará', 'CE', 82);
INSERT INTO public.state VALUES (7, 'Distrito Federal', 'DF', 82);
INSERT INTO public.state VALUES (8, 'Espírito Santo', 'ES', 82);
INSERT INTO public.state VALUES (9, 'Goiás', 'GO', 82);
INSERT INTO public.state VALUES (10, 'Maranhão', 'MA', 82);
INSERT INTO public.state VALUES (11, 'Minas Gerais', 'MG', 82);
INSERT INTO public.state VALUES (12, 'Mato Grosso do Sul', 'MS', 82);
INSERT INTO public.state VALUES (13, 'Mato Grosso', 'MT', 82);
INSERT INTO public.state VALUES (14, 'Pará', 'PA', 82);
INSERT INTO public.state VALUES (15, 'Paraíba', 'PB', 82);
INSERT INTO public.state VALUES (16, 'Pernambuco', 'PE', 82);
INSERT INTO public.state VALUES (17, 'Piauí', 'PI', 82);
INSERT INTO public.state VALUES (18, 'Paraná', 'PR', 82);
INSERT INTO public.state VALUES (19, 'Rio de Janeiro', 'RJ', 82);
INSERT INTO public.state VALUES (20, 'Rio Grande do Norte', 'RN', 82);
INSERT INTO public.state VALUES (21, 'Rondônia', 'RO', 82);
INSERT INTO public.state VALUES (22, 'Roraima', 'RR', 82);
INSERT INTO public.state VALUES (23, 'Rio Grande do Sul', 'RS', 82);
INSERT INTO public.state VALUES (24, 'Santa Catarina', 'SC', 82);
INSERT INTO public.state VALUES (25, 'Sergipe', 'SE', 82);
INSERT INTO public.state VALUES (26, 'São Paulo', 'SP', 82);
INSERT INTO public.state VALUES (27, 'Tocantins', 'TO', 82);


--
-- Name: state_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bruce
--

SELECT pg_catalog.setval('public.state_id_seq', 27, true);


--
-- Name: state state_pkey; Type: CONSTRAINT; Schema: public; Owner: bruce
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT state_pkey PRIMARY KEY (id);


--
-- Name: state fk_country_states; Type: FK CONSTRAINT; Schema: public; Owner: bruce
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT fk_country_states FOREIGN KEY (country_id) REFERENCES public.country(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict BnruZUix6TCkQyaPeTgFKVcxSSHQwDxYSpKLJ2VgruSTq1U48AVCvsrfy5dheze

