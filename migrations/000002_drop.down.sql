CREATE TABLE IF NOT EXISTS public.tags
(
  id         serial  NOT NULL,
  name       TEXT NOT NULL,
  color_name TEXT NOT NULL DEFAULT 'default',
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.projecttabs
(
  tag_id INTEGER NOT NULL,
  project_id INTEGER NOT NULL,
  FOREIGN KEY (tag_id) REFERENCES tags(id),
  FOREIGN KEY (project_id) REFERENCES projects(id)
);

ALTER TABLE IF EXISTS public."tags"
    OWNER to postgres;

ALTER TABLE IF EXISTS public."projecttabs"
    OWNER to postgres;