--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: access_tokens; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.access_tokens (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    token uuid NOT NULL,
    principal character varying(255) NOT NULL,
    scopes character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.access_tokens OWNER TO launchbox;

--
-- Name: agents; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.agents (
    id uuid NOT NULL,
    pod_name character varying(255) DEFAULT 'null'::character varying,
    ip_address character varying(255) DEFAULT 'null'::character varying,
    status character varying(255) DEFAULT 'pending'::character varying NOT NULL,
    cluster_id uuid NOT NULL,
    last_check_in timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.agents OWNER TO launchbox;

--
-- Name: applications; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.applications (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    namespace character varying(255) NOT NULL,
    tags character varying(255) DEFAULT ''::character varying NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.applications OWNER TO launchbox;

--
-- Name: cluster_applications; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.cluster_applications (
    cluster_id uuid NOT NULL,
    application_id uuid NOT NULL,
    status character varying(255) DEFAULT 'pending'::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.cluster_applications OWNER TO launchbox;

--
-- Name: clusters; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.clusters (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    provider character varying(255) DEFAULT 'null'::character varying,
    region character varying(255) DEFAULT 'null'::character varying,
    version character varying(255) DEFAULT 'null'::character varying,
    tags character varying(255) DEFAULT ''::character varying NOT NULL,
    owner_id uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid,
    owner_type character varying(255) DEFAULT 'user'::character varying NOT NULL,
    last_check_in timestamp without time zone,
    status character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.clusters OWNER TO launchbox;

--
-- Name: projects; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.projects (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    status character varying(255),
    application_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.projects OWNER TO launchbox;

--
-- Name: revisions; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.revisions (
    id uuid NOT NULL,
    status character varying(255),
    commit_sha character varying(255),
    project_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.revisions OWNER TO launchbox;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO launchbox;

--
-- Name: secrets; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.secrets (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    sensitive boolean DEFAULT true NOT NULL,
    owner_type character varying(255),
    owner_id uuid,
    cluster_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.secrets OWNER TO launchbox;

--
-- Name: users; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    email_verified boolean DEFAULT false NOT NULL,
    password_hash character varying(255) NOT NULL,
    avatar_url character varying(255),
    name character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO launchbox;

--
-- Name: vcs_connections; Type: TABLE; Schema: public; Owner: launchbox
--

CREATE TABLE public.vcs_connections (
    id uuid NOT NULL,
    provider character varying(255) NOT NULL,
    hostname character varying(255) DEFAULT 'null'::character varying,
    name character varying(255) DEFAULT 'null'::character varying,
    email character varying(255) NOT NULL,
    nickname character varying(255) DEFAULT 'null'::character varying,
    provider_user_id character varying(255) DEFAULT 'null'::character varying,
    access_token character varying(255) DEFAULT 'null'::character varying,
    expires_at timestamp without time zone,
    refresh_token character varying(255) DEFAULT 'null'::character varying,
    vcs_connection_id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.vcs_connections OWNER TO launchbox;

--
-- Name: access_tokens access_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.access_tokens
    ADD CONSTRAINT access_tokens_pkey PRIMARY KEY (id);


--
-- Name: agents agents_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.agents
    ADD CONSTRAINT agents_pkey PRIMARY KEY (id);


--
-- Name: applications applications_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pkey PRIMARY KEY (id);


--
-- Name: cluster_applications cluster_applications_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.cluster_applications
    ADD CONSTRAINT cluster_applications_pkey PRIMARY KEY (cluster_id, application_id);


--
-- Name: clusters clusters_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.clusters
    ADD CONSTRAINT clusters_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: revisions revisions_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.revisions
    ADD CONSTRAINT revisions_pkey PRIMARY KEY (id);


--
-- Name: secrets secrets_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.secrets
    ADD CONSTRAINT secrets_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: vcs_connections vcs_connections_pkey; Type: CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.vcs_connections
    ADD CONSTRAINT vcs_connections_pkey PRIMARY KEY (id);


--
-- Name: access_tokens_token_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX access_tokens_token_idx ON public.access_tokens USING btree (token);


--
-- Name: applications_name_user_id_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX applications_name_user_id_idx ON public.applications USING btree (name, user_id);


--
-- Name: applications_namespace_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX applications_namespace_idx ON public.applications USING btree (namespace);


--
-- Name: clusters_name_owner_id_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX clusters_name_owner_id_idx ON public.clusters USING btree (name, owner_id);


--
-- Name: clusters_token_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX clusters_token_idx ON public.clusters USING btree (token);


--
-- Name: projects_application_id_name_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX projects_application_id_name_idx ON public.projects USING btree (application_id, name);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: vcs_connections_provider_user_id_idx; Type: INDEX; Schema: public; Owner: launchbox
--

CREATE UNIQUE INDEX vcs_connections_provider_user_id_idx ON public.vcs_connections USING btree (provider, user_id);


--
-- Name: agents agents_cluster_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.agents
    ADD CONSTRAINT agents_cluster_id_fkey FOREIGN KEY (cluster_id) REFERENCES public.clusters(id);


--
-- Name: applications applications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: projects projects_application_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_application_id_fkey FOREIGN KEY (application_id) REFERENCES public.applications(id);


--
-- Name: revisions revisions_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.revisions
    ADD CONSTRAINT revisions_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: vcs_connections vcs_connections_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.vcs_connections
    ADD CONSTRAINT vcs_connections_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: vcs_connections vcs_connections_vcs_connection_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: launchbox
--

ALTER TABLE ONLY public.vcs_connections
    ADD CONSTRAINT vcs_connections_vcs_connection_id_fkey FOREIGN KEY (vcs_connection_id) REFERENCES public.vcs_connections(id);


--
-- PostgreSQL database dump complete
--

