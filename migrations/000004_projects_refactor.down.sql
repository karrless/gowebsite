ALTER TABLE IF EXISTS public.projects
ALTER COLUMN is_active SET NULL,
ALTER COLUMN is_archived SET NULL,
ALTER COLUMN is_developing SET NULL,
DROP COLUMN gh_link TEXT NULL,
DROP COLUMN tg_link TEXT NULL,
DROP COLUMN http_link TEXT NULL;