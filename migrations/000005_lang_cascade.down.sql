ALTER TABLE IF EXISTS public.projects
DROP CONSTRAINT IF EXISTS projects_lang_id_fkey;

ALTER TABLE IF EXISTS public.projects
ADD CONSTRAINT projects_lang_id_fkey FOREIGN KEY (lang_id) REFERENCES public.languages(id);