CREATE TABLE IF NOT EXISTS public.techs
(
  id   serial  NOT NULL,
  name TEXT NOT NULL,
  svg  TEXT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.projects
(
  id          serial  NOT NULL,
  title       TEXT NOT NULL,
  version     TEXT NOT NULL DEFAULT '0.1.0',
  description TEXT NOT NULL,
  is_active   BOOLEAN NOT NULL,
  is_archived BOOLEAN NOT NULL,
  is_developing BOOLEAN NOT NULL,
  links TEXT[] NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.project_tech
(
  tech_id INTEGER NOT NULL,
  project_id INTEGER NOT NULL,
  FOREIGN KEY (tech_id) REFERENCES techs(id),
  FOREIGN KEY (project_id) REFERENCES projects(id)
);


ALTER TABLE IF EXISTS public."languages"
    OWNER to postgres;

ALTER TABLE IF EXISTS public."projects"
    OWNER to postgres;
ALTER TABLE IF EXISTS public."techs"
    OWNER to postgres;

ALTER TABLE IF EXISTS public."project_tech"
    OWNER to postgres;