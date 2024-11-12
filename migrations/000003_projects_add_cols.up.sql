ALTER TABLE IF EXISTS public.projects
ADD is_active BOOLEAN NULL,
ADD is_archived BOOLEAN NULL,
ADD is_developing BOOLEAN NULL;